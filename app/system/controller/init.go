package controller

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine)  {
	group := r.Group("")
	{
		user := User{}
		group.GET("/system/user/list",user.List)
	}
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
}
