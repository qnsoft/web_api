package shop

import (
	"time"
)

type DataDictionary struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	Code      string    `xorm:"not null comment('编码') VARCHAR(100)"`
	Name      string    `xorm:"not null comment('名称') VARCHAR(100)"`
	Value     string    `xorm:"not null comment('更新内容') TEXT"`
	Status    int       `xorm:"not null default 0 comment('是否生效 0:不是 1:是') TINYINT(4)"`
	AdminName string    `xorm:"not null comment('管理员名称') VARCHAR(100)"`
	Recycle   int       `xorm:"not null default 0 comment('回收站 0:正常 1:回收') TINYINT(4)"`
	AddDate   time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
