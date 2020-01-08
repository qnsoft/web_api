package shop

type GroupConfig struct {
	Id         int `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int `xorm:"not null default 0 comment('商城id') INT(11)"`
	Refunmoney int `xorm:"not null comment('退款方式: 1,自动 2,手动') SMALLINT(6)"`
}
