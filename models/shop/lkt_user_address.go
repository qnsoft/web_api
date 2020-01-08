package shop

type UserAddress struct {
	Id        int    `xorm:"not null pk autoincr comment('地址id') INT(11)"`
	StoreId   int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name      string `xorm:"not null comment('收货人') VARCHAR(10)"`
	Tel       string `xorm:"not null comment('联系方式') CHAR(15)"`
	Sheng     string `xorm:"not null default '0' comment('省id') VARCHAR(200)"`
	City      string `xorm:"not null default '0' comment('市id') VARCHAR(200)"`
	Quyu      string `xorm:"not null default '0' comment('区域id') VARCHAR(200)"`
	Address   string `xorm:"not null comment('收货地址（不加省市区）') VARCHAR(200)"`
	AddressXq string `xorm:"not null comment('省市区+详细地址') VARCHAR(255)"`
	Code      int    `xorm:"not null default 0 comment('邮政编号') INT(11)"`
	Uid       string `xorm:"not null default '0' comment('用户ID') VARCHAR(32)"`
	IsDefault int    `xorm:"not null default 0 comment('是否默认地址 1默认') TINYINT(2)"`
}
