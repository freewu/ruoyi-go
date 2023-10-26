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

type Notice struct{}

var (
	notice = service.Notice{}
)

// List 获取公告列表
// @Summary 获取公告列表
// @Description 公告列表
// @Tags 公告管理
// @Param data query domain.NoticeSearchRequest true "data"
// @Success 0 {object} response.Response{data=response.PageResult{list=[]domain.Notice}} "{"code": 0, "msg": "success","data": { "list": [],"total": 0 } }"
// @Router /system/notice/list [get]
// @Security
func (u *Notice) List(c *gin.Context) {
	var searchParams domain.NoticeSearchRequest
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
	err, list, total := notice.GetList(&searchParams)
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

// Create 添加公告
// @Summary 添加公告
// @Description 添加公告
// @Tags 公告管理
// @Param data body domain.NoticeCreateRequest true "data"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/notice/create [notice]
// @Security
func (u *Notice) Create(c *gin.Context) {
	var data *domain.NoticeCreateRequest
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
	if err := notice.Create(data); err != nil {
		response.Fail(-1, "添加失败: "+err.Error(), c)
	} else {
		response.Success("添加成功", c)
	}
}

// Update 修改公告
// @Summary 修改公告
// @Description 修改公告
// @Tags 公告管理
// @Param data body domain.NoticeUpdateRequest true "data"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/notice/update [put]
// @Security
func (u *Notice) Update(c *gin.Context) {
	var data *domain.NoticeUpdateRequest
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
	if err := notice.Update(data); err != nil {
		response.Fail(-1, "修改失败: "+err.Error(), c)
	} else {
		response.Success("修改成功", c)
	}
}

// Delete 删除公告
// @Summary 删除公告
// @Description 删除公告
// @Tags 公告管理
// @Param ids body request.IdsRequest true "{ids: [1,2']}"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/notice/delete [delete]
// @Security
func (u *Notice) Delete(c *gin.Context) {
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
	err := notice.Delete(data.Ids)
	if err != nil {
		response.Fail(-1, "删除失败: "+err.Error(), c)
		return
	}
	response.Success("删除成功", c)
}

// Detail 获取指定ID的公告详情
// @Summary 获取指定ID的公告详情
// @Description 获取指定ID的公告详情
// @Tags 公告管理
// @Param data query request.GetById true "data"
// @Success 0 {object} response.Response{data=domain.Notice} "{"code": 0, "msg":"success","data": {}}"
// @Router /system/notice/detail [get]
// @Security
func (u *Notice) Detail(c *gin.Context) {
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
	record, err := notice.Detail(param.ID)
	if err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	response.OK(record, c)
}
