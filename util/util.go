package util

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"
	"weshierNext/pkg/errno"
	"weshierNext/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// RandomString generate a random str
func RandomString(len int, useLowerCase bool) string {
	bytes := make([]byte, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		if useLowerCase && b%2 == 0 {
			b += 32
		}
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func UploadFile(c *gin.Context, savePath string, saveDir string, maxSize int64) (error, string, string, int64) {
	file, err := c.FormFile("file")
	if err != nil {
		return errno.ErrBind, "", "", 0
	}
	fileSize := file.Size
	if fileSize/1024 > maxSize {
		return errno.ErrFileSizeExceed, "", "", 0
	}

	_, err = os.Stat(savePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(savePath, 0777)
		if err != nil {
			return errno.InternalServerError, "", "", 0
		}
	}

	filename := RandomString(10, true) + "_" + file.Filename
	err = c.SaveUploadedFile(file, path.Join(savePath, filename))
	if err != nil {
		logger.Logger.Error("upload file failed", zap.String("error", err.Error()))
		return errno.InternalServerError, "", "", 0
	}
	publicPrefixer := viper.GetString("publicPrefix")
	if publicPrefixer == "" {
		publicPrefixer = "/public"
	}
	if publicPrefixer[0:1] != "/" {
		publicPrefixer = "/" + publicPrefixer
	}
	url := fmt.Sprintf("%s/%s/%s", publicPrefixer, saveDir, filename)
	return nil, url, filename, fileSize
}

func FindIndexOf(slice []interface{}, judge func(item ...interface{}) bool) int {
	for index, item := range slice {
		if judge(item) {
			return index
		}
	}
	return -1
}
