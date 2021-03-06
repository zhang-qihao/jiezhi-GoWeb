package main

import (
	"MySystem/dao"
	"MySystem/models"
	"MySystem/routers"
	"MySystem/setting"
	"fmt"
)

func main() {

	// 加载配置文件
	configPath := "conf/config.ini"
	if err := setting.Init(configPath); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库,在全局变量DB中
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 数据库表与模型对应
	dao.DB.AutoMigrate(
		&models.JobInfos{},
		&models.PersonPost{},
	)

	// 注册路由并执行对应的操作
	r := routers.SetupRouter()
	err = r.Run(":8080") //启动服务器 端口为8080
	if err != nil {
		panic(err)
	} //返回错误信息

}
