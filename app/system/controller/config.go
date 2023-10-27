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

type Config struct{}

var (
	config = service.Config{}
)

// List 获取系统参数列表
// @Summary 获取系统参数列表
// @Description 系统参数列表
// @Tags 系统参数管理
// @Param data query domain.ConfigSearchRequest true "data"
// @Success 0 {object} response.Response{data=response.PageResult{list=[]domain.Config}} "{"code": 0, "msg": "success","data": { "list": [],"total": 0 } }"
// @Router /system/config/list [get]
// @Security
func (u *Config) List(c *gin.Context) {
	var searchParams domain.ConfigSearchRequest
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
	err, list, total := config.GetList(&searchParams)
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

// Create 添加系统参数
// @Summary 添加系统参数
// @Description 添加系统参数
// @Tags 系统参数管理
// @Param data body domain.ConfigCreateRequest true "data"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/config/create [config]
// @Security
func (u *Config) Create(c *gin.Context) {
	var data *domain.ConfigCreateRequest
	// 获取参数
	if err := c.Bind(&data); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	// 验证参数
	if err := validate.Struct(data); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	if err := config.Create(data); err != nil {
		response.Fail(-1, "添加失败: "+err.Error(), c)
	} else {
		response.Success("添加成功", c)
	}
}

// Update 修改系统参数
// @Summary 修改系统参数
// @Description 修改系统参数
// @Tags 系统参数管理
// @Param data body domain.ConfigUpdateRequest true "data"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/config/update [put]
// @Security
func (u *Config) Update(c *gin.Context) {
	var data *domain.ConfigUpdateRequest
	// 获取参数
	if err := c.Bind(&data); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	// 参数验证
	if err := validate.Struct(data); err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	if err := config.Update(data); err != nil {
		response.Fail(-1, "修改失败: "+err.Error(), c)
	} else {
		response.Success("修改成功", c)
	}
}

// Delete 删除系统参数
// @Summary 删除系统参数
// @Description 删除系统参数
// @Tags 系统参数管理
// @Param ids body request.IdsRequest true "{ids: [1,2']}"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/config/delete [delete]
// @Security
func (u *Config) Delete(c *gin.Context) {
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
	err := config.Delete(data.Ids)
	if err != nil {
		response.Fail(-1, "删除失败: "+err.Error(), c)
		return
	}
	response.Success("删除成功", c)
}

// Detail 获取指定ID的系统参数详情
// @Summary 获取指定ID的系统参数详情
// @Description 获取指定ID的系统参数详情
// @Tags 系统参数管理
// @Param data query request.GetById true "data"
// @Success 0 {object} response.Response{data=domain.Config} "{"code": 0, "msg":"success","data": {}}"
// @Router /system/config/detail [get]
// @Security
func (u *Config) Detail(c *gin.Context) {
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
	record, err := config.Detail(param.ID)
	if err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	response.OK(record, c)
}
