package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

/**
notes: 配置基础
*/

//声明结构

type DebugStruct bool
type DatabaseStruct struct {
	Master ConnectorStruct
	Slave  ConnectorStruct
}
type ConnectorStruct struct {
	Driver string
	Dsn    string
}

//初始变量

var Debug DebugStruct
var RouteDebug DebugStruct
var SqlDebug DebugStruct
var Database = DatabaseStruct{
	Master: ConnectorStruct{},
	Slave:  ConnectorStruct{},
}

//初始化-赋值

func init() {

	//读取根目录 .env 文件
	path, _ := os.Getwd()
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//系统调试
	if os.Getenv("APP_DEBUG") == "1" {
		Debug = true
	}
	//路由调试
	if os.Getenv("ROUTE_DEBUG") == "1" {
		RouteDebug = true
	}
	//SQL调试
	if os.Getenv("SQL_DEBUG") == "1" {
		SqlDebug = true
	}

	//主库连接
	Database.Master = ConnectorStruct{
		Driver: os.Getenv("DB_DRIVER"),
		Dsn:    os.Getenv("MASTER_DSN_1"),
	}
	//丛库连接
	Database.Slave = ConnectorStruct{
		Driver: os.Getenv("DB_DRIVER"),
		Dsn:    os.Getenv("SLAVE_DSN_1"),
	}

}
