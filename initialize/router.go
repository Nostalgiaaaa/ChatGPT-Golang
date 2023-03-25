package initialize

import (
	"chatgpt-go/html"
	"chatgpt-go/middleware"
	"chatgpt-go/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {

	chatData, err := routes.NewChatStorage()
	if err != nil {
		panic(err)
	}
	//defer chatData.Close()

	r := gin.Default()
	r.Use(middleware.SetAuthorizationHeader())
	//r.Use(middleware.AuthMiddleware())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization") // 确保允许"Authorization"请求头
	r.Use(cors.New(corsConfig))

	api := r.Group("api")
	{
		api.POST("/chat-process", routes.ChatProcess(chatData))
		api.POST("/config", routes.GetConfig)
		api.POST("/session", routes.SessionEndpoint)
		api.POST("/verify", routes.VerifyEndpoint)
	}

	r.StaticFS("/", http.FS(html.Static))

	return r
}
