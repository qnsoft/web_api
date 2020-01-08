package shop

import (
	"time"
)

type AuctionRecord struct {
	Id         int       `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	AuctionId  int       `xorm:"not null comment('竞拍商品id') INT(11)"`
	AddTime    time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Price      string    `xorm:"comment('竞拍价格') DECIMAL(12,2)"`
	UserId     string    `xorm:"comment('用户id') CHAR(50)"`
	PriceRange string    `xorm:"not null comment('加价幅度') DECIMAL(12,2)"`
}
