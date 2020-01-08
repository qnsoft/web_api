package shop

import (
	"time"
)

type MoneyPoint struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId     int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId      string    `xorm:"not null comment('用户ID') VARCHAR(15)"`
	Type        int       `xorm:"not null default 0 comment('奖金类型') TINYINT(4)"`
	AddDate     time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Money       string    `xorm:"not null default 0.00 comment('应发金额') DECIMAL(12,2)"`
	SMoney      string    `xorm:"not null default 0.00 comment('实发金额') DECIMAL(12,2)"`
	FxMoney     string    `xorm:"not null default 0.00 comment('手续费一') DECIMAL(12,2)"`
	TaxMoney    string    `xorm:"not null default 0.00 comment('手续费二') DECIMAL(12,2)"`
	OMoney      string    `xorm:"not null default 0.00 comment('手续费三') DECIMAL(12,2)"`
	SiMoney     string    `xorm:"not null default 0.00 comment('手续费四') DECIMAL(12,2)"`
	WMoney      string    `xorm:"not null default 0.00 comment('手续费五') DECIMAL(12,2)"`
	Isf         int       `xorm:"not null default 0 comment('是否延迟发放') TINYINT(4)"`
	FTime       time.Time `xorm:"not null comment('发放时间') TIME"`
	FromOrdersn string    `xorm:"not null comment('订单编号') VARCHAR(15)"`
}
