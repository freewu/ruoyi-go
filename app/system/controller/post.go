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

type Post struct{}

var (
	post = service.Post{}
)

// @Summary 获取职位列表
// @Description 职位列表
// @Tags 职位管理
// @Param data query domain.PostSearchRequest true "data"
// @Success 0 {object} response.Response{data=response.PageResult{list=[]domain.Post}} "{"code": 0, "msg": "success","data": { "list": [],"total": 0 } }"
// @Router /system/post/list [get]
// @Security
func (u *Post) List(c *gin.Context) {
	var searchParams domain.PostSearchRequest
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
	err, list, total := Post.GetList(&searchParams)
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

// @Summary 添加职位
// @Description 添加职位
// @Tags 职位管理
// @Param data body domain.PostCreateRequest true "data"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/post/create [post]
// @Security
func (u *Post) Create(c *gin.Context) {
	var data *domain.PostCreateRequest
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
	if err := Post.Create(data); err != nil {
		response.Fail(-1, "添加失败: "+err.Error(), c)
	} else {
		response.Success("添加成功", c)
	}
}

// @Summary 修改职位
// @Description 修改职位
// @Tags 职位管理
// @Param data body domain.PostUpdateRequest true "data"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/post/update [put]
// @Security
func (u *Post) Update(c *gin.Context) {
	var data *domain.PostUpdateRequest
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
	if err := Post.Update(data); err != nil {
		response.Fail(-1, "修改失败: "+err.Error(), c)
	} else {
		response.Success("修改成功", c)
	}
}

// @Summary 删除职位
// @Description 删除职位
// @Tags 职位管理
// @Param ids body request.IdsRequest true "{ids: [1,2']}"
// @Success 0 {object} response.Response "{"code": 0, "msg":"success","data": null}"
// @Router /system/post/delete [delete]
// @Security
func (u *Post) Delete(c *gin.Context) {
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
	// 如果有子职位存在，禁止删除,需要先删除完子职位
	_, _, total := Post.GetList(&domain.PostSearchRequest{
		ParentIDIn: data.Ids,
		PageInfo:   request.PageInfo{PageSize: 1},
	})
	if total > 0 {
		response.Fail(-1, "存在子职位，禁止删除", c)
		return
	}
	err := Post.Delete(data.Ids)
	if err != nil {
		response.Fail(-1, "删除失败: "+err.Error(), c)
		return
	}
	response.Success("删除成功", c)
}

// @Summary 获取指定ID的职位详情
// @Description 获取指定ID的职位详情
// @Tags 职位管理
// @Param data query request.GetById true "data"
// @Success 0 {object} response.Response{data=domain.Post} "{"code": 0, "msg":"success","data": {}}"
// @Router /system/post/detail [get]
// @Security
func (u *Post) Detail(c *gin.Context) {
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
	record, err := Post.Detail(param.ID)
	if err != nil {
		response.Fail(-1, err.Error(), c)
		return
	}
	response.OK(record, c)
}
