package shop

import (
	"time"
)

type ImgGroup struct {
	Id        int       `xorm:"not null pk autoincr comment('分组id') INT(11)"`
	StoreId   int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name      string    `xorm:"not null comment('分组名称') VARCHAR(10)"`
	Sort      int       `xorm:"default 100 comment('排序') INT(11)"`
	AddDate   time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	IsDefault int       `xorm:"not null default 0 comment('0不是默认 1默认') TINYINT(2)"`
}
