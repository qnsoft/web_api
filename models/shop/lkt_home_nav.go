package shop

import (
	"time"
)

type HomeNav struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"comment('商城id') INT(11)"`
	StoreType string    `xorm:"default '' comment('软件类型') VARCHAR(255)"`
	Name      string    `xorm:"not null default '' comment('图标名称') VARCHAR(255)"`
	Uniquely  string    `xorm:"not null pk comment('标识') VARCHAR(255)"`
	Url       string    `xorm:"not null default '' comment('页面路径') VARCHAR(500)"`
	OpenType  string    `xorm:"not null default '' comment('打开方式') VARCHAR(255)"`
	PicUrl    string    `xorm:"not null comment('图标url') LONGTEXT"`
	AddTime   time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Sort      int       `xorm:"not null default 100 comment('排序，升序') INT(11)"`
	IsDelete  int       `xorm:"not null default 0 SMALLINT(1)"`
	IsHide    int       `xorm:"not null default 0 comment('是否隐藏 0 显示 1隐藏 ') SMALLINT(1)"`
}
