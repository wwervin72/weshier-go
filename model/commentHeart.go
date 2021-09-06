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

// Create heart comment
func (chm *CommentHeartModel) Create() error {
	return DB.Self.Model(chm).Create(&chm).Error
}

// Delete cancel heart comment
func (chm *CommentHeartModel) Delete() error {
	return DB.Self.Model(chm).Delete(&chm).Error
}

func CountCommentHeartByIdAndUser(commentId uint64, userId uint64) (count uint16, err error) {
	chm := &CommentHeartModel{}
	tx := DB.Self.Model(chm).Where("comment_id = ? and user_id = ?", commentId, userId).Count(&count)
	return count, tx.Error
}

func CountCommentHeartById(commentId uint64) (count uint, err error) {
	chm := &CommentHeartModel{}
	tx := DB.Self.Model(chm).Where("comment_id = ?", commentId).Count(&count)
	return count, tx.Error
}

func QueryCommentHeartUserById(commentId uint64) (userIds []uint64, err error) {
	list := &[]CommentHeartModel{}
	tx := DB.Self.Model(&CommentHeartModel{}).Select("user_id").Where("comment_id = ?", commentId).Find(list)
	for _, item := range *list {
		userIds = append(userIds, item.UserID)
	}
	return userIds, tx.Error
}

func CountCommentHeartByUser(userId uint64) (count uint16, err error) {
	chm := &CommentHeartModel{}
	tx := DB.Self.Model(chm).Where("user_id = ?", userId).Count(&count)
	return count, tx.Error
}
