package shop

import (
	"time"
)

type TwelveDrawField struct {
	Id           int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId      int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	ActivityName string    `xorm:"not null default '' comment('活动名称') VARCHAR(100)"`
	Price        int       `xorm:"default 0 comment('所需金额') INT(10)"`
	Starttime    time.Time `xorm:"comment('活动开始时间') TIMESTAMP"`
	Endtime      time.Time `xorm:"comment('活动结束时间') TIMESTAMP"`
	Rate         string    `xorm:"comment('场次中间几率') CHAR(10)"`
	Gtype        int       `xorm:"default 1 comment('抽奖规则/次') INT(2)"`
	Productnum   int       `xorm:"comment('用户可参与的次数 0代表不限制 大于0就是次数') INT(2)"`
	Status       string    `xorm:"not null default '1' comment('状态 0:关闭 1:打开') CHAR(6)"`
	Sort         int       `xorm:"comment('排序号') INT(11)"`
	Addtime      time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
}
