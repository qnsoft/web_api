package shop

import (
	"time"
)

type SignRecord struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId    string    `xorm:"not null comment('用户ID') VARCHAR(15)"`
	SignScore int       `xorm:"not null default 0 comment('签到积分') INT(11)"`
	Record    string    `xorm:"comment('事件') CHAR(20)"`
	SignTime  time.Time `xorm:"comment('签到时间') TIMESTAMP"`
	Type      int       `xorm:"not null default 0 comment('类型: 0:签到 1:消费 2:首次关注得积分 3:转积分给好友 4:好友转积分 5:系统扣除 6:系统充值 7:抽奖') INT(4)"`
	Recovery  int       `xorm:"default 0 comment('是否删除 0.未删除 1.
删除') TINYINT(2)"`
}
