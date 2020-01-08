package shop

type GroupProduct struct {
	Id         int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	AttrId     int    `xorm:"not null comment('规格id') INT(11)"`
	ProductId  int    `xorm:"comment('产品id') INT(11)"`
	GroupLevel string `xorm:"not null default '' comment('拼团等级价格参数') VARCHAR(200)"`
	GroupData  string `xorm:"not null default '' comment('拼团参数数据') VARCHAR(300)"`
	GStatus    int    `xorm:"not null default 1 comment('活动状态: 1--未开始 2--活动中 3--已结束') TINYINT(4)"`
	IsShow     int    `xorm:"not null default 1 comment('是否显示') TINYINT(4)"`
}
