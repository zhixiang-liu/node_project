package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nodejs_project/static"
)

var DB *gorm.DB

func OpenToMysql() {
	username := "root"        //账号
	password := "l1074822002" //密码
	host := "127.0.0.1"       //数据库地址，可以是Ip或者域名
	port := 3306              //数据库端口
	Dbname := "node_db"       //数据库名
	timeout := "10s"          //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}
	// 连接成功
	DB = db
	err = db.AutoMigrate(
		&static.UserInfo{},
		&static.LogInfo{},
		&static.TestMsg{},
	)
	if err != nil {
		fmt.Println(err)
	}

}
