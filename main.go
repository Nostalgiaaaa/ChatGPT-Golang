package main

import (
	"chatgpt-go/chatgpt"
	"chatgpt-go/core"
	"chatgpt-go/global"
	"chatgpt-go/html"
	"chatgpt-go/middleware"
	"chatgpt-go/routes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//global.OpenAIKey = os.Getenv("OPENAI_API_KEY")
	//
	//if global.OpenAIKey == "" {
	//	if terminal.IsTerminal(int(os.Stdin.Fd())) {
	//		reader := bufio.NewReader(os.Stdin)
	//		fmt.Print("Enter your OpenAI API Key: ")
	//		var err error
	//		global.OpenAIKey, err = reader.ReadString('\n')
	//		global.OpenAIKey = strings.TrimSpace(global.OpenAIKey) // 移除末尾的换行符
	//		if err != nil {
	//			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	//			os.Exit(1)
	//		}
	//	} else {
	//		fmt.Println("OPENAI_API_KEY is not provided and terminal non-interactive. Exiting.")
	//		os.Exit(1)
	//	}
	//}

	global.ViperConfig = core.Viper()

	fmt.Println(global.Config)

	chatStorage, err := chatgpt.NewChatStorage()
	if err != nil {
		panic(err)
	}
	defer chatStorage.Close()

	r := gin.Default()
	r.Use(middleware.SetAuthorizationHeader())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization") // 确保允许"Authorization"请求头
	r.Use(cors.New(corsConfig))

	r.POST("/chat-process", routes.ChatProcess(chatStorage))
	r.POST("/config", routes.GetConfig)
	r.POST("/session", routes.Session)
	r.POST("/verify", routes.Verify)

	r.StaticFS("/", http.FS(html.Static))

	port := ":3002"
	fmt.Printf("Server is running on http://localhost:%d\n", port)

	r.Run(port)

}
