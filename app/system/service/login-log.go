package service

import (
	"gorm.io/gorm"
	"ruoyi-go/app/system/domain"
	"ruoyi-go/common/lib"
	"strings"
)

type LoginLog struct{}

// GetCount 获取访问日志列表数据数量
//@author: [bluefrog](https://github.com/freewu)
//@function: GetCount
//@description: 获取访问日志列表数据数量
//@param: searchParams *domain.LoginLogSearchRequest
//@return: err error, total int64
func (s LoginLog) GetCount(searchParams *domain.LoginLogSearchRequest) (err error, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.LoginLog{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	return err, total
}

// GetList 获取访问日志列表数据
//@author: [bluefrog](https://github.com/freewu)
//@function: GetList
//@description: 获取访问日志列表数据
//@param: searchParams *domain.LoginLogSearchRequest
//@return: err error, list interface{}, total int64
func (s LoginLog) GetList(searchParams *domain.LoginLogSearchRequest) (err error, list []domain.LoginLog, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.LoginLog{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	// 如果数据 0,也没有必要处理以下动作了
	if total > 0 {
		// 排序
		// orderBy := "id desc"
		db.Order("info_id desc")
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
//@description: 访问日志列表搜索条件过滤
//@param: db *gorm.DB
//@param: searchParams *domain.LoginLogSearchRequest
//@return: *gorm.DB
func (s LoginLog) parseFilter(db *gorm.DB, searchParams *domain.LoginLogSearchRequest) *gorm.DB {
	// 条件过滤
	if searchParams.ID != nil { // 访问日志编号
		db = db.Where("info_id = ?", searchParams.ID)
	}
	if len(searchParams.IDIn) > 0 { // 访问日志编号 数组
		db = db.Where("info_id IN (?)", searchParams.IDIn)
	}
	if len(searchParams.IDNotIn) > 0 { // IDNotIn
		db = db.Where("info_id NOT IN (?)", searchParams.IDNotIn)
	}
	if searchParams.Name != "" { // 登录账号
		db = db.Where("login_name = ?", searchParams.Name)
	}
	if searchParams.IPAddress != "" { // 登录IP地址
		db = db.Where("ipaddr = ?", searchParams.IPAddress)
	}
	if searchParams.Location != "" { // 登录地点
		db = db.Where("login_location = ?", searchParams.Location)
	}
	if searchParams.Browser != "" { // 浏览器类型
		db = db.Where("browser = ?", searchParams.Browser)
	}
	if searchParams.OS != "" { // 操作系统
		db = db.Where("os = ?", searchParams.OS)
	}
	if searchParams.Module != "" { // 登录模块
		db = db.Where("module = ?", searchParams.Module)
	}
	if searchParams.Message != "" { // 提示消息
		db = db.Where("msg LIKE ?", "%"+searchParams.Message+"%")
	}
	if searchParams.Status != nil { // 登录状态（0成功 1失败）
		db = db.Where("status = ?", searchParams.Status)
	}

	if searchParams.BeginDate != "" { // 开始时间
		t, err := lib.StrToTime(searchParams.BeginDate + "00:00:00")
		if err == nil {
			db = db.Where("login_time >= ?", t)
		}
	}
	if searchParams.EndDate != "" { // 结束时间
		t, err := lib.StrToTime(searchParams.EndDate + " 23:59:59")
		if err == nil {
			db = db.Where("login_time <= ?", t)
		}
	}

	if searchParams.Keyword != "" { // 关键词
		k := strings.Trim(searchParams.Keyword, " \t\r\n")
		k1 := "%" + k + "%"
		db = db.Where("(`login_name` LIKE ? OR "+
			"`ipaddr` LIKE ? OR "+
			"`login_location` LIKE ? OR "+
			"`browser` LIKE ? OR "+
			"`os` LIKE ? OR "+
			"`msg` LIKE ? OR "+
			"`module` LIKE ? OR "+
			"`info_id` = ? )", k1, k1, k1, k1, k1, k1, k1, k)
	}
	return db
}

// Create 添加访问日志数据
//@author: [bluefrog](https://github.com/freewu)
//@function: Create
//@description: 添加访问日志数据
//@param: data *domain.LoginLogAddRequest
//@return: error
func (s LoginLog) Create(data *domain.LoginLogCreateRequest) error {
	config := new(domain.LoginLog)
	config.IPAddress = data.IPAddress
	config.Name = data.Name
	config.Location = data.Location
	config.Browser = data.Browser
	config.OS = data.OS
	config.Message = data.Message
	config.Module = data.Module
	config.Time = data.Time

	return domain.DB.Create(&config).Error
}

// Update 修改访问日志
//@author: [bluefrog](https://github.com/freewu)
//@function: Update
//@description: 修改访问日志
//@param: data *model.LoginLogEditRequest
//@return: err error
func (s LoginLog) Update(data *domain.LoginLogUpdateRequest) (err error) {
	record := make(map[string]interface{})
	record["Name"] = data.Name
	record["IPAddress"] = data.IPAddress
	record["Location"] = data.Location
	record["Browser"] = data.Browser
	record["OS"] = data.OS
	record["Message"] = data.Message
	record["Module"] = data.Module
	record["Time"] = data.Time

	return s.UpdateByMap(data.ID, record)
}

// UpdateByMap 更新指定ID的访问日志信息
//@author: [bluefrog](https://github.com/freewu)
//@function: UpdateByMap
//@description: 更新指定ID的访问日志信息
//@param: id uint
//@param: updateMap map[string]interface{}
//@return: err error
func (s LoginLog) UpdateByMap(id uint, updateMap map[string]interface{}) (err error) {
	if err := domain.DB.Model(&domain.LoginLog{}).Where("info_id = ?", id).Updates(updateMap).Error; err != nil {
		return err
	}
	return err
}

// Delete 删除访问日志
//@author: [bluefrog](https://github.com/freewu)
//@function: Delete
//@description: 删除访问日志
//@param: ids []uint
//@return: err error
func (s LoginLog) Delete(ids []uint) (err error) {
	if err = domain.DB.Where("info_id in (?) ", ids).Unscoped().Delete(&domain.LoginLog{}).Error; err != nil {
		return err
	}
	return nil
}

// Detail 获取指定ID访问日志详情
//@author: [bluefrog](https://github.com/freewu)
//@function: Detail
//@description: 获取指定ID访问日志详情
//@param: id uint
//@return: *domain.LoginLog,err error
func (s LoginLog) Detail(id uint) (*domain.LoginLog, error) {
	var detail domain.LoginLog
	db := domain.DB.Where("info_id = ?", id)
	if err := db.First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}
