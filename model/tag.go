package model

import (
	"encoding/json"
	"fmt"
	"weshierNext/model/request"
	"weshierNext/model/response"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"
	"weshierNext/pkg/validate"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// TagModel tag model
type TagModel struct {
	BaseModel
	Name   string `json:"name" gorm:"no null" binding:"required" validate:"min=1,max=20"`
	Desc   string `json:"desc" validate:"max=100"`
	UserID uint64 `json:"userId" gorm:"no null;column:user_id"`
}

// TableName tag table name
func (tm *TagModel) TableName() string {
	return "ws_tag"
}

// Create create article
func (tm *TagModel) Create() error {
	return DB.Self.Model(tm).Create(&tm).Error
}

// Delete delete article
func (tm *TagModel) Delete() error {
	tx := DB.Self.Model(tm).Begin()
	err := tx.Delete(&tm).Error
	if err != nil {
		tx.Rollback()
		logger.Logger.Error("tag delete failed", zap.String("error", err.Error()))
		return errno.ErrDatabase
	}
	// 删除 tag 下的 article
	err, articles := QueryArticleListByTag(tm.ID)
	if err != nil {
		tx.Rollback()
		logger.Logger.Error("articles of tag query failed", zap.String("error", err.Error()))
		return errno.ErrDatabase
	}
	for _, article := range articles {
		err = tx.Model(article).Delete(&article).Error
		if err != nil {
			tx.Rollback()
			logger.Logger.Error("article of tag delete failed", zap.String("error", err.Error()))
			return errno.ErrDatabase
		}
	}
	return nil
}

// Validate Validate the field
func (tm *TagModel) Validate() error {
	return validate.Validate(*tm)
}

func CreateTag(data *request.CreateTagReqStruct, userID uint64) (error, *TagModel) {
	tag := &TagModel{
		Name:   data.Name,
		Desc:   data.Desc,
		UserID: userID,
	}
	err := tag.Validate()
	if err != nil {
		logger.Logger.Debug("tag validate failed", zap.String("error", err.Error()))
		return err, nil
	}
	err, existTag := QueryUserTagByName(userID, data.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return errno.ErrDatabase, nil
	}
	if existTag != nil {
		return &errno.Errno{
			Code:    errno.ErrValidation.Code,
			Message: "该标签名已存在，请修改",
		}, nil
	}
	err = tag.Create()
	if err != nil {
		logger.Logger.Error("tag insert failed", zap.String("error", err.Error()))
		return err, nil
	}
	return nil, tag
}

func UpdateTag(data *request.UpdateTagReqStruct, userID uint64) error {
	tag := &TagModel{
		Name: data.Name,
		Desc: data.Desc,
	}
	err := tag.Validate()
	if err != nil {
		logger.Logger.Debug("tag validate failed", zap.String("error", err.Error()))
		return err
	}
	err, existTag := QueryUserTagByName(userID, data.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return errno.ErrDatabase
	}
	if existTag != nil {
		return &errno.Errno{
			Code:    errno.ErrValidation.Code,
			Message: "该标签名已存在，请修改",
		}
	}
	j, err := json.Marshal(tag)
	if err != nil {
		return errno.InternalServerError
	}
	tMap := make(map[string]interface{})
	err = json.Unmarshal(j, &tMap)
	if err != nil {
		return errno.InternalServerError
	}
	err = DB.Self.Model(&tag).Select("name", "desc").Where("user_id = ? and id = ?", userID, data.ID).Updates(tMap).Error
	if err != nil {
		logger.Logger.Error("tag update failed", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func QueryTagByID(tagId uint64) (error, *TagModel) {
	tag := &TagModel{}
	data := DB.Self.Model(tag).Where("id=?", tagId).First(tag)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, tag
}

func QueryTagByIDAndUser(tagId uint64, userId uint64) (error, *TagModel) {
	tag := &TagModel{}
	data := DB.Self.Model(tag).Where("id = ? and user_id = ?", tagId, userId).First(tag)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, tag
}

func QueryTagByName(tagName string) (error, *TagModel) {
	tag := &TagModel{}
	data := DB.Self.Model(tag).Where("name=?", tagName).First(tag)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, tag
}

func QueryUserTagByName(userId uint64, tagName string) (error, *TagModel) {
	tag := &TagModel{}
	data := DB.Self.Model(tag).Where("name=? and user_id=?", tagName, userId).First(tag)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, tag
}

func QueryTagList() (error, []TagModel) {
	tm := &TagModel{}
	list := []TagModel{}
	tx := DB.Self.Model(tm).Find(&list)
	if tx.Error != nil {
		logger.Logger.Debug("tag list query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func QueryUserTagList(userId uint64) (error, []TagModel) {
	tm := &TagModel{}
	list := []TagModel{}
	tx := DB.Self.Model(tm).Where("user_id = ?", userId).Find(&list)
	if tx.Error != nil {
		logger.Logger.Debug("tag list query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func QueryTagArticleStatistic() (error, []response.ArticleTagStatisticResStruct) {
	list := []response.ArticleTagStatisticResStruct{}
	tm := TagModel{}
	atm := ArticleTagModel{}
	tx := DB.Self.Raw(
		fmt.Sprintf(`
			SELECT
				COUNT(b.article_id) AS count,
				a.id AS tag_id,
				a.NAME AS tag_name,
				a.DESC AS tag_desc
			FROM
				%s AS a
				LEFT JOIN %s AS b ON a.id = b.tag_id
			GROUP BY
				a.id
			HAVING count>0;
			`,
			tm.TableName(),
			atm.TableName(),
		),
	).Scan(&list)
	if tx.Error != nil {
		logger.Logger.Debug("tag list count query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func QueryUserTagArticleStatistic(userId uint64) (error, []response.ArticleTagStatisticResStruct) {
	list := []response.ArticleTagStatisticResStruct{}
	tm := TagModel{}
	atm := ArticleTagModel{}
	tx := DB.Self.Raw(
		fmt.Sprintf(`
			SELECT
				COUNT( b.article_id ) AS count,
				a.id AS tag_id,
				a.NAME AS tag_name,
				a.DESC AS tag_desc,
				a.user_id AS tag_user
			FROM
				%s AS a
				LEFT JOIN %s AS b ON a.id = b.tag_id
			GROUP BY
				a.id
			HAVING
				count > 0
				AND tag_user=?;
			`,
			tm.TableName(),
			atm.TableName(),
		),
		userId,
	).Scan(&list)
	if tx.Error != nil {
		logger.Logger.Debug("tag list count query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func TagPagination(data *request.PaginationReqStruct) (error, []TagModel, uint) {
	pagesize := data.PageSize
	pageNumber := data.PageNumber
	if pagesize == 0 {
		pagesize = 10
	}
	offset := pagesize * pageNumber
	list := make([]TagModel, pagesize)
	tm := &TagModel{}
	var count uint
	tx := DB.Self.Model(tm).Count(&count)
	if tx.Error != nil {
		logger.Logger.Error("tag count query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, list, 0
	}
	tx = DB.Self.Model(tm).Limit(pagesize).Offset(offset).Find(&list)
	if tx.Error != nil {
		logger.Logger.Error("tag pagination query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil, 0
	}
	return nil, list, count
}
