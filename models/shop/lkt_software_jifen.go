package shop

type SoftwareJifen struct {
	Id           int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId      int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Jifennum     int    `xorm:"comment('首次关注小程序积分') INT(11)"`
	Switch       int    `xorm:"not null default 0 comment('是否开启积分转让（0.关闭 1.开启）') INT(11)"`
	Rule         string `xorm:"comment('积分规则') TEXT"`
	Xiaofeibili  int    `xorm:"not null default 10 comment('积分消费比例') INT(11)"`
	Xiaofeiguize string `xorm:"comment('积分消费规则') TEXT"`
}
