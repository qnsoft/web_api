package shop

import (
	"time"
)

type Guide struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Image   string    `xorm:"not null default '' comment('图片') VARCHAR(200)"`
	Source  int       `xorm:"default 0 comment('来源 1:小程序 2:APP') TINYINT(4)"`
	Type    int       `xorm:"comment('类型') INT(11)"`
	Sort    int       `xorm:"comment('排序') INT(11)"`
	AddDate time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
