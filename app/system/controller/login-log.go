package controller

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/system/domain"
	"ruoyi-go/app/system/service"
	"ruoyi-go/common/model/request"
	resp "ruoyi-go/common/model/response"
	"ruoyi-go/common/response"
	"ruoyi-go/common/validate"
)

type LoginLog struct{}

var (
	loginLog = service.LoginLog{}
)

// List 获取访问日志列表
// @Summary 获取访问日志列表
// @Description 访问日志列表
// @Tags 访问日志管理
// @Param data query domain.LoginLogSearchRequest true "data"
// @Success 0 {object} response.Response{data=response.PageResult{list=[]domain.LoginLog}} "{"code": 0, "msg": "success","data": { "list": [],"total": 0 } }"
// @Router /system/login-log/list [get]
// @Security
func (u *LoginLog) List(c *gin.Context) {
	var searchParams domain.LoginLogSearchRequest
	// 获取查询参数
	if err := c.Bind(&searchParams); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	// 验证参数
	if err := validate.Struct(&searchParams); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	err, list, total := loginLog.GetList(&searchParams)
	if err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	response.OK(resp.PageResult{
		List:     list,
		Total:    total,
		Page:     searchParams.Page,
		PageSize: searchParams.PageSize,
	}, c)
}

// Delete 删除访问日志
// @Summary 删除访问日志
// @Description 删除访问日志
// @Tags 访问日志管理
// @Param ids body request.IdsRequest true "{ids: [1,2']}"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/login-log/delete [delete]
// @Security
func (u *LoginLog) Delete(c *gin.Context) {
	var data *request.IdsRequest
	// 获取参数
	if err := c.Bind(&data); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	if len(data.Ids) == 0 {
		response.Fail(-1, "参数不能为空", c)
		return
	}
	err := loginLog.Delete(data.Ids)
	if err != nil {
		response.Fail(-1, "删除失败: "+err.Error(), c)
		return
	}
	response.Success("删除成功", c)
}

// Detail 获取指定ID的访问日志详情
// @Summary 获取指定ID的访问日志详情
// @Description 获取指定ID的访问日志详情
// @Tags 访问日志管理
// @Param data query request.GetById true "data"
// @Success 0 {object} response.Response{data=domain.LoginLog} "{"code": 0, "msg":"success","data": {}}"
// @Router /system/login-log/detail [get]
// @Security
func (u *LoginLog) Detail(c *gin.Context) {
	var param request.GetById
	if err := c.Bind(&param); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	// 验证参数
	if err := validate.Struct(&param); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	record, err := loginLog.Detail(param.ID)
	if err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	response.OK(record, c)
}
