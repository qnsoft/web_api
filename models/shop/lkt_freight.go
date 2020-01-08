package shop

import (
	"time"
)

type Freight struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name      string    `xorm:"not null default '' comment('名称') VARCHAR(60)"`
	Type      int       `xorm:"not null default 0 comment('类型 0:件 1:重量') TINYINT(4)"`
	Freight   string    `xorm:"comment('运费') TEXT"`
	IsDefault int       `xorm:"not null default 0 comment('类型 0:不默认 1:默认') TINYINT(4)"`
	AddTime   time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
