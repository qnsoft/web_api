package shop

import (
	"time"
)

type Draw struct {
	Id             int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId        int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name           string    `xorm:"comment('活动名称') VARCHAR(255)"`
	DrawBrandid    int       `xorm:"not null comment('该抽奖所参与商品ID') INT(11)"`
	FoundTime      time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	StartTime      time.Time `xorm:"comment('开始时间') DATETIME"`
	EndTime        time.Time `xorm:"comment('结束时间') DATETIME"`
	Num            int       `xorm:"comment('每个团所需人数') INT(11)"`
	SpellingNumber int       `xorm:"comment('可抽中奖次数（默认为1）') INT(11)"`
	CollageNumber  int       `xorm:"comment('最少开奖团数（默认为1）') INT(11)"`
	State          int       `xorm:"default 0 comment('该团的状态（默认为0：未开始 ，1：进行中 ， 2：已结束）') INT(11)"`
	Price          string    `xorm:"comment('抽奖金额') DECIMAL(10,2)"`
	Cishu          int       `xorm:"comment('次数 ') INT(11)"`
	Type           string    `xorm:"comment('备注') VARCHAR(255)"`
}
