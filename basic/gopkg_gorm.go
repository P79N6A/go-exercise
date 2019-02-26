package basic

import (
	"code.byted.org/gopkg/dbutil/conf"
	"code.byted.org/gopkg/dbutil/gormdb"
	//"github.com/jinzhu/gorm"

	"code.byted.org/gopkg/gorm"
	"code.byted.org/gopkg/logs"
	"fmt"
)

var Db *gorm.DB
//var Dbhandler *gormdb.DBHandler

func initGopkgGormDb() {
	dbHandler := gormdb.NewDBHandler()
	err := dbHandler.ConnectDB(&conf.DBOptional{
		DriverName:   "mysql",
		User:         "appclouddev",
		Password:     "password",
		DBName:       "appcloud_test",
		DBCharset:    "utf8mb4",
		DBHostname:   "10.8.121.175",
		DBPort:       "3306",
		Timeout:      "100ms",
		ReadTimeout:  "2.0s",
		WriteTimeout: "5.0s",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
	})
	if err != nil {
		logs.Error("init db error:%v", err)
		panic(fmt.Sprintf("Init db: connect db error, %v", err))
	}
	if err != nil {
		logs.Error("init db error:%v", err)
		panic(fmt.Sprint("Init db: connect db error, "))
	}
	dbClient, err := dbHandler.GetConnection()
	if err != nil {
		panic(fmt.Sprintf("Init db: connect db error, "))
	}
	Db = dbClient

	//Dbhandler = dbHandler
}
func TestGopkgGormConnect() {
	initGopkgGormDb()
	//defer db.Close()
	//err := db.CreateTable(&Product{Code: "L1212", Price: 1000}).Error
	tx := Db.Begin()
	//err := Db.CreateTable(&Product{Code: "L1212", Price: 1000}).Error
	//err := tx.Create(&Product{Code: "L1212", Price: 1000}).Error
	err := tx.Table("product").Where("code='abc'").Update("code","abc1").Error
	fmt.Println("create table  err,err=", err)
	err = Db.Table("product").Create(&Product{Code: "L1211", Price: 1000}).Error
	fmt.Println("create err,err=", err)
	tx.Commit()
}
