package service

import (
	"errors"
	"gorm.io/gorm"
	"ruoyi-go/app/system/domain"
	"strings"
	"time"
)

type Config struct{}

// GetCount 获取参数列表数据数量
//@author: [bluefrog](https://github.com/freewu)
//@function: GetCount
//@description: 获取参数列表数据数量
//@param: searchParams *domain.ConfigSearchRequest
//@return: err error, total int64
func (s Config) GetCount(searchParams *domain.ConfigSearchRequest) (err error, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Config{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	return err, total
}

// GetList 获取参数列表数据
//@author: [bluefrog](https://github.com/freewu)
//@function: GetList
//@description: 获取参数列表数据
//@param: searchParams *domain.ConfigSearchRequest
//@return: err error, list interface{}, total int64
func (s Config) GetList(searchParams *domain.ConfigSearchRequest) (err error, list []domain.Config, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Config{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	// 如果数据 0,也没有必要处理以下动作了
	if total > 0 {
		// 排序
		// orderBy := "id desc"
		db.Order("config_id desc")
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
//@description: 参数列表搜索条件过滤
//@param: db *gorm.DB
//@param: searchParams *domain.ConfigSearchRequest
//@return: *gorm.DB
func (s Config) parseFilter(db *gorm.DB, searchParams *domain.ConfigSearchRequest) *gorm.DB {
	// 条件过滤
	if searchParams.ID != nil { // 参数编号
		db = db.Where("config_id = ?", searchParams.ID)
	}
	if len(searchParams.IDIn) > 0 { // 参数编号 数组
		db = db.Where("config_id IN (?)", searchParams.IDIn)
	}
	if len(searchParams.IDNotIn) > 0 { // IDNotIn
		db = db.Where("config_id NOT IN (?)", searchParams.IDNotIn)
	}
	if searchParams.Name != "" { // 参数名称
		db = db.Where("config_name = ?", searchParams.Name)
	}
	if searchParams.Key != "" { // 参数键名
		db = db.Where("config_key = ?", searchParams.Key)
	}
	if searchParams.Value != "" { // 参数键值
		db = db.Where("config_value LIKE ?", "%"+searchParams.Value+"%")
	}
	if searchParams.Type != "" { // 是否系统内置（Y是 | N否）
		db = db.Where("config_type = ?", searchParams.Type)
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
		db = db.Where("(`config_name` LIKE ? OR "+
			"`config_key` LIKE ? OR "+
			"`config_value` LIKE ? OR "+
			"`config_id` = ? )", k1, k1, k1, k)
	}
	return db
}

// Create 添加参数数据
//@author: [bluefrog](https://github.com/freewu)
//@function: Create
//@description: 添加参数数据
//@param: data *domain.ConfigAddRequest
//@return: error
func (s Config) Create(data *domain.ConfigCreateRequest) error {
	// 判断参数名称是否唯一
	if s.NameIsExist(data.Name, nil) {
		return errors.New("已存在相同的参数名称")
	}
	// 判断参数键名是否唯一
	if s.KeyIsExist(data.Key, nil) {
		return errors.New("已存在相同的参数键名")
	}
	config := new(domain.Config)
	config.Key = data.Key
	config.Name = data.Name
	config.Type = data.Type
	config.Value = data.Value

	config.CreateBy = "" // todo 通过登录服务获取
	config.CreateTime = time.Now()
	config.UpdateBy = "" // todo 通过登录服务获取
	config.UpdateTime = time.Now()
	config.Remark = data.Remark

	return domain.DB.Create(&config).Error
}

// Update 修改参数
//@author: [bluefrog](https://github.com/freewu)
//@function: Update
//@description: 修改参数
//@param: data *model.ConfigEditRequest
//@return: err error
func (s Config) Update(data *domain.ConfigUpdateRequest) (err error) {
	// 判断名字是否唯一
	if s.NameIsExist(data.Name, &[]uint{data.ID}) {
		return errors.New("已存在相同的参数名称")
	}
	// 判断参数键名是否唯一
	if s.KeyIsExist(data.Key, &[]uint{data.ID}) {
		return errors.New("已存在相同的参数键名")
	}
	record := make(map[string]interface{})
	record["Name"] = data.Name
	record["Key"] = data.Key
	record["Type"] = data.Type
	record["Value"] = data.Value
	record["Remark"] = data.Remark
	record["UpdateBy"] = "" // todo 通过登录服务获取
	record["UpdateTime"] = time.Now()

	return s.UpdateByMap(data.ID, record)
}

// UpdateByMap 更新指定ID的参数信息
//@author: [bluefrog](https://github.com/freewu)
//@function: UpdateByMap
//@description: 更新指定ID的参数信息
//@param: id uint
//@param: updateMap map[string]interface{}
//@return: err error
func (s Config) UpdateByMap(id uint, updateMap map[string]interface{}) (err error) {
	if err := domain.DB.Model(&domain.Config{}).Where("config_id = ?", id).Updates(updateMap).Error; err != nil {
		return err
	}
	return err
}

// Delete 删除参数
//@author: [bluefrog](https://github.com/freewu)
//@function: Delete
//@description: 删除参数
//@param: ids []uint
//@return: err error
func (s Config) Delete(ids []uint) (err error) {
	if err = domain.DB.Where("config_id in (?) ", ids).Unscoped().Delete(&domain.Config{}).Error; err != nil {
		return err
	}
	return nil
}

// Detail 获取指定ID参数详情
//@author: [bluefrog](https://github.com/freewu)
//@function: Detail
//@description: 获取指定ID参数详情
//@param: id uint
//@return: *domain.Config,err error
func (s Config) Detail(id uint) (*domain.Config, error) {
	var detail domain.Config
	db := domain.DB.Where("config_id = ?", id)
	if err := db.First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

// NameIsExist 判断参数名称是否存在
//@author: [bluefrog](https://github.com/freewu)
//@function: NameIsExist
//@description: 判断参数名称是否存在
//@param: name string 参数名称
//@param: excludeIds *[]uint 不包含的参数ID
//@return: err error
func (s Config) NameIsExist(name string, excludeIds *[]uint) bool {
	// 判断名字是否唯一
	var filter = &domain.ConfigSearchRequest{
		Name: name,
	}
	if excludeIds != nil {
		filter.IDNotIn = *excludeIds
	}
	_, total := s.GetCount(filter)
	return total > 0
}

// KeyIsExist 判断参数键名是否存在
//@author: [bluefrog](https://github.com/freewu)
//@function: KeyIsExist
//@description: 判断参数键名是否存在
//@param: key string 参数键名
//@param: excludeIds *[]uint 不包含的区域类型ID
//@return: err error
func (s Config) KeyIsExist(key string, excludeIds *[]uint) bool {
	// 判断名字是否唯一
	var filter = &domain.ConfigSearchRequest{
		Key: key,
	}
	if excludeIds != nil {
		filter.IDNotIn = *excludeIds
	}
	_, total := s.GetCount(filter)
	return total > 0
}
