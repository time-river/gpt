package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"scancenter/config"
)

//连接db
func connectTestDataDb() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=false",
		config.UserName, config.Password,
		config.NetWork, config.Server,
		config.Port, config.DataBase)
	dataBase, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	if _, err := dataBase.DB(); err != nil {
		return nil
	} else {
		return dataBase
	}
}
