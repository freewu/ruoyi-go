package main

import "github.com/gin-gonic/gin"

// @title rouyi-go API文档
// @version 1.0
// @description rouyi-go 在线API文档
// @host localhost
// @BasePath /
func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run() // listen and serve on 0.0.0.0:8080
}