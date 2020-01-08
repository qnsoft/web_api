package shop

import (
	"time"
)

type Edition struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Edition    string    `xorm:"not null comment('版本号') VARCHAR(100)"`
	EditionUrl string    `xorm:"not null comment('路径') VARCHAR(100)"`
	Type       int       `xorm:"not null default 0 comment('是否是热更新 0:是 1:不是') TINYINT(4)"`
	Content    string    `xorm:"not null comment('更新内容') TEXT"`
	AddDate    time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
