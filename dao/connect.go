package dao

import (
	"MySystem/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	DB = db.Debug()
	// 模型Model绑定数据库的private_account、business_account表go
	return DB.DB().Ping()
}
