package module

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config/xml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
	//db连接信息
	dbinfo string

)
type Database struct {
	User string
	Password string
	Host string
	Database string
}
func NewDatabase() string{
	database :=   Database{
		User:     beego.AppConfig.String("mysqluser"),
		Password: beego.AppConfig.String("mysqlpass"),
		Host:     beego.AppConfig.String("mysqlurl"),
		Database: beego.AppConfig.String("mysqldb"),
	}
	return  fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",database.User,database.Password,database.Host,database.Database)
}

func init() {
	//注册数据库 mysql是数据库类型
	//拼接数据库连接信息
	dbinfo = NewDatabase()
	//初始化db
	Db,err  = gorm.Open("mysql", dbinfo)
	if err != nil {
		fmt.Println("mysql打开失败",err)
		return
	}
	//创建表关联user结构体
	Db.AutoMigrate(&User{})
}




