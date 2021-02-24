package dao

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
)
//数据库连接
var DB *sql.DB
func InitDB()(err error){
	dsn:="root:123456@tcp(127.0.0.1:3306)/bilibili?charset=utf8mb4&parseTime=True"
	DB,err=sql.Open("mysql",dsn)
	if err!=nil{
		return err
	}
	err=DB.Ping()//校对密码是否正确
	if err!=nil{
		return err
	}
	return nil
}
func Close(){
	DB.Close()
}