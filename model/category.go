package model

import (
	"weshierNext/model/request"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"
	"weshierNext/pkg/validate"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// CategoryModel category model
type CategoryModel struct {
	BaseModel
	Name   string    `json:"name" gorm:"no null" binding:"required" validate:"min=1,max=10"`
	Desc   string    `json:"desc" validate:"max=100"`
	UserID uint64    `json:"userId" gorm:"no null;column:user_id"`
	User   UserModel `zh:"用户" json:"user" validate:"-" gorm:"foreignKey:UserID"`
}

// TableName category table name
func (cm *CategoryModel) TableName() string {
	return `ws_category`
}

// Create create category
func (cm *CategoryModel) Create() error {
	return DB.Self.Model(cm).Create(&cm).Error
}

// Delete delete category
func (cm *CategoryModel) Delete() error {
	// return DB.Self.Model(cm).Delete(&cm).Error
	tx := DB.Self.Model(cm).Begin()
	err := tx.Delete(&cm).Error
	if err != nil {
		tx.Rollback()
		logger.Logger.Error("catagory delete failed", zap.String("error", err.Error()))
		return errno.ErrDatabase
	}
	// 删除 cata 下的 article
	err, articles := QueryArticleListByCata(cm.ID)
	if err != nil {
		tx.Rollback()
		logger.Logger.Error("articles of catagory query failed", zap.String("error", err.Error()))
		return errno.ErrDatabase
	}
	for _, article := range articles {
		err = tx.Model(article).Delete(&article).Error
		if err != nil {
			tx.Rollback()
			logger.Logger.Error("article of catagory delete failed", zap.String("error", err.Error()))
			return errno.ErrDatabase
		}
	}
	return nil
}

// Validate Validate the field
func (cm *CategoryModel) Validate() error {
	return validate.Validate(*cm)
}

func CreateCategory(c *gin.Context, data *request.CreateCategoryReqStruct, userID uint64) (error, *CategoryModel) {
	category := &CategoryModel{
		Name:   data.Name,
		Desc:   data.Desc,
		UserID: userID,
	}
	err := category.Validate()
	if err != nil {
		logger.Logger.Debug("admin user validate failed", zap.String("error", err.Error()))
		return err, nil
	}
	err, existCate := QueryUserCategoryByName(userID, data.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return errno.ErrDatabase, nil
	}
	if existCate != nil {
		return &errno.Errno{
			Code:    errno.ErrValidation.Code,
			Message: "该分类名已存在，请修改",
		}, nil
	}
	err = category.Create()
	if err != nil {
		logger.Logger.Error("category insert failed", zap.String("error", err.Error()))
		return err, nil
	}
	return nil, category
}

func QueryCategoryByID(id uint64) (error, *CategoryModel) {
	cate := &CategoryModel{}
	data := DB.Self.Model(cate).Where("id=?", id).First(cate)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, cate
}

func QueryCategoryByIDAndUser(id uint64, userID uint64) (error, *CategoryModel) {
	cate := &CategoryModel{}
	data := DB.Self.Model(cate).Where("id = ? and user_id = ?", id, userID).First(cate)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, cate
}

func QueryCategoryByName(name string) (error, *CategoryModel) {
	cate := &CategoryModel{}
	data := DB.Self.Model(cate).Where("name=?", name).First(cate)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, cate
}

func QueryUserCategoryByName(userID uint64, name string) (error, *CategoryModel) {
	cate := &CategoryModel{}
	data := DB.Self.Model(cate).Where("name=? and user_id=?", name, userID).First(cate)
	if data.Error == gorm.ErrRecordNotFound {
		return data.Error, nil
	}
	return data.Error, cate
}

func CategoryPagination(data *request.PaginationReqStruct) (error, []CategoryModel, uint) {
	pagesize := data.PageSize
	pageNumber := data.PageNumber
	if pagesize == 0 {
		pagesize = 10
	}
	offset := pagesize * pageNumber
	list := make([]CategoryModel, pagesize)
	cm := &CategoryModel{}
	var count uint
	tx := DB.Self.Model(cm).Count(&count)
	if tx.Error != nil {
		logger.Logger.Error("category count query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, list, 0
	}
	tx = DB.Self.Model(cm).Limit(pagesize).Offset(offset).Find(&list)
	if tx.Error != nil {
		logger.Logger.Error("category pagination query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil, 0
	}
	return nil, list, count
}

func QueryCategoryList() (error, []CategoryModel) {
	cm := &CategoryModel{}
	list := []CategoryModel{}
	tx := DB.Self.Model(cm).Find(&list)
	if tx.Error != nil {
		logger.Logger.Debug("category list query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func QueryUserCategoryList(userId uint64) (error, []CategoryModel) {
	cm := &CategoryModel{}
	list := []CategoryModel{}
	tx := DB.Self.Model(cm).Where("user_id = ?", userId).Find(&list)
	if tx.Error != nil {
		logger.Logger.Debug("category list query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil
	}
	return nil, list
}

func UpdateCategory(data *request.UpdateCategoryReqStruct, userID uint64) error {
	cate := &CategoryModel{
		Name: data.Name,
		Desc: data.Desc,
	}
	err := cate.Validate()
	if err != nil {
		logger.Logger.Debug("cate validate failed", zap.String("error", err.Error()))
		return err
	}
	err, existTag := QueryUserCategoryByName(userID, data.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return errno.ErrDatabase
	}
	if existTag != nil {
		return &errno.Errno{
			Code:    errno.ErrValidation.Code,
			Message: "该分类名已存在，请修改",
		}
	}
	aMap := make(map[string]interface{})
	aMap["name"] = data.Name
	aMap["desc"] = data.Desc
	err = DB.Self.Model(&cate).Select("name", "desc").Where("user_id=? and id=?", userID, data.ID).Updates(aMap).Error
	if err != nil {
		logger.Logger.Error("cate update failed", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func CataPagination(data *request.PaginationReqStruct) (error, []CategoryModel, uint) {
	pagesize := data.PageSize
	pageNumber := data.PageNumber
	if pagesize == 0 {
		pagesize = 10
	}
	offset := pagesize * pageNumber
	list := make([]CategoryModel, pagesize)
	cm := &CategoryModel{}
	var count uint
	tx := DB.Self.Model(cm).Count(&count).Limit(pagesize).Offset(offset).Find(&list)
	if tx.Error != nil {
		logger.Logger.Error("catagory pagination query failed", zap.String("error", tx.Error.Error()))
		return errno.ErrDatabase, nil, 0
	}
	DB.Self.Model(cm).Preload("User").Find(&list)
	return nil, list, count
}
