# 项目说明
> 使用 gin + gorm 实现 ruoyi 的功能，做为一个 go 后台管理系统快速开发框架

# 内置功能

* [待实现]用户管理：用户是系统操作者，该功能主要完成系统用户配置。
* [待实现]部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
* [待实现]岗位管理：配置系统用户所属担任职务。
* [待实现]菜单管理：配置系统菜单，操作权限，按钮权限标识等。
* [待实现]角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
* [待实现]字典管理：对系统中经常使用的一些较为固定的数据进行维护。
* [待实现]参数管理：对系统动态配置常用参数。
* [待实现]操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
* [待实现]登录日志：系统登录日志记录查询包含登录异常。
* [待实现]在线用户：当前系统中活跃用户状态监控。
* [待实现]定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
* [待实现]代码生成：前后端代码的生成。
* [待实现]服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。

# 项目启动
```bash
# 启动

  go run main.go

```

# 使用组件
```
github.com/gin-gonic/gin 路由分发
gorm.io/gorm db操作
github.com/sirupsen/logrus 日志
github.com/spf13/viper 配置文件
github.com/go-playground/validator/v10 结构体参数验证
github.com/swaggo/swag 接口文档
```

# 接口文档生成
```bash
# 生成文档
swag init -o ./app/swagger

# 访问
http://127.0.0.1:8080/swagger.json
```

# swag 安装
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```