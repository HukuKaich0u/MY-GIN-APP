package main

import (
	"my-gin-app/controllers"
	"my-gin-app/database"
	"my-gin-app/middleware"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func main() {
	router := gin.Default()
	database.ConnectDB()

	router.Use(middleware.CustomRecovery())

	router.GET("/", func(c *gin.Context) {
		panic("意図的なパニック発生！")
	})

	// Cookieストアを作成。秘密鍵は任意の文字列に変更すること。
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Session("mysession", store))

	// 公開ルート
	router.POST("/register", controllers.RegisterHandler)
	router.POST("/login", controllers.LoginHandler)

	// 認証が必要なグループ
	auth := router.Group("/auth")
	auth.Use(middleware.AuthRequired())
	{
		auth.POST("/posts", controllers.CreatePostHandler)
		// 他の保護されたルートもここに追加
	}

	router.Run(":8080")
}
