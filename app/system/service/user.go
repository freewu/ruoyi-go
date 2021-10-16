package service

import "ruoyi-go/app/system/domain"

type User struct {}

//@author: [freewu](http://git.yibianli.com/freewu)
//@function: GetList
//@description: 获取用户列表数据
//@param: searchParams *domain.UserSearchRequest
//@return: err error, list interface{}, total int64
func (s User) GetList(params *domain.UserSearchRequest) (err error, list []domain.User, total int64) {
	return err,list,total
}
