package domain

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"ruoyi-go/common/config"
	"time"
)

var (
	DB *gorm.DB
)

func InitDB(r *gin.Engine) {
	log.Info("app system db init")
	prefix := config.String("database.admin.prefix")
	dsn := config.String("database.admin.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix, // 表名前缀，`User` 的表名应该是 `prefix_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `prefix_user`
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	maxIdle := config.Int("database.admin.max-idle")         // 连接池最大闲置的连接数
	maxOpen := config.Int("database.admin.max-open")         // 连接池最大打开的连接数
	maxLifetime := config.Int("database.admin.max-lifetime") // 连接对象可重复使用的时间长度(秒)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpen)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime))

	autoMigrate(db)

	if config.Bool("database.admin.debug") {
		db = db.Debug()
	}
	DB = db
}

func autoMigrate(db *gorm.DB) {
	// migrate model
	err := db.AutoMigrate(
		User{},       // 后台用户表
		Department{}, // 部门表
		Post{},       // 职位表
	)
	if err != nil {
		log.Fatalf("migrate error: %v", err.Error())
		os.Exit(0)
	}
}
