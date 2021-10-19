package system

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/system/controller"
	"ruoyi-go/app/system/domain"
)

func Initialize(r *gin.Engine)  {
	// 注册路由
	controller.InitRouter(r)
	// 注册DB
	domain.InitDB(r)
}
