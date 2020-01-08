package shop

type BargainGoods struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	StoreId    int    `xorm:"not null default 0 index INT(11)"`
	GoodsId    int    `xorm:"not null default 0 index INT(11)"`
	AttrId     int    `xorm:"not null comment('产品属性id') INT(11)"`
	MinPrice   string `xorm:"not null default 0.00 comment('最低价') DECIMAL(11,2)"`
	BeginTime  string `xorm:"not null default '0' comment('活动开始时间') CHAR(20)"`
	EndTime    string `xorm:"not null default '0' comment('活动结束时间') CHAR(20)"`
	Buytime    int    `xorm:"not null default 0 comment('购买时间') INT(11)"`
	ManNum     int    `xorm:"not null default 0 comment('砍价人数 0为不限制') SMALLINT(6)"`
	StatusData string `xorm:"not null default '' comment('砍价方式数据') VARCHAR(255)"`
	IsDelete   int    `xorm:"not null default 0 comment('是否上线:0下线   1上线') SMALLINT(6)"`
	IsShow     int    `xorm:"not null default 1 comment('是否显示:0,不显示   1,显示') SMALLINT(6)"`
	Addtime    string `xorm:"not null default '0' CHAR(20)"`
}
