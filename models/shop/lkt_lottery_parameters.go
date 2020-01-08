package shop

type LotteryParameters struct {
	Id         int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Parameters string `xorm:"not null comment('参数值') VARCHAR(255)"`
	UserId     string `xorm:"not null comment('用户id') VARCHAR(20)"`
}
