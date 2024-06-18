package entity

import (
	"app/base/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"sync"
	"time"
)

//DOC: https://gorm.io/zh_CN/docs/index.html

/**
notes: 实体模型基础
*/

var connector *gorm.DB                        // 连接实例
var connectorList = make(map[string]*gorm.DB) // 连接实例-列表
var connectorSync = sync.Map{}                // 连接实例-原子标记

type BaseEntityInterface interface {
	Connector(connName string) *gorm.DB
}

type BaseEntityStruct struct {
	BaseEntityInterface
}

// 连接实例-初始化

func (instance *BaseEntityStruct) Connector(connName string) *gorm.DB {

	//获取原子标记
	connSync, _ := connectorSync.Load(connName)
	if connSync == nil {
		//记录原子标记
		connectorSync.Store(connName, "1")

		//公共配置
		var conf = gorm.Config{
			SkipDefaultTransaction: true, //禁用默认事务
			PrepareStmt:            true, //缓存预编译语句
			//命名策略
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   "gin_", // 表名的前缀, 如：User结构，表名为：`bored_users`
				//SingularTable: true,   // 使用单数命名表,如：User结构，在数据库中表名为：user。若不设置，默认为复数。
				//NoLowerCase:   true,   // 禁用snake_casing命名法
			},
		}

		//主库: 注册 源
		var err error
		connector, err = gorm.Open(mysql.Open(config.Database.Master.Dsn), &conf)
		if err != nil {
			log.Fatal(err)
		}
		connDb, _ := connector.DB()
		connDb.SetMaxIdleConns(10)               //设置连接池的最大闲置连接数
		connDb.SetMaxOpenConns(100)              //设置连接池中的最大连接数量
		connDb.SetConnMaxLifetime(1 * time.Hour) //设置连接的最大复用时间

		//主从库: 注册 源/副本
		var traceResolver = false
		if config.Debug || config.SqlDebug {
			traceResolver = true
		}
		var resolverConf = dbresolver.Config{
			Sources:           []gorm.Dialector{mysql.Open(config.Database.Master.Dsn)}, //源
			Replicas:          []gorm.Dialector{mysql.Open(config.Database.Slave.Dsn)},  //副本
			Policy:            dbresolver.RandomPolicy{},                                //源/副本 负载平衡策略
			TraceResolverMode: traceResolver,                                            //打印 源/副本 日志
		}
		var resolverReg = dbresolver.Register(resolverConf).
			SetMaxIdleConns(10). //设置连接池的最大闲置连接数
			SetMaxOpenConns(100). //设置连接池中的最大连接数量
			SetConnMaxIdleTime(time.Hour). //设置连接的最大闲置时间
			SetConnMaxLifetime(1 * time.Hour) //设置连接的最大复用时间
		connector.Use(resolverReg)

		// 是否打开日志
		if config.Debug || config.SqlDebug {
			connector.Debug()
		}

		connectorList[connName] = connector
		log.Println("connectorList[Connector]::", connector) //
	}

	return connectorList[connName]
}
