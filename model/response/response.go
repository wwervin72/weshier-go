package response

import "weshierNext/model/request"

// Response 服务返回消息体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ObjectList []interface{}

type PaginationDataStruct struct {
	request.PaginationReqStruct
	List  interface{} `json:"list"`
	Total uint        `json:"total"`
}

// CreateArticleResStruct create article response body
type CreateArticleResStruct struct {
}

type DropdownListResStruct struct {
	Label string `json:"label"`
	Value uint64 `json:"value"`
}

type FileListResStruct struct {
	DropdownListResStruct
	Url string `json:"url"`
}

// CreateArticleResStruct create article response body
type CreateCategoryResStruct struct {
}

type UploadImageResStruct struct {
	DropdownListResStruct
	Url string `json:"url"`
}

type UploadAnnexResStruct struct {
	ID    uint64 `json:"id"`
	Value uint64 `json:"value"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Size  int64  `json:"size"`
}

type PaginationResStruct struct {
	Response
	Data PaginationDataStruct `json:"data"`
}

// CreateArticleResStruct create article response body
type CreateTagResStruct struct {
}

type ArticleMonthStatisticResStruct struct {
	Count        uint64 `json:"count"`
	AuthorId     uint64 `json:"authorId"`
	CreatedMonth string `json:"createdMonth"`
}
type ArticleTagStatisticResStruct struct {
	Tag     string `json:"tag"`
	TagID   string `json:"tagId"`
	TagName string `json:"tagName"`
	TagDesc string `json:"tagDesc"`
	Count   uint64 `json:"count"`
}
