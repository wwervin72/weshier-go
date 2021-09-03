package request

// CreateArticleReqStruct create article request body
type CreateArticleReqStruct struct {
	Title        string   `json:"title" binding:"required" validate:"min=2,max=50"`
	Abstract     string   `json:"abstract" binding:"required" validate:"min=1,max=200"`
	Content      string   `json:"content" binding:"required" validate:"min=1"`
	AllowComment int8     `json:"allowComment" validate:"min=1,max=1"`
	Password     string   `json:"password" validate:"max=10"`
	Thumbnail    string   `json:"thumbnail"`
	Topping      int8     `json:"topping"`
	CategoryID   uint64   `json:"categoryId" binding:"required"`
	Tags         []uint64 `json:"tags"`
	Annexs       []uint64 `json:"annexs"`
}

type UpdateArticleReqStruct struct {
	ID uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT;"`
	CreateArticleReqStruct
}

type PaginationReqStruct struct {
	PageSize   uint8 `json:"pageSize" form:"pageSize"`
	PageNumber uint8 `json:"pageNumber" form:"pageNumber"`
}

// CreateArticleReqStruct create article request body
type CreateTagReqStruct struct {
	Name string `json:"name" binding:"required" validate:"min=3,max=10"`
	Desc string `json:"desc" validate:"max=100"`
}

type UpdateTagReqStruct struct {
	ID uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT;"`
	CreateTagReqStruct
}

// CreateArticleReqStruct create article request body
type CreateCategoryReqStruct struct {
	Name string `json:"name" binding:"required" validate:"min=3,max=10"`
	Desc string `json:"desc" validate:"max=100"`
}

type UpdateCategoryReqStruct struct {
	ID uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT;"`
	CreateCategoryReqStruct
}

type CreateCommentReqStruct struct {
	CommentID uint64 `json:"commentId"`
	Content   string `json:"content" binding:"required" validate:"min=1,max=250"`
}

type DelCommentReqStruct struct {
	CommentID uint64 `json:"commentId"`
}

type CreateArticleCommentReqStruct struct {
	CreateCommentReqStruct
	ArticleID uint64 `json:"articleId" binding:"required"`
}

type ChangeAccountPwdReqStruct struct {
	OldPwd string `json:"oldPwd" binding:"required"`
	NewPwd string `json:"newPwd" binding:"required" validate:"min=5,max=50"`
}
