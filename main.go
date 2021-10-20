package main

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/bootstrap"
	"ruoyi-go/common/config"
)

// @title ruoyi-go API文档
// @version 1.0
// @description ruoyi-go 在线API文档
// @host localhost
// @BasePath /
func main() {
	if config.Bool("service.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	_ = bootstrap.Bootstrap(r)
	_ = r.Run(config.String("service.port")) // listen and serve on 0.0.0.0:8080
}