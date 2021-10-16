package controller

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/system/service"
)

type User struct {}

// @Summary 获取用户列表
// @Description 用户列表
// @Tags 用户管理
// @Param data query model.UserSearchParam true "data"
// @Success 200 {object} response.Response{data=response.PageResult{list=[]model.User}} "{"code": 0, "data": { "list": [] } }"
// @Router /system/user/list [get]
// @Security
func (c *User) List(r *gin.Context) {
	var searchParams *model.AreaSearchParam

	if err := r.Parse(&searchParams); err != nil {
		response.FailJson(true, r, err.Error())
	}
	err, list, total := service.User.GetList(searchParams)
	if err != nil {
		response.FailJson(true, r, err.Error())
	}
	response.SusJson(true, r, "成功", model.PageResult{
		List: list,
		Total: total,
		Page: searchParams.Page,
		PageSize: searchParams.PageSize,
	})
}

//// @Summary 获取指定ID的用户详情
//// @Description 获取指定ID的用户详情
//// @Tags 用户管理
//// @Param data query model.GetById true "data"
//// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
//// @Router /system/user/detail [get]
//// @Security
//func (c *User) Detail(r *gin.Context) {
//	var param *model.GetById
//
//	if err := r.Parse(&param); err != nil {
//		response.FailJson(true, r, err.Error())
//	}
//	area, err := service.GetArea(param.ID)
//	if err != nil {
//		response.FailJson(true, r, err.Error())
//	}
//	response.SusJson(true, r, "成功", area)
//}
//
//// @Summary 添加用户
//// @Description 添加用户
//// @Tags 用户管理
//// @Param data body model.AreaAddRequest true "data"
//// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
//// @Router /system/user/add [put]
//// @Security
//func (c *User) Add(r *gin.Context) {
//	var data *model.AreaAddRequest
//	// 获取参数
//	if err := r.Parse(&data); err != nil {
//		response.FailJson(true, r, err.Error())
//	}
//	if err := service.CreateArea(data); err != nil {
//		response.FailJson(true, r, "添加失败: " + err.Error())
//	} else {
//		response.SusJson(true, r, "添加成功")
//	}
//}
//
//// @Summary 修改用户
//// @Description 修改用户
//// @Tags 用户管理
//// @Param data body model.AreaEditRequest true "data"
//// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
//// @Router /system/user/edit [post]
//// @Security
//func (c *User) Edit(r *gin.Context) {
//	var data *model.AreaEditRequest
//	// 获取参数
//	if err := r.Parse(&data); err != nil {
//		response.FailJson(true, r, err.Error())
//	}
//	if err := service.UpdateArea(data); err != nil {
//		response.FailJson(true, r, "修改失败: " + err.Error())
//	} else {
//		response.SusJson(true, r, "修改成功")
//	}
//}
//
//// @Summary 删除用户
//// @Description 删除用户
//// @Tags 用户管理
//// @Param ids body model.IdsReq true "{ids: [1,2']}"
//// @Success 0 {object} response.Response "{"code": 200, "data": [...]}"
//// @Router /system/user/delete [delete]
//// @Security
//func (c *User) Delete(r *gin.Context) {
//	ids := r.GetInts("ids")
//	if len(ids) == 0 {
//		response.FailJson(true, r, "参数不能为空")
//	}
//	err := service.DeleteArea(ids)
//	if err != nil {
//		response.FailJson(true, r, "删除失败: " + err.Error())
//	}
//	response.SusJson(true, r, "删除成功")
//}