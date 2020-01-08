package shop

import (
	"time"
)

type PlugIns struct {
	Id             int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	Name           string    `xorm:"not null default '' comment('插件名称') VARCHAR(60)"`
	Identification string    `xorm:"not null comment('标识') VARCHAR(60)"`
	Sort           int       `xorm:"comment('排序') INT(11)"`
	Status         int       `xorm:"not null default 0 comment('状态 0：未启用 1：启用') TINYINT(4)"`
	AddTime        time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
