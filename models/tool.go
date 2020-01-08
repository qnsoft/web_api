package model

import (
	_ "database/sql"
	_ "time"

	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
)

func TableName(name string) string {
	return beego.AppConfig.String("database::db_prefix") + name
}
