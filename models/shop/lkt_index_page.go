package shop

import (
	"time"
)

type IndexPage struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Value     string    `xorm:"not null comment('值') TEXT"`
	Name      string    `xorm:"not null comment('模块名称') VARCHAR(100)"`
	Sort      int       `xorm:"default 10 comment('排序') INT(11)"`
	StoreType string    `xorm:"not null default '0' comment('软件类型') VARCHAR(11)"`
	Style     string    `xorm:"default '0' comment('样式类型') VARCHAR(10)"`
	AddDate   time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
