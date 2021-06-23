package router

import (
	_ "weshierNext/docs"

	"fmt"
	"weshierNext/handler/article"
	"weshierNext/handler/category"
	"weshierNext/handler/comment"
	"weshierNext/handler/file"
	"weshierNext/handler/page"
	"weshierNext/handler/sd"
	"weshierNext/handler/tag"
	"weshierNext/handler/user"
	"weshierNext/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Load 加载路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(page.NotFound)
	// 配置 swagger
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", viper.GetString("port")))
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	g.GET("/", page.Home)

	api := g.Group("/api")
	{
		svcd := api.Group("/sd")
		{
			svcd.GET("/health", sd.HealthCheck)
			svcd.GET("/disk", sd.DiskCheck)
			svcd.GET("/cpu", sd.CPUCheck)
			svcd.GET("/ram", sd.RAMCheck)
		}
		userGroup := api.Group("/user")
		{
			userGroup.GET("", middleware.LoginRequired, user.QueryUserInfo)
			// github login
			userGroup.GET("/auth/github/callback", user.GithubLogin)
			userGroup.POST("/login", user.Login)
			userGroup.PUT("/changepwd", middleware.LoginRequired, user.ChangePwd)
			userGroup.GET("/logout", middleware.LoginRequired, user.Logout)
		}
		articleGroup := api.Group("/article")
		{
			articleGroup.POST("", middleware.LoginRequired, middleware.IsAdmin, article.Create)
			articleGroup.PUT("", middleware.LoginRequired, middleware.IsAdmin, article.Update)
			articleGroup.GET("/detail/:articleId", article.QueryArticleDetailByID)
			articleGroup.GET("/page", article.QueryArticleList)
			articleGroup.GET("/statistic/month", article.QueryArticleStatisticByMonth)
			articleGroup.GET("/statistic/month/:userId", article.QueryArticleStatisticByMonthAndUser)
		}
		commentGroup := api.Group("/comment")
		{
			commentGroup.POST("", middleware.LoginRequired, comment.LeaveComment)
			commentGroup.POST("/article", middleware.LoginRequired, comment.CommentArticle)
			commentGroup.GET("/article/:articleId/page", comment.ArticleCommentPagination)
			commentGroup.GET("/page", comment.CommentPagination)
		}
		tagGroup := api.Group("/tag")
		{
			tagGroup.POST("", middleware.LoginRequired, middleware.IsAdmin, tag.Create)
			tagGroup.PUT("", middleware.LoginRequired, middleware.IsAdmin, tag.Update)
			tagGroup.GET("/list", tag.QueryTagList)
			tagGroup.GET("/user/list", middleware.LoginRequired, tag.QueryUserTagList)
			tagGroup.DELETE("/:tagId", middleware.LoginRequired, middleware.IsAdmin, tag.DeleteTag)
			tagGroup.GET("/detail/:tagId", middleware.LoginRequired, middleware.IsAdmin, tag.QueryTagInfo)
			tagGroup.GET("/page", tag.QueryTagPagination)
			tagGroup.GET("/statistic/article", tag.QueryTagArticleStatistic)
			tagGroup.GET("/statistic/article/:userId", tag.QueryUserTagArticleStatistic)
		}
		cateGroup := api.Group("/cate")
		{
			cateGroup.POST("", middleware.LoginRequired, middleware.IsAdmin, category.Create)
			cateGroup.PUT("", middleware.LoginRequired, middleware.IsAdmin, category.Update)
			cateGroup.GET("/list", category.QueryCategoryList)
			cateGroup.GET("/user/list", middleware.LoginRequired, category.QueryUserCategoryList)
			cateGroup.DELETE("/:cateId", middleware.LoginRequired, middleware.IsAdmin, category.DeleteCategory)
			cateGroup.GET("/detail/:cateId", middleware.LoginRequired, middleware.IsAdmin, category.QueryCategoryInfo)
			cateGroup.GET("/page", category.QueryCatePagination)
		}
		fileGroup := api.Group("/file")
		{
			fileGroup.POST("/upload/img", middleware.LoginRequired, middleware.IsAdmin, file.UploadImg)
			fileGroup.POST("/upload/thumbnail", middleware.LoginRequired, middleware.IsAdmin, file.UploadThumbnail)
			fileGroup.POST("/upload/annex", middleware.LoginRequired, middleware.IsAdmin, file.UploadAnnex)
			fileGroup.GET("/thumbnails", middleware.LoginRequired, middleware.IsAdmin, file.QueryThumbnailList)
			fileGroup.DELETE("/thumbnails/:thumbnailId", middleware.LoginRequired, middleware.IsAdmin, file.DeleteThumbnail)
		}
	}
	return g
}
