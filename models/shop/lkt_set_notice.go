package shop

import (
	"time"
)

type SetNotice struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name    string    `xorm:"not null comment('公告名称') VARCHAR(50)"`
	Detail  string    `xorm:"TEXT"`
	ImgUrl  string    `xorm:"not null default '' CHAR(30)"`
	Time    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	User    string    `xorm:"not null comment('发布者') VARCHAR(55)"`
}
