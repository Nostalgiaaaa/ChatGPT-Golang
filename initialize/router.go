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

	r.POST("/api/chat-process", routes.ChatProcess(chatData))
	r.POST("/api/config", routes.GetConfig)
	r.POST("/api/session", routes.Session)
	r.POST("/api/verify", routes.Verify)

	r.StaticFS("/", http.FS(html.Static))

	return r
}
