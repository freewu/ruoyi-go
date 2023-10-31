package controller

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	group := r.Group("/system")
	{
		// 用户相关接口
		user := User{}
		group.GET("/user/list", user.List)
		group.POST("/user/create", user.Create)
		group.PUT("/user/update", user.Update)
		group.DELETE("/user/delete", user.Delete)
		group.GET("/user/detail", user.Detail)

		// 部门相关接口
		department := Department{}
		group.GET("/department/list", department.List)
		group.POST("/department/create", department.Create)
		group.PUT("/department/update", department.Update)
		group.DELETE("/department/delete", department.Delete)
		group.GET("/department/detail", department.Detail)

		// 职位相关接口
		post := Post{}
		group.GET("/post/list", post.List)
		group.POST("/post/create", post.Create)
		group.PUT("/post/update", post.Update)
		group.DELETE("/post/delete", post.Delete)
		group.GET("/post/detail", post.Detail)

		// 公告相关接口
		notice := Notice{}
		group.GET("/notice/list", notice.List)
		group.POST("/notice/create", notice.Create)
		group.PUT("/notice/update", notice.Update)
		group.DELETE("/notice/delete", notice.Delete)
		group.GET("/notice/detail", notice.Detail)

		// 系统参数相关接口
		config := Config{}
		group.GET("/config/list", config.List)
		group.POST("/config/create", config.Create)
		group.PUT("/config/update", config.Update)
		group.DELETE("/config/delete", config.Delete)
		group.GET("/config/detail", config.Detail)

		// 访问日志相关接口
		loginLog := LoginLog{}
		group.GET("/login-log/list", loginLog.List)
		group.DELETE("/login-log/delete", loginLog.Delete)
		group.GET("/login-log/detail", loginLog.Detail)

		// 角色相关接口

		//

	}
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
}
