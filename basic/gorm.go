package basic

import (
	"code.byted.org/gopkg/logs"
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func initDb() {
	dbGorm, err := gorm.Open("mysql", "appclouddev:password@tcp(10.8.121.175:3306)/appcloud_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logs.Error("init db error:%v", err)
		panic(fmt.Sprint("Init db: connect db error, %v", err))
	}
	db = dbGorm
}
func TestGormConnect() {
	initDb()
	defer db.Close()
	tx := db.Begin()
	//err := Db.CreateTable(&Product{Code: "L1212", Price: 1000}).Error
	err := tx.Table("products").Where("code=1").Update("code","abc1").Error
	fmt.Println("create table  err,err=", err)
	err = db.Table("products").Create(&Product{Code: "L1211", Price: 1000}).Error
	fmt.Println("create err,err=", err)
	tx.Commit()
}
