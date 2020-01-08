package shop

type MchConfig struct {
	Id         int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Logo       string `xorm:"comment('店铺默认logo') VARCHAR(255)"`
	Settlement string `xorm:"not null default '0' comment('结算方式') VARCHAR(50)"`
	MinCharge  string `xorm:"default 0.00 comment('最低提现金额

') DECIMAL(12,2)"`
	MaxCharge string `xorm:"default 0.00 comment('最大提现金额

') DECIMAL(12,2)"`
	ServiceCharge string `xorm:"default 0.00 comment('手续费') DECIMAL(12,2)"`
	Illustrate    string `xorm:"comment('提现说明') TEXT"`
	Agreement     string `xorm:"comment('入驻协议') TEXT"`
}
