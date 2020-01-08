package shop

type Setscore struct {
	Id       int `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId  int `xorm:"not null default 0 comment('商城id') INT(11)"`
	Lever    int `xorm:"comment('等级') SMALLINT(6)"`
	Ordernum int `xorm:"comment('订单金额') INT(11)"`
	Scorenum int `xorm:"comment('可抵用消费金额') INT(11)"`
}
