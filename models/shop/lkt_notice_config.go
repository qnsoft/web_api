package shop

import (
	"time"
)

type NoticeConfig struct {
	Id         int       `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	CName      string    `xorm:"comment('中文名称') VARCHAR(50)"`
	EName      string    `xorm:"VARCHAR(255)"`
	StockId    string    `xorm:"comment('微信模板库id') VARCHAR(50)"`
	StockKey   string    `xorm:"comment('组合关键词') VARCHAR(255)"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
}
