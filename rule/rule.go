package rule

import (
	"github.com/casbin/casbin"
	"just.for.test/rbacdemo/model"
)

var (
	// AuthEnforce 全局鉴权中心
	AuthEnforce *casbin.Enforcer
)

// Init 初始化
func Init() {
	// 自文件中读取poliy
	// e, err := casbin.NewEnforcer("./rbac/rbac.conf", "./rbac/rbac.csv")
	// 自数据库中读取poliy
	e, err := casbin.NewEnforcer("./rbac/rbac.conf", model.GormAdapter)
	if err != nil {
		panic(err)
	}
	AuthEnforce = e
	if err := AuthEnforce.LoadPolicy(); err != nil {
		panic(err)
	}
}
