package connect

import (
	"MySystem/conf"
	"MySystem/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitMySQL(cfg *setting.MySQLConfig) (DB *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 模型Model绑定数据库的private_account、business_account表go
	DB.AutoMigrate(
		&conf.JobInfo{},
	)
	return DB, DB.DB().Ping()
}
