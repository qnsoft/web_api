package shop

import (
	"time"
)

type Banner struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Image     string    `xorm:"not null default '' comment('图片') VARCHAR(200)"`
	OpenType  string    `xorm:"comment('跳转方式') VARCHAR(100)"`
	Url       string    `xorm:"not null default '' comment('链接') VARCHAR(100)"`
	Sort      int       `xorm:"comment('排序') INT(11)"`
	Type      int       `xorm:"default 1 comment('类型 1:小程序 2:app') TINYINT(4)"`
	AddDate   time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	ClassType int       `xorm:"comment('分类：4课堂轮播 默认1') TINYINT(1)"`
}
