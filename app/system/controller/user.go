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

type User struct {}

var (
	user = service.User{}
)

// @Summary 获取用户列表
// @Description 用户列表
// @Tags 用户管理
// @Param data query domain.UserSearchRequest true "data"
// @Success 0 {object} response.Response{data=response.PageResult{list=[]domain.User}} "{"code": 0, "data": { "list": [] } }"
// @Router /system/user/list [get]
// @Security
func (u *User) List(c *gin.Context) {
	var searchParams domain.UserSearchRequest
	// 获取查询参数
	if err := c.Bind(&searchParams); err != nil {
		response.Fail(-1, err.Error(),c)
	}
	log.Printf("searchParam: %#v",searchParams)
	// 验证参数
	if err := validate.Struct(&searchParams); err != nil {
		response.Fail(-1, err.Error(),c)
	}
	err, list, total := user.GetList(&searchParams)
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

// @Summary 添加用户
// @Description 添加用户
// @Tags 用户管理
// @Param data body domain.UserAddRequest true "data"
// @Success 0 {object} response.Response "{"code": 0, "data": [...]}"
// @Router /system/user/add [post]
// @Security
func (u *User) Add(c *gin.Context) {
	var data *domain.UserAddRequest
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
	if err := user.Create(data); err != nil {
		response.Fail(-1, "添加失败: " + err.Error(),c)
	} else {
		response.Success("添加成功",c)
	}
}

// @Summary 修改用户
// @Description 修改用户
// @Tags 用户管理
// @Param data body domain.UserEditRequest true "data"
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /system/user/edit [put]
// @Security
func (u *User) Edit(c *gin.Context) {
	var data *domain.UserEditRequest
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
	if err := user.Update(data); err != nil {
		response.Fail(-1, "修改失败: " + err.Error(),c)
	} else {
		response.Success("修改成功",c)
	}
}

// @Summary 删除用户
// @Description 删除用户
// @Tags 用户管理
// @Param ids body request.IdsRequest true "{ids: [1,2']}"
// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /system/user/delete [delete]
// @Security
func (u *User) Delete(c *gin.Context) {
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
	err := user.Delete(data.Ids)
	if err != nil {
		response.Fail(-1, "删除失败: " + err.Error(),c)
		return
	}
	response.Success("删除成功",c)
}

// @Summary 获取指定ID的用户详情
// @Description 获取指定ID的用户详情
// @Tags 用户管理
// @Param data query request.GetById true "data"
// @Success 0 {object} response.Response{data=domain.User} "{"code": 200, "data": [...]}"
// @Router /system/user/detail [get]
// @Security
func (u *User) Detail(c *gin.Context) {
	var param *request.GetById

	if err := c.Bind(&param); err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	user, err := user.Detail(param.ID)
	if err != nil {
		response.Fail(-1, err.Error(),c)
		return
	}
	response.OK(user,c)
}