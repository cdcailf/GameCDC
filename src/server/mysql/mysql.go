package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
)

var (
	db *gorm.DB
)

func OpenDB() {
	fmt.Println("mysqldb->open db")
	db1, err := gorm.Open("mysql", "xxx:xxx@tcp(xxx.xxx.xxx.xxx:3306)/game?parseTime=true")
	if err != nil {
		panic("connect db error")
	}
	db = db1
}

func MysqlDB() *gorm.DB {
	return db
}
