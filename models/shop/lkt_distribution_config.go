package shop

type DistributionConfig struct {
	Id      int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Sets    string `xorm:"TEXT"`
}
