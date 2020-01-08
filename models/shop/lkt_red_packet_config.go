package shop

type RedPacketConfig struct {
	Id       int    `xorm:"not null pk autoincr comment('ID') INT(11)"`
	StoreId  int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Ordinary string `xorm:"comment('普通红包数据') TEXT"`
	Rand     string `xorm:"comment('拼手气红包数据') TEXT"`
}
