package swagger

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"net/http"
)

func Initialize(r *gin.Engine)  {
	// 注册路由
	r.GET("/swagger.json", func(c *gin.Context) {
		jsonStr, err := swag.ReadDoc()
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		record := make(map[string]interface{})
		if err := json.Unmarshal([]byte(jsonStr), &record); err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, record)
	})
}
