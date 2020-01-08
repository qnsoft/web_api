package shop

import (
	"time"
)

type DrawUser struct {
	Id            int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId       int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	DrawId        int       `xorm:"comment('拼团ID') INT(11)"`
	UserId        string    `xorm:"comment('用户ID') VARCHAR(30)"`
	Time          time.Time `xorm:"comment('用户参团时间') TIMESTAMP"`
	Role          string    `xorm:"default '0' comment('用户角色（默认 0：团长  userid:该用户分享进来的用户）') VARCHAR(30)"`
	LotteryStatus int       `xorm:"default 0 comment('抽奖状态（0.参团中 1.待抽奖 2.参团失败 3.抽奖失败 4.抽奖成功）') INT(11)"`
}
