package model

// CommentModel article model
type ArticleHeartModel struct {
	BaseModel
	UserID    uint64       `zh:"用户" json:"userId" gorm:"not null;column:user_id" validate:"required"`
	User      UserModel    `zh:"用户" json:"user" validate:"-" gorm:"foreignKey:UserID"`
	ArticleID uint64       `zh:"文章" json:"articleId" gorm:"column:article_id;"`
	Article   ArticleModel `json:"article" validate:"-" gorm:"foreignKey:ArticleID"`
}

// TableName comment table name
func (ahm *ArticleHeartModel) TableName() string {
	return "ws_article_heart"
}

// Create create article
func (ahm *ArticleHeartModel) Create() error {
	return DB.Self.Model(ahm).Create(&ahm).Error
}

// Delete delete article
func (ahm *ArticleHeartModel) Delete() error {
	return DB.Self.Model(ahm).Delete(&ahm).Error
}

func CountArticleHeartByIdAndUser(commentId uint64, userId uint64) (count uint16, err error) {
	ahm := &ArticleHeartModel{}
	tx := DB.Self.Model(ahm).Where("article_id = ? and user_id = ?", commentId, userId).Count(&count)
	return count, tx.Error
}

func CountArticleHeartById(commentId uint64) (count uint, err error) {
	ahm := &ArticleHeartModel{}
	tx := DB.Self.Model(ahm).Where("article_id = ?", commentId).Count(&count)
	return count, tx.Error
}

func QueryArticleHeartUserById(commentId uint64) (userIds []uint64, err error) {
	list := &[]ArticleHeartModel{}
	tx := DB.Self.Model(&ArticleHeartModel{}).Select("user_id").Where("article_id = ?", commentId).Find(list)
	for _, item := range *list {
		userIds = append(userIds, item.UserID)
	}
	return userIds, tx.Error
}

func CountArticleHeartByUser(userId uint64) (count uint16, err error) {
	ahm := &ArticleHeartModel{}
	tx := DB.Self.Model(ahm).Where("user_id = ?", userId).Count(&count)
	return count, tx.Error
}
