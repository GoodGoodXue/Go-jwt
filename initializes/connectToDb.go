package initializes

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/web?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数形式，例如User的表名默认是users,
		}})

	if err != nil {
		panic("Failed to connect database,err" + err.Error())
	}
}
