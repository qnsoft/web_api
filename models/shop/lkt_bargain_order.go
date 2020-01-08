package shop

type BargainOrder struct {
	Id            int    `xorm:"not null pk autoincr INT(11)"`
	StoreId       int    `xorm:"not null default 0 index INT(11)"`
	UserId        string `xorm:"not null default '0' index CHAR(11)"`
	OrderNo       string `xorm:"not null default '0' VARCHAR(255)"`
	AttrId        int    `xorm:"not null default 0 INT(11)"`
	OriginalPrice string `xorm:"not null default 0.00 comment('剩余价格') DECIMAL(10,2)"`
	MinPrice      string `xorm:"not null default 0.00 comment('商品最低价') DECIMAL(10,2)"`
	GoodsId       int    `xorm:"not null INT(11)"`
	Status        int    `xorm:"not null default 0 comment('状态 0--进行中 1--成功 2--失败') INT(11)"`
	IsDelete      int    `xorm:"not null default 0 INT(11)"`
	Addtime       int    `xorm:"not null default 0 INT(11)"`
	StatusData    string `xorm:"not null default '' comment('砍价方式数据') VARCHAR(255)"`
}
