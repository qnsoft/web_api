package shop

type PaymentConfig struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	StoreId    int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Pid        int    `xorm:"not null comment('支付方式id') INT(11)"`
	Status     int    `xorm:"not null default 1 comment('状态 0启用 1禁用') TINYINT(1)"`
	ConfigData string `xorm:"comment('配置参数,json数据对象') TEXT"`
}
