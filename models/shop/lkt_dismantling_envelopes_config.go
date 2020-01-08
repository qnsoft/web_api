package shop

import (
	"time"
)

type DismantlingEnvelopesConfig struct {
	Id               int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId          int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	PlugInsId        int       `xorm:"not null default 0 comment('插件id') INT(11)"`
	PeopleNum        int       `xorm:"comment('拆红包人数') INT(11)"`
	CollectPeopleNum int       `xorm:"comment('确认收货拆红包人数') INT(11)"`
	UpperLimitMoney  string    `xorm:"not null default 0.00 comment('上限金额') DECIMAL(12,2)"`
	Unit             string    `xorm:"not null default '元' comment('小程序钱包单位') VARCHAR(50)"`
	InvalidTime      int       `xorm:"comment('逾期失效时间') INT(11)"`
	ActivityOverdue  int       `xorm:"not null default 0 comment('活动过期删除时间') INT(11)"`
	AddTime          time.Time `xorm:"comment('修改时间') TIMESTAMP"`
}
