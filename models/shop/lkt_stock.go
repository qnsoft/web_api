package shop

import (
	"time"
)

type Stock struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId     int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	ProductId   int       `xorm:"not null default 0 comment('商品id') INT(11)"`
	AttributeId int       `xorm:"not null default 0 comment('属性id') INT(11)"`
	FlowingNum  int       `xorm:"not null default 0 comment('入库/出库') INT(11)"`
	Type        int       `xorm:"not null default 0 comment('类型 0.入库 1.出库 2.预警') TINYINT(4)"`
	UserId      string    `xorm:"not null default '' comment('购买方') VARCHAR(100)"`
	AddDate     time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
