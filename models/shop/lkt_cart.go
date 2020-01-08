package shop

import (
	"time"
)

type Cart struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Token      string    `xorm:"comment('token') VARCHAR(255)"`
	UserId     string    `xorm:"not null comment('用户id') VARCHAR(20)"`
	Uid        string    `xorm:"comment('微信id') VARCHAR(32)"`
	GoodsId    int       `xorm:"comment('产品id') INT(11)"`
	GoodsNum   int       `xorm:"comment('数量') INT(11)"`
	CreateTime time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	SizeId     string    `xorm:"not null default '' comment('商品属性id') CHAR(15)"`
}
