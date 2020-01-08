package shop

type BargainRecord struct {
	Id      int    `xorm:"not null pk autoincr comment('id') index INT(11)"`
	StoreId int    `xorm:"INT(11)"`
	OrderNo string `xorm:"not null default '' comment('砍价订单') CHAR(20)"`
	UserId  string `xorm:"not null comment('用户ID') VARCHAR(15)"`
	Money   string `xorm:"not null default 0.00 comment('金额') DECIMAL(12,2)"`
	AddTime string `xorm:"comment('添加时间') CHAR(20)"`
	Status  int    `xorm:"not null default 0 comment('状态 0:砍价中 1:砍价成功 2:逾期失效 3:生成订单') TINYINT(4)"`
}
