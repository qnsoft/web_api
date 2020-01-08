package shop

import (
	"time"
)

type FinanceConfig struct {
	Id               int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId          int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	PlugInsId        int       `xorm:"not null default 0 comment('插件id') INT(11)"`
	MinCz            string    `xorm:"default 0.00 comment('最小充值金额') DECIMAL(12,2)"`
	MinAmount        string    `xorm:"not null default 0.00 comment('最少提现金额') DECIMAL(12,2)"`
	MaxAmount        string    `xorm:"not null default 0.00 comment('最大提现金额') DECIMAL(12,2)"`
	ServiceCharge    string    `xorm:"not null default 0.00 comment('手续费') DECIMAL(12,2)"`
	Unit             string    `xorm:"not null default '元' comment('小程序钱包单位') VARCHAR(50)"`
	ModifyDate       time.Time `xorm:"comment('修改时间') TIMESTAMP"`
	Multiple         int       `xorm:"default 0 comment('提现倍数') INT(6)"`
	TransferMultiple int       `xorm:"default 0 comment('转账倍数') INT(6)"`
	CzMultiple       int       `xorm:"default 100 comment('充值倍数') INT(6)"`
}
