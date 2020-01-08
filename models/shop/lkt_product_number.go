package shop

import (
	"time"
)

type ProductNumber struct {
	Id            int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId       int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	MchId         string    `xorm:"comment('店铺id') CHAR(15)"`
	Operation     string    `xorm:"comment('操作人账号') VARCHAR(150)"`
	ProductNumber string    `xorm:"not null comment('商品编号') VARCHAR(100)"`
	Status        int       `xorm:"not null default 1 comment('状态：1.使用 2.撤销') TINYINT(3)"`
	Addtime       time.Time `xorm:"comment('时间') TIMESTAMP"`
}
