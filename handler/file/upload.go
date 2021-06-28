package file

import (
	"path"
	"weshierNext/handler"
	"weshierNext/model"
	"weshierNext/model/response"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"
	"weshierNext/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// UploadImg
// @Summary 文件上传
// @Description 文件上传
// @Tags file
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} handler.Response{data=response.UploadImageResStruct} "图片上传成功"
// @Router /api/file/upload/img [post]
func UploadImg(c *gin.Context) {
	savePath := path.Join(viper.GetString("upload.saveDir"), model.ImageDir)
	err, url, _, _ := util.UploadFile(c, savePath, model.ImageDir, 100)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, &response.UploadImageResStruct{
		Url: url,
	})
}

// UploadThumbnail
// @Summary 缩略图上传
// @Description 缩略图上传
// @Tags file
// @Produce json
// @Param file formData file true "缩略图"
// @Success 200 {object} handler.Response{data=response.UploadImageResStruct} "缩略图上传成功"
// @Router /api/file/upload/thumbnail [post]
func UploadThumbnail(c *gin.Context) {
	savePath := path.Join(viper.GetString("upload.saveDir"), model.ThumbnailDir)
	err, url, filename, size := util.UploadFile(c, savePath, model.ThumbnailDir, 100)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	tm := &model.ThumbnailModel{
		Name:   filename,
		UserID: user.ID,
		Size:   size,
	}
	err = tm.Create()
	if err != nil {
		logger.Logger.Error("upload thumbnail failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, &response.UploadImageResStruct{
		Url: url,
		DropdownListResStruct: response.DropdownListResStruct{
			Label: tm.Name,
			Value: tm.ID,
		},
	})
}

// UploadAnnex
// @Summary 附件上传
// @Description 附件上传
// @Tags file
// @Produce json
// @Param annex formData file true "附件"
// @Success 200 {object} handler.Response{data=response.UploadAnnexResStruct} "附件上传成功"
// @Router /api/file/upload/annex [post]
func UploadAnnex(c *gin.Context) {
	savePath := path.Join(viper.GetString("upload.saveDir"), model.AnnexDir)
	err, url, filename, size := util.UploadFile(c, savePath, model.AnnexDir, 100)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	err, user := handler.GetUserFromContext(c)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	am := &model.AnnexModel{
		Name:   filename,
		UserID: user.ID,
		Size:   size,
	}
	err = am.Create()
	if err != nil {
		logger.Logger.Error("upload thumbnail failed", zap.String("error", err.Error()))
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, &response.UploadAnnexResStruct{
		Url:   url,
		Name:  filename,
		Value: am.ID,
		ID:    am.ID,
	})
	return
}
