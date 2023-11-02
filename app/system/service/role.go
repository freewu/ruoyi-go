package service

import (
	"errors"
	"gorm.io/gorm"
	"ruoyi-go/app/system/domain"
	"strings"
	"time"
)

type Role struct{}

// GetCount 获取角色列表数据数量
//@author: [bluefrog](https://github.com/freewu)
//@function: GetCount
//@description: 获取角色列表数据数量
//@param: searchParams *domain.RoleSearchRequest
//@return: err error, total int64
func (s Role) GetCount(searchParams *domain.RoleSearchRequest) (err error, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Role{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	return err, total
}

// GetList 获取角色列表数据
//@author: [bluefrog](https://github.com/freewu)
//@function: GetList
//@description: 获取角色列表数据
//@param: searchParams *domain.RoleSearchRequest
//@return: err error, list interface{}, total int64
func (s Role) GetList(searchParams *domain.RoleSearchRequest) (err error, list []domain.Role, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Role{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	// 如果数据 0,也没有必要处理以下动作了
	if total > 0 {
		// 排序
		// orderBy := "id desc"
		db.Order("role_id desc")
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
//@description: 角色列表搜索条件过滤
//@param: db *gorm.DB
//@param: searchParams *domain.RoleSearchRequest
//@return: *gorm.DB
func (s Role) parseFilter(db *gorm.DB, searchParams *domain.RoleSearchRequest) *gorm.DB {
	// 条件过滤
	if searchParams.ID != nil { // 角色编号
		db = db.Where("role_id = ?", searchParams.ID)
	}
	if len(searchParams.IDIn) > 0 { // 角色编号 数组
		db = db.Where("role_id IN (?)", searchParams.IDIn)
	}
	if len(searchParams.IDNotIn) > 0 { // IDNotIn
		db = db.Where("role_id NOT IN (?)", searchParams.IDNotIn)
	}
	if searchParams.Name != "" { // 角色名称
		db = db.Where("role_name = ?", searchParams.Name)
	}
	if searchParams.Key != "" { // 角色权限字符串
		db = db.Where("role_key = ?", searchParams.Key)
	}
	if searchParams.Status != nil { // 角色状态（0正常 1停用）
		db = db.Where("status = ?", searchParams.Status)
	}
	if searchParams.MenuCheckStrictly != nil { // 菜单树选择项是否关联显示（0否 1是）
		db = db.Where("menu_check_strictly = ?", searchParams.MenuCheckStrictly)
	}
	if searchParams.DepartmentCheckStrictly != nil { // 部门树选择项是否关联显示（0否 1是）
		db = db.Where("dept_check_strictly = ?", searchParams.Status)
	}
	if searchParams.DataScope != nil { // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		db = db.Where("data_scope = ?", searchParams.Status)
	}
	if searchParams.DeleteFlag != nil { // 删除标志（0代表存在 2代表删除）
		db = db.Where("del_flag = ?", searchParams.Status)
	}
	if searchParams.CreateBy != "" { // 创建者
		db = db.Where("create_by LIKE ?", "%"+searchParams.CreateBy+"%")
	}
	if searchParams.UpdateBy != "" { // 更新者
		db = db.Where("update_by LIKE ?", "%"+searchParams.UpdateBy+"%")
	}
	if searchParams.Keyword != "" { // 关键词
		k := strings.Trim(searchParams.Keyword, " \t\r\n")
		k1 := "%" + k + "%"
		db = db.Where("(`role_name` LIKE ? OR `role_key` LIKE ? OR `role_id` = ? )", k1, k1, k)
	}
	return db
}

// Create 添加角色数据
//@author: [bluefrog](https://github.com/freewu)
//@function: Create
//@description: 添加角色数据
//@param: data *domain.RoleAddRequest
//@return: error
func (s Role) Create(data *domain.RoleCreateRequest) error {
	// 判断角色名称是否唯一
	if s.NameIsExist(data.Name, nil) {
		return errors.New("已存在相同的角色名称")
	}
	// 判断角色权限字符串是否唯一
	if s.KeyIsExist(data.Key, nil) {
		return errors.New("已存在相同的角色权限字符串")
	}
	role := new(domain.Role)
	role.Key = data.Key
	role.Name = data.Name
	role.Status = data.Status
	role.Sort = data.Sort
	role.MenuCheckStrictly = data.MenuCheckStrictly
	role.DepartmentCheckStrictly = data.DepartmentCheckStrictly
	role.DataScope = data.DataScope

	role.CreateBy = "" // todo 通过登录服务获取
	role.CreateTime = time.Now()
	role.UpdateBy = "" // todo 通过登录服务获取
	role.UpdateTime = time.Now()

	return domain.DB.Create(&role).Error
}

// Update 修改角色
//@author: [bluefrog](https://github.com/freewu)
//@function: Update
//@description: 修改角色
//@param: data *model.RoleEditRequest
//@return: err error
func (s Role) Update(data *domain.RoleUpdateRequest) (err error) {
	// 判断名字是否唯一
	if s.NameIsExist(data.Name, &[]uint{data.ID}) {
		return errors.New("已存在相同的角色名称")
	}
	// 判断角色权限字符串是否唯一
	if s.KeyIsExist(data.Key, &[]uint{data.ID}) {
		return errors.New("已存在相同的角色权限字符串")
	}
	record := make(map[string]interface{})
	record["Name"] = data.Name
	record["Key"] = data.Key
	record["MenuCheckStrictly"] = data.MenuCheckStrictly
	record["DepartmentCheckStrictly"] = data.DepartmentCheckStrictly
	record["DataScope"] = data.DataScope
	record["DeleteFlag"] = data.DeleteFlag
	record["Status"] = data.Status
	record["Sort"] = data.Sort
	record["UpdateBy"] = "" // todo 通过登录服务获取
	record["UpdateTime"] = time.Now()

	return s.UpdateByMap(data.ID, record)
}

// UpdateByMap 更新指定ID的角色信息
//@author: [bluefrog](https://github.com/freewu)
//@function: UpdateByMap
//@description: 更新指定ID的角色信息
//@param: id uint
//@param: updateMap map[string]interface{}
//@return: err error
func (s Role) UpdateByMap(id uint, updateMap map[string]interface{}) (err error) {
	if err := domain.DB.Model(&domain.Role{}).Where("role_id = ?", id).Updates(updateMap).Error; err != nil {
		return err
	}
	return err
}

// Delete 删除角色
//@author: [bluefrog](https://github.com/freewu)
//@function: Delete
//@description: 删除角色
//@param: ids []uint
//@return: err error
func (s Role) Delete(ids []uint) (err error) {
	if err = domain.DB.Where("role_id in (?) ", ids).Unscoped().Delete(&domain.Role{}).Error; err != nil {
		return err
	}
	return nil
}

// Detail 获取指定ID角色详情
//@author: [bluefrog](https://github.com/freewu)
//@function: Detail
//@description: 获取指定ID角色详情
//@param: id uint
//@return: *domain.Role,err error
func (s Role) Detail(id uint) (*domain.Role, error) {
	var detail domain.Role
	db := domain.DB.Where("role_id = ?", id)
	if err := db.First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

// NameIsExist 判断角色名称是否存在
//@author: [bluefrog](https://github.com/freewu)
//@function: NameIsExist
//@description: 判断角色名称是否存在
//@param: name string 角色名称
//@param: excludeIds *[]uint 不包含的角色ID
//@return: err error
func (s Role) NameIsExist(name string, excludeIds *[]uint) bool {
	// 判断名字是否唯一
	var filter = &domain.RoleSearchRequest{
		Name: name,
	}
	if excludeIds != nil {
		filter.IDNotIn = *excludeIds
	}
	_, total := s.GetCount(filter)
	return total > 0
}

// KeyIsExist 判断角色编号是否存在
//@author: [bluefrog](https://github.com/freewu)
//@function: KeyIsExist
//@description: 判断角色编号是否存在
//@param: code string 角色编号
//@param: excludeIds *[]uint 不包含的区域类型ID
//@return: err error
func (s Role) KeyIsExist(key string, excludeIds *[]uint) bool {
	// 判断名字是否唯一
	var filter = &domain.RoleSearchRequest{
		Key: key,
	}
	if excludeIds != nil {
		filter.IDNotIn = *excludeIds
	}
	_, total := s.GetCount(filter)
	return total > 0
}
