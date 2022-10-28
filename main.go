package main

import (
	"github.com/afrimadhia/task-5-vix-btpns-afrima_dhia_defara/config"
	"github.com/afrimadhia/task-5-vix-btpns-afrima_dhia_defara/controllers"
	"github.com/afrimadhia/task-5-vix-btpns-afrima_dhia_defara/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	config.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.GET("/home", middleware.Authorization, controllers.Index)
	// r.GET("/", controllers.Index)
	r.GET("/login", controllers.LoginGet)
	r.POST("/login", controllers.LoginPost)
	r.GET("/logout", controllers.Logout)
	r.GET("/signup", controllers.SignupGet)
	r.POST("/signup", controllers.SignupPost)
	r.GET("/upload", middleware.Authorization, controllers.CreateFilesGet)
	r.POST("/upload", middleware.Authorization, controllers.CreateFilesPost)

	r.Run()
}
