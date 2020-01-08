package shop

type TaobaokeConfig struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	StoreId     int    `xorm:"default 1 INT(11)"`
	AppKey      string `xorm:"VARCHAR(200)"`
	AppSecret   string `xorm:"VARCHAR(200)"`
	Pid         string `xorm:"VARCHAR(200)"`
	Many        int    `xorm:"default 0 comment('-1直推 0多层级') INT(4)"`
	Comm        string `xorm:"default '' comment('分佣比例') VARCHAR(255)"`
	InviterCode string `xorm:"comment('邀请码') VARCHAR(255)"`
	Accesstoken string `xorm:"comment('accesstoken ') VARCHAR(255)"`
}
