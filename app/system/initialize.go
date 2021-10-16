package system

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/system/controller"
)

func Initialize(r *gin.Engine)  {
	// 注册路由
	controller.InitRouter(r)
	//
}
