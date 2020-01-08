package model

import (
	_ "database/sql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//整个系统启用session
	beego.BConfig.WebConfig.Session.SessionOn = true //也可以通过配置文件sessionon = true 来启用session,在这里我要强制使用
	db_type := beego.AppConfig.String("database::db_type")
	db_path := beego.AppConfig.String("database::db_path")
	// 注册驱动
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", db_type, db_path)
	//orm.RunSyncdb("default", false, true) //没有默认数据库该句不可用
}
