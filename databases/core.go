package database

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ContentMysql(tableName string) *gorm.DB {
	//读取.ini里面的数据库配置

	//dsn := "root:123456@tcp(192.168.0.6:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := fmt.Sprintf("root:ql2252528@tcp(172.30.15.90:3306)/bill202212?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "root:ql2252528@tcp(127.0.0.1:3306)/" + tableName + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true, //打印sql
		SkipDefaultTransaction: true,
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}
	return DB
}
