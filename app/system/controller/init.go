package controller

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	group := r.Group("/system")
	{
		// 用户相关接口
		user := User{}
		group.GET("/user/list",user.List)
		group.POST("/user/create",user.Create)
		group.PUT("/user/update",user.Update)
		group.DELETE("/user/delete",user.Delete)
		group.GET("/user/detail",user.Detail)

		// 部门相关接口
		department := Department{}
		group.GET("/department/list",department.List)
		group.POST("/department/create",department.Create)
		group.PUT("/department/update",department.Update)
		group.DELETE("/department/delete",department.Delete)
		group.GET("/department/detail",department.Detail)

		// 职位相关接口

		// 角色相关接口

		//

	}
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
}
