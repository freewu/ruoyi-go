package controller

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/system/domain"
	"ruoyi-go/app/system/service"
	"ruoyi-go/common/model/request"
	resp "ruoyi-go/common/model/response"
	"ruoyi-go/common/response"
	"ruoyi-go/common/validate"

	log "github.com/sirupsen/logrus"
)

type Department struct {}

var (
	department = service.Department{}
)

// @Summary 获取部门列表
// @Description 部门列表
// @Tags 部门管理
// @Param data query domain.DepartmentSearchRequest true "data"
// @Success 0 {object} response.Response{data=response.PageResult{list=[]domain.Department}} "{"code": 0, "data": { "list": [] } }"
// @Router /system/department/list [get]
// @Security
func (u *Department) List(c *gin.Context) {
	var searchParams domain.DepartmentSearchRequest
	// 获取查询参数
	if err := c.Bind(&searchParams); err != nil {
		response.Fail(-1, err.Error(),c)
	}
	log.Printf("searchParam: %#v",searchParams)
	// 验证参数
	if err := validate.Struct(&searchParams); err != nil {
		response.Fail(-1, err.Error(),c)
	}
	err, list, total := department.GetList(&searchParams)
	if err != nil {
		response.Fail(-1, err.Error(),c)
	}
	response.OK(resp.PageResult{
		List: list,
		Total: total,
		Page: searchParams.Page,
		PageSize: searchParams.PageSize,
	},c)
}

// @Summary 添加部门
// @Description 添加部门
// @Tags 部门管理
// @Param data body domain.DepartmentAddRequest true "data"
// @Success 0 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /system/department/add [post]
// @Security
func (u *Department) Add(c *gin.Context) {
	var data *domain.DepartmentAddRequest
	// 获取参数
	if err := c.Bind(&data); err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	// 验证参数
	if err := validate.Struct(data); err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	if err := department.Create(data); err != nil {
		response.Fail(-1, "添加失败: " + err.Error(),c)
	} else {
		response.Success("添加成功",c)
	}
}

// @Summary 修改部门
// @Description 修改部门
// @Tags 部门管理
// @Param data body domain.DepartmentEditRequest true "data"
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /system/department/edit [put]
// @Security
func (u *Department) Edit(c *gin.Context) {
	var data *domain.DepartmentEditRequest
	// 获取参数
	if err := c.Bind(&data); err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	// 参数验证
	if err := validate.Struct(data); err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	if err := department.Update(data); err != nil {
		response.Fail(-1, "修改失败: " + err.Error(),c)
	} else {
		response.Success("修改成功",c)
	}
}

// @Summary 删除部门
// @Description 删除部门
// @Tags 部门管理
// @Param ids body request.IdsRequest true "{ids: [1,2']}"
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /system/department/delete [delete]
// @Security
func (u *Department) Delete(c *gin.Context) {
	var data *request.IdsRequest
	// 获取参数
	if err := c.Bind(&data); err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	if len(data.Ids) == 0 {
		response.Fail(-1, "参数不能为空",c)
		return
	}
	err := department.Delete(data.Ids)
	if err != nil {
		response.Fail(-1, "删除失败: " + err.Error(),c)
		return
	}
	response.Success("删除成功",c)
}

// @Summary 获取指定ID的部门详情
// @Description 获取指定ID的部门详情
// @Tags 部门管理
// @Param data query request.GetById true "data"
// @Success 0 {object} response.Response{data=domain.Department} "{"code": 200, "data": [...]}"
// @Router /system/department/detail [get]
// @Security
func (u *Department) Detail(c *gin.Context) {
	var param *request.GetById

	if err := c.Bind(&param); err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	record, err := department.Detail(param.ID)
	if err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	response.OK(record,c)
}
