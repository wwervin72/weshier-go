package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"weshierNext/model/request"
	"weshierNext/model/response"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"
	"weshierNext/pkg/validate"
	"weshierNext/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ArticleModel article model
type ArticleModel struct {
	BaseModel
	Title           string              `zh:"文章标题" json:"title" gorm:"not null;column:title" validate:"required,min=1"`
	Abstract        string              `zh:"文章摘要" json:"abstract" gorm:"column:abstract;type:text"`
	Content         string              `zh:"文章内容" json:"content" gorm:"not null;column:content;type:text" validate:"required,min=1"`
	AllowComment    int8                `zh:"允许评论" json:"allowComment" gorm:"column:allow_comment;default:0"`
	Password        string              `zh:"允许评论" json:"password" gorm:"column:password"`
	Thumbnail       string              `zh:"缩略图" json:"thumbnail" gorm:"column:thumbnail"`
	ThumbnailEntity ThumbnailModel      `zh:"缩略图" json:"thumbnailEntity" gorm:"foreignKey:Thumbnail;"`
	Topping         int8                `zh:"置顶" json:"topping" gorm:"column:topping;default:0"`
	AuthorID        uint64              `zh:"作者" json:"authorId" gorm:"column:author_id;"`
	Author          UserModel           `json:"author" validate:"-" gorm:"foreignKey:AuthorID"`
	CategoryID      uint64              `zh:"分类" json:"categoryId" gorm:"column:category_id" validate:"required"`
	Tags            []uint64            `zh:"标签" json:"tags" gorm:"-"`
	TagsEntity      []ArticleTagModel   `zh:"标签" json:"tagsEntity" gorm:"FOREIGNKEY:ArticleID;ASSOCIATION_FOREIGNKEY:ID"`
	Annexs          []uint64            `zh:"附件" json:"annexs" gorm:"-"`
	AnnexsEntity    []ArticleAnnexModel `zh:"附件" json:"annexsEntity" gorm:"FOREIGNKEY:ArticleID;ASSOCIATION_FOREIGNKEY:ID"`
}

type ArticleReqStruct struct {
	ArticleModel
	Author UserModel `json:"author"`
}

// TableName article table name
func (am *ArticleModel) TableName() string {
	return "ws_article"
}

func (am *ArticleModel) AfterFind(tx *gorm.DB) (err error) {
	am.Thumbnail = strings.TrimSpace(am.Thumbnail)
	// 如果有未找到的缩略图 就使用默认的
	if am.Thumbnail == "" || am.ThumbnailEntity.ID == 0 {
		am.Thumbnail = fmt.Sprintf("%s/asset/image/thumbnail.jpg", viper.GetString("publicPrefix"))
	} else if am.ThumbnailEntity.ID != 0 {
		am.Thumbnail = am.ThumbnailEntity.Url
	}
	var tags []uint64
	for _, tag := range am.TagsEntity {
		tags = append(tags, tag.TagID)
	}
	if am.Tags == nil {
		am.Tags = tags
	}
	am.CreatedTime = util.TimeFormat(*am.CreatedAt)
	am.UpdatedTime = util.TimeFormat(*am.UpdatedAt)
	return
}

// Create create article
func (am *ArticleModel) Create() error {
	return DB.Self.Model(am).Create(&am).Error
}

// Delete delete article
func (am *ArticleModel) Delete() error {
	return DB.Self.Model(am).Delete(&am).Error
}

// Validate Validate the field
func (am *ArticleModel) Validate() error {
	return validate.Validate(*am)
}

func CreateArticle(c *gin.Context, data *request.CreateArticleReqStruct, userId uint64) (error, *ArticleModel) {
	article := &ArticleModel{
		Title:        data.Title,
		Content:      data.Content,
		Abstract:     data.Abstract,
		AllowComment: data.AllowComment,
		Password:     data.Password,
		Thumbnail:    data.Thumbnail,
		Topping:      data.Topping,
		AuthorID:     userId,
		CategoryID:   data.CategoryID,
		Tags:         data.Tags,
		Annexs:       data.Annexs,
	}
	err := article.Validate()
	if err != nil {
		logger.Logger.Debug("admin user validate failed", zap.String("error", err.Error()))
		return err, nil
	}
	err, _ = QueryCategoryByIDAndUser(data.CategoryID, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &errno.Errno{
				Code:    errno.ErrNotFound.Code,
				Message: "分类不存在",
			}, nil
		}
		logger.Logger.Debug("query category by id failed", zap.String("error", err.Error()))
		return errno.ErrDatabase, nil
	}
	tx := DB.Self.Model(article).Begin()
	err = tx.Create(article).Error
	if err != nil {
		tx.Rollback()
		logger.Logger.Debug("article created failed", zap.String("error", err.Error()))
		return errno.ErrDatabase, nil
	}
	for _, tagId := range data.Tags {
		err, _ := QueryTagByIDAndUser(tagId, userId)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return &errno.Errno{
					Code:    errno.ErrNotFound.Code,
					Message: "标签不存在",
				}, nil
			}
			logger.Logger.Error("query tag by id failed", zap.String("error", err.Error()))
			return errno.ErrDatabase, nil
		}
		articleTag := &ArticleTagModel{
			TagID:     tagId,
			ArticleID: article.ID,
		}
		err = tx.Model(articleTag).Create(articleTag).Error
		if err != nil {
			tx.Rollback()
			logger.Logger.Error("articleTag created failed", zap.String("error", err.Error()))
			return errno.ErrDatabase, nil
		}
	}
	for _, annexId := range data.Annexs {
		err, _ := QueryAnnexByIDAndUser(annexId, userId)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return &errno.Errno{
					Code:    errno.ErrNotFound.Code,
					Message: "附件不存在",
				}, nil
			}
			logger.Logger.Error("query annex by id failed", zap.String("error", err.Error()))
			return errno.ErrDatabase, nil
		}
		articleAnnex := &ArticleAnnexModel{
			AnnexID:   annexId,
			ArticleID: article.ID,
		}
		err = tx.Model(articleAnnex).Create(articleAnnex).Error
		if err != nil {
			tx.Rollback()
			logger.Logger.Error("articleAnnex created failed", zap.String("error", err.Error()))
			return errno.ErrDatabase, nil
		}
	}
	return tx.Commit().Error, article
}

func UpdateArticle(c *gin.Context, data *request.UpdateArticleReqStruct, userId uint64) error {
	am := &ArticleModel{
		Content:      data.Content,
		Title:        data.Title,
		Abstract:     data.Abstract,
		Topping:      data.Topping,
		AllowComment: data.AllowComment,
		Thumbnail:    data.Thumbnail,
		Tags:         data.Tags,
		CategoryID:   data.CategoryID,
		Password:     data.Password,
	}
	err := am.Validate()
	if err != nil {
		logger.Logger.Debug("article update validate failed", zap.String("error", err.Error()))
		return err
	}
	j, err := json.Marshal(am)
	if err != nil {
		return errno.InternalServerError
	}
	aMap := make(map[string]interface{})
	err = json.Unmarshal(j, &aMap)
	if err != nil {
		return errno.InternalServerError
	}
	tx := DB.Self.Model(&am).Begin()
	atx := tx.Select("title", "abstract", "content", "allow_comment", "password", "thumbnail", "topping", "category_id").Where("author_id=? and id=?", userId, data.ID).Updates(aMap)
	if atx.RowsAffected < 1 {
		tx.Rollback()
		return errno.ErrNotFound
	}
	if tx.Error != nil {
		tx.Rollback()
		logger.Logger.Error("article update failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase
	}
	// 更新 tags
	atm := &ArticleTagModel{}
	atmList := []ArticleTagModel{}
	err = tx.Model(atm).Where("article_id = ?", data.ID).Find(&atmList).Error
	if err != nil {
		tx.Rollback()
		logger.Logger.Error("article tags query failed", zap.String("error", err.Error()))
		return errno.InternalServerError
	}
	for _, atm := range atmList {
		exist := false
		// 找出不在的新的 tag 列表中的数据，然后删除掉
		for _, newTagId := range data.Tags {
			if newTagId == atm.TagID {
				exist = true
				break
			}
		}
		if !exist {
			err = tx.Model(atm).Unscoped().Where("id = ?", atm.ID).Delete(&atm).Error
			if err != nil {
				tx.Rollback()
				logger.Logger.Error("article tag delete failed", zap.String("error", err.Error()))
				return errno.InternalServerError
			}
		}
	}
	for _, newTagId := range am.Tags {
		exist := false
		for _, atm := range atmList {
			if newTagId == atm.TagID {
				exist = true
				break
			}
		}
		if !exist {
			newTag := &TagModel{}
			// 查找 tag 是否存在
			err = tx.Model(newTag).Where("user_id = ? and id = ?", userId, newTagId).Find(newTag).Error
			if err != nil {
				tx.Rollback()
				if err == gorm.ErrRecordNotFound {
					logger.Logger.Debug("article tag not found", zap.String("error", err.Error()))
					return errno.ErrNotFound
				}
				logger.Logger.Error("article tag query failed", zap.String("error", err.Error()))
				return errno.InternalServerError
			}
			// 插入新的 article tag model
			atm := &ArticleTagModel{
				TagID:     newTag.ID,
				ArticleID: data.ID,
			}
			err = tx.Model(atm).Create(atm).Error
			if err != nil {
				tx.Rollback()
				logger.Logger.Error("article tag insert failed", zap.String("error", err.Error()))
				return errno.InternalServerError
			}
		}
	}
	tx.Commit()
	return nil
}

func QueryArticleByID(articleId uint64) (error, *ArticleModel) {
	am := &ArticleModel{}
	data := DB.Self.Model(am).Preload("ThumbnailEntity").Preload("Author").Preload("AnnexsEntity").Preload("AnnexsEntity.Annex").Preload("TagsEntity").Preload("TagsEntity.Tag").Where("id=?", articleId).First(am)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, am
}

// QueryArticlesByTag
func QueryArticleListByTag(tagId uint64) (error, []ArticleModel) {
	err, list := QueryArticleTagListByTagId(tagId)
	if err != nil {
		return err, nil
	}
	articles := []ArticleModel{}
	for _, article := range list {
		articles = append(articles, article.Article)
	}
	return nil, articles
}

// QueryArticleListByCata query articles of catagory
func QueryArticleListByCata(cataId uint64) (error, []ArticleModel) {
	am := &ArticleModel{}
	list := []ArticleModel{}
	tx := DB.Self.Model(am).Where("category_id=?", cataId).Find(&list)
	if tx.Error != nil {
		logger.Logger.Debug("article list of catagory query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func QueryArticleStatisticByMonth() (error, []response.ArticleMonthStatisticResStruct) {
	am := ArticleModel{}
	list := []response.ArticleMonthStatisticResStruct{}
	tx := DB.Self.Raw(
		fmt.Sprintf(`
			SELECT
				COUNT( id ) AS count,
				DATE_FORMAT(
					created_at,
				?) AS created_month
			FROM
				%s
			GROUP BY
				created_month;
		`, am.TableName()),
		"%Y-%m",
	).Scan(&list)
	if tx.Error != nil {
		logger.Logger.Debug("article statistic of month query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func QueryArticleStatisticByMonthAndUser(userId uint64) (error, []response.ArticleMonthStatisticResStruct) {
	am := ArticleModel{}
	list := []response.ArticleMonthStatisticResStruct{}
	tx := DB.Self.Raw(
		fmt.Sprintf(`
			SELECT
				COUNT( id ) AS count,
				DATE_FORMAT(
					created_at,
				?) AS created_month,
				author_id
			FROM
				%s
			where author_id=?
			GROUP BY
				created_month;
		`, am.TableName()),
		"%Y-%m",
		userId,
	).Scan(&list)
	if tx.Error != nil {
		logger.Logger.Error("article statistic of month query failed: %s", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func ArticlePagination(data *request.PaginationReqStruct) (error, []ArticleReqStruct, uint) {
	pagesize := data.PageSize
	pageNumber := data.PageNumber
	if pagesize == 0 {
		pagesize = 10
	}
	offset := pagesize * pageNumber
	list := make([]ArticleReqStruct, pagesize)
	am := &ArticleModel{}
	var count uint
	tx := DB.Self.Model(am).Count(&count).Order("topping desc, created_at desc").Limit(pagesize).Offset(offset).Preload("ThumbnailEntity").Preload("Author").Preload("TagsEntity").Preload("TagsEntity.Tag").Find(&list)
	if tx.Error != nil {
		logger.Logger.Error("article pagination query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil, 0
	}
	return nil, list, count
}

func TagArticlePagination(tagId uint64, data *request.PaginationReqStruct) (error, []ArticleReqStruct, uint) {
	pagesize := data.PageSize
	pageNumber := data.PageNumber
	if pagesize == 0 {
		pagesize = 10
	}
	offset := pagesize * pageNumber
	list := make([]ArticleReqStruct, pagesize)
	am := &ArticleModel{}
	atm := &ArticleTagModel{}
	var count uint
	tx := DB.Self.Model(am).Where("id in (?)", DB.Self.Model(atm).Select("article_id").Where("tag_id = ?", tagId).SubQuery()).Count(&count).Order("topping desc, created_at desc").Limit(pagesize).Offset(offset).Preload("ThumbnailEntity").Preload("Author").Preload("TagsEntity").Preload("TagsEntity.Tag").Find(&list)
	if tx.Error != nil {
		logger.Logger.Error("article pagination query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil, 0
	}
	return nil, list, count
}
