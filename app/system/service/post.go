package service

import (
	"errors"
	"gorm.io/gorm"
	"ruoyi-go/app/system/domain"
	"ruoyi-go/common/model/request"
	"strings"
	"time"
)

type Post struct{}

// GetCount 获取职位列表数据数量
//@author: [bluefrog](https://github.com/freewu)
//@function: GetCount
//@description: 获取职位列表数据数量
//@param: searchParams *domain.PostSearchRequest
//@return: err error, total int64
func (s Post) GetCount(searchParams *domain.PostSearchRequest) (err error, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Post{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	return err, total
}

// GetList 获取职位列表数据
//@author: [bluefrog](https://github.com/freewu)
//@function: GetList
//@description: 获取职位列表数据
//@param: searchParams *domain.PostSearchRequest
//@return: err error, list interface{}, total int64
func (s Post) GetList(searchParams *domain.PostSearchRequest) (err error, list []domain.Post, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.Post{})
	// 条件过滤
	db = s.parseFilter(db, searchParams)
	// 统计数据
	err = db.Count(&total).Error
	// 如果数据 0,也没有必要处理以下动作了
	if total > 0 {
		// 排序
		// orderBy := "id desc"
		db.Order("post_id desc")
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
//@description: 职位列表搜索条件过滤
//@param: db *gorm.DB
//@param: searchParams *domain.PostSearchRequest
//@return: *gorm.DB
func (s Post) parseFilter(db *gorm.DB, searchParams *domain.PostSearchRequest) *gorm.DB {
	// 条件过滤
	if searchParams.ID != nil { // 职位编号
		db = db.Where("post_id = ?", searchParams.ID)
	}
	if len(searchParams.IDIn) > 0 { // 职位编号 数组
		db = db.Where("post_id IN (?)", searchParams.IDIn)
	}
	if len(searchParams.IDNotIn) > 0 { // IDNotIn
		db = db.Where("post_id NOT IN (?)", searchParams.IDNotIn)
	}
	if searchParams.Name != "" { // 职位名称
		db = db.Where("post_name = ?", searchParams.Name)
	}
	if searchParams.Code != "" { // 职位编码
		db = db.Where("post_code = ?", searchParams.Code)
	}
	if searchParams.Status != nil { // 职位状态（0正常 1停用）
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
		//db = db.Where("(`post_name` LIKE '%?%` OR `post_code` LIKE '%?%' OR `post_id` = ? )", k, k, k)
		db = db.Where("(`post_name` LIKE ? OR `post_code` LIKE ? OR `post_id` = ? )", k1, k1, k)
	}
	return db
}

// Create 添加职位数据
//@author: [bluefrog](https://github.com/freewu)
//@function: Create
//@description: 添加职位数据
//@param: data *domain.PostAddRequest
//@return: error
func (s Post) Create(data *domain.PostCreateRequest) error {
	// 判断职位名称是否唯一
	if s.NameIsExist(data.Name, nil) {
		return errors.New("已存在相同的职位名称")
	}
	// 判断职位编号是否唯一
	if s.CodeIsExist(data.Code, nil) {
		return errors.New("已存在相同的职位编号")
	}
	post := new(domain.Post)
	post.Code = data.Code
	post.Name = data.Name
	post.Status = data.Status
	post.Sort = data.Sort

	post.CreateBy = "" // todo 通过登录服务获取
	post.CreateTime = time.Now()
	post.UpdateBy = "" // todo 通过登录服务获取
	post.UpdateTime = time.Now()

	return domain.DB.Create(&post).Error
}

// Update 修改职位
//@author: [bluefrog](https://github.com/freewu)
//@function: Update
//@description: 修改职位
//@param: data *model.PostEditRequest
//@return: err error
func (s Post) Update(data *domain.PostUpdateRequest) (err error) {
	// 判断名字是否唯一
	if s.NameIsExist(data.Name, &[]uint{data.ID}) {
		return errors.New("已存在相同的职位名称")
	}
	// 判断职位编号是否唯一
	if s.CodeIsExist(data.Code, &[]uint{data.ID}) {
		return errors.New("已存在相同的职位编号")
	}
	record := make(map[string]interface{})
	record["Name"] = data.Name
	record["Code"] = data.Code
	record["Status"] = data.Status
	record["Sort"] = data.Sort
	record["UpdateBy"] = "" // todo 通过登录服务获取
	record["UpdateTime"] = time.Now()

	return s.UpdateByMap(data.ID, record)
}

// UpdateByMap 更新指定ID的职位信息
//@author: [bluefrog](https://github.com/freewu)
//@function: UpdateByMap
//@description: 更新指定ID的职位信息
//@param: id uint
//@param: updateMap map[string]interface{}
//@return: err error
func (s Post) UpdateByMap(id uint, updateMap map[string]interface{}) (err error) {
	if err := domain.DB.Model(&domain.Post{}).Where("post_id = ?", id).Updates(updateMap).Error; err != nil {
		return err
	}
	return err
}

// Delete 删除职位
//@author: [bluefrog](https://github.com/freewu)
//@function: Delete
//@description: 删除职位
//@param: ids []uint
//@return: err error
func (s Post) Delete(ids []uint) (err error) {
	if err = domain.DB.Where("post_id in (?) ", ids).Unscoped().Delete(&domain.Post{}).Error; err != nil {
		return err
	}
	return nil
}

// Detail 获取指定ID职位详情
//@author: [bluefrog](https://github.com/freewu)
//@function: Detail
//@description: 获取指定ID职位详情
//@param: id uint
//@return: *domain.Post,err error
func (s Post) Detail(id uint) (*domain.Post, error) {
	var detail domain.Post
	db := domain.DB.Where("post_id = ?", id)
	if err := db.First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

// NameIsExist 判断职位名称是否存在
//@author: [bluefrog](https://github.com/freewu)
//@function: NameIsExist
//@description: 判断职位名称是否存在
//@param: name string 职位名称
//@param: excludeIds *[]uint 不包含的职位ID
//@return: err error
func (s Post) NameIsExist(name string, excludeIds *[]uint) bool {
	// 判断名字是否唯一
	var filter = &domain.PostSearchRequest{
		Name:     name,
		PageInfo: request.PageInfo{PageSize: 1},
	}
	if excludeIds != nil {
		filter.IDNotIn = *excludeIds
	}
	_, total := s.GetCount(filter)
	return total > 0
}

// CodeIsExist 判断职位编号是否存在
//@author: [bluefrog](https://github.com/freewu)
//@function: CodeIsExist
//@description: 判断职位编号是否存在
//@param: code string 职位编号
//@param: excludeIds *[]uint 不包含的区域类型ID
//@return: err error
func (s Post) CodeIsExist(code string, excludeIds *[]uint) bool {
	// 判断名字是否唯一
	var filter = &domain.PostSearchRequest{
		Code:     code,
		PageInfo: request.PageInfo{PageSize: 1},
	}
	if excludeIds != nil {
		filter.IDNotIn = *excludeIds
	}
	_, total := s.GetCount(filter)
	return total > 0
}
