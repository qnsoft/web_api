package shop

import (
	"time"
)

type ProductConfig struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Config  string    `xorm:"comment('配置') TEXT"`
	AddDate time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
