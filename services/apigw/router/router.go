package router

import (
	"github.com/gin-gonic/gin"
	"go-disk/midware"
	"go-disk/services/apigw/handler"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/static", "./static")
	router.StaticFile("/hom", "./static/view/home.html")

	router.Use(midware.Cors())

	//create and init group
	userGroup := router.Group("/users")
	uploadGroup := router.Group("/files/upload")

	userServiceRoute(userGroup)
	uploadServiceRoute(uploadGroup)

	return router
}

func uploadServiceRoute(group *gin.RouterGroup) {
	group.StaticFile("/", "./static/view/index.html")
	group.GET("/endpoint", handler.GetUploadServiceEndpoint())
}

func userServiceRoute(group *gin.RouterGroup) {
	group.POST("/register", handler.RegisterUser())
	group.POST("/login", handler.UserLogin())

	group.Use(midware.AuthorizeInterceptor())
	group.GET("/info", handler.QueryUserInfo())
}
