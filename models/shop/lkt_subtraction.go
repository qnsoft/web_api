package shop

import (
	"time"
)

type Subtraction struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId     int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Goods       string    `xorm:"comment('所选商品') TEXT"`
	Status      int       `xorm:"not null default 2 comment('活动状态 1.开启 2.关闭') TINYINT(3)"`
	Subtraction string    `xorm:"comment('满减') TEXT"`
	Startdate   string    `xorm:"comment('活动开始时间') CHAR(20)"`
	Enddate     string    `xorm:"comment('活动结束时间') CHAR(20)"`
	AddDate     time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
}
