package shop

import (
	"time"
)

type DistributionWithdraw struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId  string    `xorm:"not null default '' comment('用户id') VARCHAR(20)"`
	Name    string    `xorm:"not null default '' comment('名称') VARCHAR(20)"`
	WxId    string    `xorm:"comment('微信id') VARCHAR(50)"`
	Mobile  string    `xorm:"comment('手机') VARCHAR(20)"`
	BankId  int       `xorm:"not null comment('银行卡id') INT(11)"`
	Money   string    `xorm:"not null default 0.00 comment('提现金额') DECIMAL(12,2)"`
	ZMoney  string    `xorm:"not null default 0.00 comment('剩余金额') DECIMAL(12,2)"`
	SCharge string    `xorm:"not null default 0.00 comment('手续费') DECIMAL(12,2)"`
	Status  int       `xorm:"not null default 0 comment('状态 0：审核中 1：审核通过 2：拒绝') TINYINT(4)"`
	AddDate time.Time `xorm:"comment('申请时间') TIMESTAMP"`
	Refuse  string    `xorm:"comment('拒绝原因') TEXT"`
	IsMch   int       `xorm:"default 0 comment('是否是店铺提现') INT(2)"`
}
