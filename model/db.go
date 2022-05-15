package model

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

/**
 * 配置连接数据库
 */

var db *gorm.DB
var err error

func InitDb() {
	// 连接数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 设置表名不为负数
			SingularTable: true,
		},
		// 禁用自动创建数据库外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	//defer func(sqlDB *sql.DB) {
	//	err := sqlDB.Close()
	//	if err != nil {
	//		fmt.Println("关闭数据库失败", err)
	//	}
	//}(sqlDB)
}
