package service

import (
	"errors"
	"gorm.io/gorm"
	"ruoyi-go/app/system/domain"
	"ruoyi-go/common/model/request"
	"strings"
)

type Department struct{}

// GetList
//@author: [bluefrog](https://github.com/freewu)
//@function: GetList
//@description: 获取部门列表数据
//@param: searchParams *domain.DepartmentSearchRequest
//@return: err error, list interface{}, total int64
func (s Department) GetList(searchParams *domain.DepartmentSearchRequest) (err error, list []domain.Department, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Department{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	// 如果数据 0,也没有必要处理以下动作了
	if total > 0 {
		// 排序
		// orderBy := "id desc"
		db.Order("id desc")
		if searchParams.OrderBy != "" {
			db.Order(searchParams.OrderBy)
		}
		// 分页
		if searchParams.PageSize > 0 {
			limit := searchParams.PageSize
			offset := searchParams.PageSize * (searchParams.Page - 1)
			db.Limit(limit).Offset(offset)
		}
		err = db.Find(&list).Error
	}
	return err, list, total
}

//@author: [bluefrog](https://github.com/freewu)
//@function: parseFilter
//@description: 部门列表搜索条件过滤
//@param: db *gorm.DB
//@param: searchParams *domain.DepartmentSearchRequest
//@return: *gorm.DB
func (s Department) parseFilter(db *gorm.DB, searchParams *domain.DepartmentSearchRequest) *gorm.DB {
	// 条件过滤
	if searchParams.ParentID != nil { // 父部门ID
		db = db.Where("parent_id = ?", searchParams.ParentID)
	}
	if len(searchParams.ParentIDIn) > 0 { // 父部门ID 数组
		db = db.Where("parent_id IN (?)", searchParams.ParentIDIn)
	}
	if len(searchParams.IDNotIn) > 0 { // IDNotIn
		db = db.Where("id NOT IN (?)", searchParams.IDNotIn)
	}
	if len(searchParams.IDIn) > 0 { // IDIn
		db = db.Where("id IN (?)", searchParams.IDIn)
	}
	if searchParams.Name != "" { // 部门名称
		db = db.Where("name like ?", "%"+searchParams.Name+"%")
	}
	if searchParams.Ancestors != "" { // 祖级列表
		db = db.Where("ancestors LIKE ?", "%"+searchParams.Ancestors+"%")
	}
	if searchParams.Status != nil { // 部门状态（0正常 1停用）
		db = db.Where("status = ?", searchParams.Status)
	}
	if searchParams.CreateBy != "" { // 创建者
		db = db.Where("create_by LIKE ?", "%"+searchParams.CreateBy+"%")
	}
	if searchParams.UpdateBy != "" { // 更新者
		db = db.Where("update_by LIKE ?", "%"+searchParams.UpdateBy+"%")
	}
	if searchParams.Keyword != "" { // 关键词
		k1 := strings.Trim(searchParams.Keyword, " \t\r\n")
		k := "%" + k1 + "%"
		db = db.Where("(name LIKE ?  )", k)
	}
	return db
}

//@author: [bluefrog](https://github.com/freewu)
//@function: Create
//@description: 添加部门数据
//@param: data *domain.DepartmentAddRequest
//@return: error
func (s Department) Create(data *domain.DepartmentCreateRequest) error {
	// 判断部门名称是否唯一
	if s.IsExist(data.Name, nil) {
		return errors.New("已存在相同的部门名称")
	}
	department := new(domain.Department)
	department.ParentID = data.ParentID
	department.Name = data.Name
	department.Ancestors = data.Ancestors
	department.Order = data.Order
	department.Status = data.Status

	department.CreateBy = "" // todo 通过登录服务获取
	department.UpdateBy = "" // todo 通过登录服务获取

	return domain.DB.Create(&department).Error
}

//@author: [bluefrog](https://github.com/freewu)
//@function: Update
//@description: 修改区域
//@param: data *model.DepartmentEditRequest
//@return: err error
func (s Department) Update(data *domain.DepartmentUpdateRequest) (err error) {
	// 判断名字是否唯一
	if s.IsExist(data.Name, &[]uint{data.ID}) {
		return errors.New("已存在相同的部门名称")
	}
	record := make(map[string]interface{})
	record["ParentID"] = data.ParentID
	record["Name"] = data.Name
	record["Ancestors"] = data.Ancestors
	record["Order"] = data.Order
	record["Status"] = data.Status

	record["UpdateBy"] = "" // todo 通过登录服务获取

	return s.UpdateByMap(data.ID, record)
}

//@author: [bluefrog](https://github.com/freewu)
//@function: UpdateByMap
//@description: 更新指定ID的部门信息
//@param: id uint
//@param: updateMap map[string]interface{}
//@return: err error
func (s Department) UpdateByMap(id uint, updateMap map[string]interface{}) (err error) {
	if err := domain.DB.Model(&domain.Department{}).Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return err
	}
	return err
}

//@author: [bluefrog](https://github.com/freewu)
//@function: Delete
//@description: 删除部门
//@param: ids []uint
//@return: err error
func (s Department) Delete(ids []uint) (err error) {
	// 删除区域
	if err = domain.DB.Where("id in (?) ", ids).Unscoped().Delete(&domain.Department{}).Error; err != nil {
		return err
	}
	return nil
}

//@author: [bluefrog](https://github.com/freewu)
//@function: Detail
//@description: 获取指定ID部门详情
//@param: id uint
//@return: *domain.Department,err error
func (s Department) Detail(id uint) (*domain.Department, error) {
	var detail domain.Department
	db := domain.DB.Where("id = ?", id)
	if err := db.First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

//@author: [bluefrog](https://github.com/freewu)
//@function: IsExistAreaType
//@description: 判断区域类型是否存在
//@param: name string 区域类型名称
//@param: excludeIds *[]uint 不包含的区域类型ID
//@return: err error
func (s Department) IsExist(name string, excludeIds *[]uint) bool {
	// 判断名字是否唯一
	var filter = &domain.DepartmentSearchRequest{
		Name:     name,
		PageInfo: request.PageInfo{PageSize: 1},
	}
	if excludeIds != nil {
		filter.IDNotIn = *excludeIds
	}
	_, _, total := s.GetList(filter)
	return total > 0
}
