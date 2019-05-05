package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_frame/config"
	"go_frame/cron"
	"go_frame/router"
)

func main() {
	if setting.Product == setting.Config.RunMode {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("gin.Version", gin.Version)
	}
	app := gin.Default()
	router.Router(app)
	cron.New().Start()
	_ = app.Run(":8080")
}
