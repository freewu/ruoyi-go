package service

import (
	"gorm.io/gorm"
	"ruoyi-go/app/system/domain"
	"time"
)

type User struct {}

//@author: [bluefrog](https://github.com/freewu)
//@function: GetList
//@description: 获取用户列表数据
//@param: searchParams *domain.UserSearchRequest
//@return: err error, list interface{}, total int64
func (s User) GetList(searchParams *domain.UserSearchRequest) (err error, list []domain.User, total int64) {
	// 创建db
	db := domain.DB.Model(&domain.User{})
	// 条件过滤
	db = s.parseFilter(db,searchParams)
	// 统计数据
	err = db.Count(&total).Error
	// 如果数据 0,也没有必要处理以下动作了
	if total > 0 {
		// 排序
		orderBy := "id desc"
		if searchParams.OrderBy != "" {
			db.Order(searchParams.OrderBy)
		}
		db.Order(orderBy)
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
//@description: 用户列表搜索条件过滤
//@param: db *gorm.DB
//@param: searchParams *domain.UserSearchRequest
//@return: *gorm.DB
func (s User) parseFilter(db *gorm.DB, searchParams *domain.UserSearchRequest) *gorm.DB {
	// 条件过滤
	return db
}

//@author: [bluefrog](https://github.com/freewu)
//@function: Create
//@description: 添加用户数据
//@param: data *domain.UserAddRequest
//@return: error
func (s User) Create(data *domain.UserAddRequest) error {
	// todo 判断用户名是否已存在
	// todo 判断手机号是否已存在
	// todo 判断邮箱是否已存在
	user := new(domain.User)
	user.Username = data.Username
	user.Mobile = data.Mobile
	user.Nickname = data.Nickname
	user.Birthday = data.Birthday
	user.Status = data.Status
	user.Email = data.Email
	user.Gender = data.Gender
	user.Avatar = data.Avatar
	user.DepartmentID = data.DepartmentID
	user.Remark = data.Remark
	user.IsAdmin = data.IsAdmin
	user.RegisterTime = time.Now() // 注册时间

	return domain.DB.Create(&user).Error
}

//@author: [bluefrog](https://github.com/freewu)
//@function: Update
//@description: 修改区域
//@param: data *model.UserEditRequest
//@return: err error
func (s User) Update(data *domain.UserEditRequest) (err error) {
	// todo 判断用户名是否已存在
	// todo 判断手机号是否已存在
	// todo 判断邮箱是否已存在

	record := make(map[string]interface{})
	record["Username"] = data.Username
	record["Mobile"] = data.Mobile
	record["Nickname"] = data.Nickname
	record["Birthday"] = data.Birthday
	record["Status"] = data.Status
	record["Email"] = data.Email
	record["Gender"] = data.Gender
	record["Avatar"] = data.Avatar
	record["DepartmentID"] = data.DepartmentID
	record["Remark"] = data.Remark
	record["IsAdmin"] = data.IsAdmin

	return s.UpdateByMap(data.ID, record)
}

//@author: [bluefrog](https://github.com/freewu)
//@function: UpdateByMap
//@description: 更新指定ID的用户信息
//@param: id uint
//@param: updateMap map[string]interface{}
//@return: err error
func (s User) UpdateByMap(id uint, updateMap map[string]interface{}) (err error) {
	if err := domain.DB.Model(&domain.User{}).Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return err
	}
	return err
}

//@author: [bluefrog](https://github.com/freewu)
//@function: Delete
//@description: 删除用户
//@param: ids []string
//@return: err error
func (s User) Delete(ids []string) (err error) {
	// 删除区域
	if err = domain.DB.Where("id in (?) ",ids).Unscoped().Delete(&domain.User{}).Error; err != nil {
		return err
	}
	return nil
}

//@author: [freewu](http://git.yibianli.com/freewu)
//@function: Detail
//@description: 获取指定ID用户详情
//@param: id uint
//@return: *domain.User,err error
func (s User) Detail(id uint) (*domain.User, error) {
	var detail domain.User
	db := domain.DB.Where("id = ?", id)
	if err := db.First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}