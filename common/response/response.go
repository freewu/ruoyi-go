package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code	int 		`json:"code" example:"0"` // 代码
	Data	interface{} `json:"data"`// 数据集
	Msg 	string 		`json:"msg"` // 消息
}

var response *Response

func (r *Response) Error(code int,msg string,c *gin.Context) {
	response = &Response{}
	response.Code = code
	response.Msg = msg
	c.JSON(http.StatusInternalServerError,response)
}

func (r *Response) Success(data interface{},c *gin.Context) {
	response = &Response{}
	response.Code = 0
	response.Msg = "success"
	response.Data = data
	c.JSON(http.StatusOK,response)
}

func Fail(code int,msg string,r *gin.Context) {
	response.Error(code,msg,r)
}

func OK(data interface{},r *gin.Context) {
	response.Success(data,r)
}

func Success(msg string,r *gin.Context) {
	response.Error(0,msg,r)
}