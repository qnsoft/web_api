package shop

import (
	"time"
)

type AuctionPromise struct {
	Id        int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	UserId    string    `xorm:"comment('用户id') VARCHAR(50)"`
	Promise   string    `xorm:"comment('保证金额') DECIMAL(12,2)"`
	AddTime   time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	AId       int       `xorm:"comment('竞拍商品id') INT(11)"`
	TradeNo   string    `xorm:"comment('订单编号') CHAR(32)"`
	IsPay     int       `xorm:"comment('是否成功支付，0失败，1成功') TINYINT(2)"`
	IsBack    int       `xorm:"comment('是否退款成功  0-成功 1-失败') TINYINT(2)"`
	StoreId   int       `xorm:"comment('商城id') INT(11)"`
	Type      int       `xorm:"not null comment('支付方式 1微信支付 2-钱包支付') TINYINT(2)"`
	AllowBack int       `xorm:"default 1 comment('是否符合退款标准  0：不符合 1：符合') TINYINT(2)"`
}
