package model

// CommentModel article model
type CommentHeartModel struct {
	BaseModel
	UserID    uint64       `zh:"用户" json:"userId" gorm:"not null;column:user_id" validate:"required"`
	User      UserModel    `zh:"用户" json:"user" validate:"-" gorm:"foreignKey:UserID"`
	CommentID uint64       `zh:"评论" json:"commentId" gorm:"column:comment_id;"`
	Comment   CommentModel `json:"comment" validate:"-" gorm:"foreignKey:CommentID"`
}

// TableName comment heart table name
func (chm *CommentHeartModel) TableName() string {
	return "ws_comment_heart"
}
