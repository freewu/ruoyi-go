package bootstrap

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/swagger"
	"ruoyi-go/app/system"
)

func Bootstrap(r *gin.Engine) error {
	// todo 后期可以做成配置反射调用
	system.Initialize(r) // 后台管理服务
	swagger.Initialize(r) // swagger 服务
	return  nil
}