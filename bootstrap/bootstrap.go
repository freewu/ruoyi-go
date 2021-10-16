package bootstrap

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/system"
)

func Bootstrap(r *gin.Engine) error {
	system.Initialize(r)
	return  nil
}