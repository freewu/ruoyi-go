package service

import (
	"errors"
	"gorm.io/gorm"
	"ruoyi-go/app/system/domain"
	"ruoyi-go/common/model/request"
	"strings"
	"time"
)

type Notice struct{}

// GetCount 获取公告列表数据数量
//@author: [bluefrog](https://github.com/freewu)
//@function: GetCount
//@description: 获取公告列表数据数量
//@param: searchParams *domain.NoticeSearchRequest
//@return: err error, total int64
func (s Notice) GetCount(searchParams *domain.NoticeSearchRequest) (err error, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Notice{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	return err, total
}

// GetList 获取公告列表数据
//@author: [bluefrog](https://github.com/freewu)
//@function: GetList
//@description: 获取公告列表数据
//@param: searchParams *domain.NoticeSearchRequest
//@return: err error, list interface{}, total int64
func (s Notice) GetList(searchParams *domain.NoticeSearchRequest) (err error, list []domain.Notice, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Notice{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	// 如果数据 0,也没有必要处理以下动作了
	if total > 0 {
		// 排序
		// orderBy := "id desc"
		db.Order("notice_id desc")
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
//@description: 公告列表搜索条件过滤
//@param: db *gorm.DB
//@param: searchParams *domain.NoticeSearchRequest
//@return: *gorm.DB
func (s Notice) parseFilter(db *gorm.DB, searchParams *domain.NoticeSearchRequest) *gorm.DB {
	// 条件过滤
	if searchParams.ID != nil { // 公告编号
		db = db.Where("notice_id = ?", searchParams.ID)
	}
	if len(searchParams.IDIn) > 0 { // 公告编号 数组
		db = db.Where("notice_id IN (?)", searchParams.IDIn)
	}
	if len(searchParams.IDNotIn) > 0 { // IDNotIn
		db = db.Where("notice_id NOT IN (?)", searchParams.IDNotIn)
	}
	if searchParams.Title != "" { // 公告标题
		db = db.Where("notice_title like ?", "%"+searchParams.Title+"%")
	}
	if searchParams.Content != "" { // 公告内容
		db = db.Where("notice_content like ?", "%"+searchParams.Content+"%")
	}
	if searchParams.Type != nil { // 公告类型（1通知 2公告）
		db = db.Where("notice_type = ?", searchParams.Type)
	}
	if searchParams.Status != nil { // 公告状态（0正常 1关闭）
		db = db.Where("status = ?", searchParams.Status)
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
		db = db.Where("(`notice_title` LIKE ? OR `notice_content` LIKE ? OR `notice_id` = ? )", k1, k1, k)
		//db = db.Where("(`notice_title` LIKE '%?%' OR `notice_content` LIKE '%?%' OR `notice_id` = ? )", k, k, k)
	}
	return db
}

// Create 添加公告数据
//@author: [bluefrog](https://github.com/freewu)
//@function: Create
//@description: 添加公告数据
//@param: data *domain.NoticeAddRequest
//@return: error
func (s Notice) Create(data *domain.NoticeCreateRequest) error {
	// 判断公告title是否唯一
	//if s.TileIsExist(data.Title, nil) {
	//	return errors.New("已存在相同的公告标题")
	//}
	if data.Type < 0 {
		return errors.New("请选择公告类型")
	}
	notice := new(domain.Notice)
	notice.Title = data.Title
	notice.Content = data.Content
	notice.Type = data.Type
	notice.Status = data.Status

	notice.CreateBy = "" // todo 通过登录服务获取
	notice.CreateTime = time.Now()
	notice.UpdateBy = "" // todo 通过登录服务获取
	notice.UpdateTime = time.Now()

	return domain.DB.Create(&notice).Error
}

// Update 修改公告
//@author: [bluefrog](https://github.com/freewu)
//@function: Update
//@description: 修改公告
//@param: data *model.NoticeEditRequest
//@return: err error
func (s Notice) Update(data *domain.NoticeUpdateRequest) (err error) {
	// 判断公告标题是否唯一
	//if s.TitleIsExist(data.Title, &[]uint{data.ID}) {
	//	return errors.New("已存在相同的公告标题")
	//}
	if data.Type < 0 {
		return errors.New("请选择公告类型")
	}
	record := make(map[string]interface{})
	record["Title"] = data.Title
	record["Content"] = data.Content
	record["Type"] = data.Type
	record["Status"] = data.Status
	record["UpdateBy"] = "" // todo 通过登录服务获取
	record["UpdateTime"] = time.Now()

	return s.UpdateByMap(data.ID, record)
}

// UpdateByMap 更新指定ID的公告信息
//@author: [bluefrog](https://github.com/freewu)
//@function: UpdateByMap
//@description: 更新指定ID的公告信息
//@param: id uint
//@param: updateMap map[string]interface{}
//@return: err error
func (s Notice) UpdateByMap(id uint, updateMap map[string]interface{}) (err error) {
	if err := domain.DB.Model(&domain.Notice{}).Where("notice_id = ?", id).Updates(updateMap).Error; err != nil {
		return err
	}
	return err
}

// Delete 删除公告
//@author: [bluefrog](https://github.com/freewu)
//@function: Delete
//@description: 删除公告
//@param: ids []uint
//@return: err error
func (s Notice) Delete(ids []uint) (err error) {
	if err = domain.DB.Where("notice_id in (?) ", ids).Unscoped().Delete(&domain.Notice{}).Error; err != nil {
		return err
	}
	return nil
}

// Detail 获取指定ID公告详情
//@author: [bluefrog](https://github.com/freewu)
//@function: Detail
//@description: 获取指定ID公告详情
//@param: id uint
//@return: *domain.Notice,err error
func (s Notice) Detail(id uint) (*domain.Notice, error) {
	var detail domain.Notice
	db := domain.DB.Where("notice_id = ?", id)
	if err := db.First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

// TitleIsExist 判断公告标题是否存在
//@author: [bluefrog](https://github.com/freewu)
//@function: TitleIsExist
//@description: 判断公告标题是否存在
//@param: title string 公告标题
//@param: excludeIds *[]uint 不包含的公告ID
//@return: err error
func (s Notice) TitleIsExist(title string, excludeIds *[]uint) bool {
	// 判断名字是否唯一
	var filter = &domain.NoticeSearchRequest{
		Title:    title,
		PageInfo: request.PageInfo{PageSize: 1},
	}
	if excludeIds != nil {
		filter.IDNotIn = *excludeIds
	}
	_, total := s.GetCount(filter)
	return total > 0
}
