package file

import (
	"fmt"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func QueryThumbnailList(c *gin.Context) {
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err, list := model.QueryThumbnailsByUser(user.ID)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	res := []response.FileListResStruct{}
	for _, item := range *list {
		res = append(res, response.FileListResStruct{
			DropdownListResStruct: response.DropdownListResStruct{
				Label: item.Name,
				Value: item.ID,
			},
			Url: fmt.Sprintf("%s/%s/%s", viper.GetString("publicPrefix"), model.ThumbnailDir, item.Name),
		})
	}
	handler.SendResponse(c, nil, res)
	return
}
