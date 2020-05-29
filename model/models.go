package model

import (
	"fmt"
	"time"

	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"

	// mysql原语
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/makiko-fly/logrus"
)

// User 数据库表的原型
type User struct {
	ID       int64
	UserName string
	Password string
	Role     string
}

// UserAccess 数据库表的原型
type UserAccess struct {
	UserRole string
	Resource string
	Action   string
}

// TestDb 全局的DB
var TestDb *gorm.DB

// GormAdapter 调整鉴权规则的数据库，每次调整完成后，需要LoadPolicy一下
var GormAdapter *gormadapter.Adapter

// InitDb 数据库连接初始化
func InitDb() {
	TestDb = createDb("root", "123456", "127.0.0.1", "3306", "csco", "utf8", 200, 30)
	// gormadapter 会自动在数库名参数后面加上casbin后缀，所有需要新建一个数据库dbName+casbin
	gadapter, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/csco", false)
	if err != nil || gadapter == nil {
		logrus.Errorln("mysql adapter init failure")
		return
	}
	GormAdapter = gadapter
	migrate()
	logrus.Infoln("Finished database migration!")
}

func createDb(user, password, host, port, dbName, charset string, maxLifeTime, maxConn int) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		user,
		password,
		host,
		port,
		dbName,
		charset,
		/* url.QueryEscape("Asia/Shanghai") */)
	if gormDb, err := gorm.Open("mysql", connStr); err != nil {
		logrus.Panicf(fmt.Sprintf("Failed to connect to database with connection str: [%v], err: %v", connStr, err))
	} else {
		gormDb.DB().SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)
		gormDb.DB().SetMaxOpenConns(maxConn)
		return gormDb
	}
	return nil
}

func migrate() {
	TestDb.AutoMigrate(&User{})
	TestDb.AutoMigrate(&UserAccess{})
}
