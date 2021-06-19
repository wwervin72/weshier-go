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
