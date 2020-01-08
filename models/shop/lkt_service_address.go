package shop

type ServiceAddress struct {
	Id        int    `xorm:"not null pk autoincr comment('售后地址id') INT(11)"`
	StoreId   int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name      string `xorm:"not null comment('收货人姓名') VARCHAR(50)"`
	Tel       string `xorm:"not null comment('联系电话') CHAR(20)"`
	Sheng     int    `xorm:"not null default 0 comment('省id') INT(11)"`
	City      int    `xorm:"not null default 0 comment('市id') INT(11)"`
	Quyu      int    `xorm:"not null default 0 comment('区域id') INT(11)"`
	Address   string `xorm:"comment('收货地址（不加省，市，区）') VARCHAR(255)"`
	AddressXq string `xorm:"not null comment('省市区+详细地址') VARCHAR(255)"`
	Code      int    `xorm:"not null default 0 comment('邮政编号') INT(11)"`
	Uid       string `xorm:"not null default '0' comment('代表该地址是售后地址') VARCHAR(50)"`
	IsDefault int    `xorm:"not null default 0 comment('是否为默认收货地址 1.是  0.不是') TINYINT(2)"`
}
