package shop

import (
	"time"
)

type SystemTell struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Title     string    `xorm:"not null default '' comment('公告标题') VARCHAR(300)"`
	Type      int       `xorm:"comment('公告类型: 1--系统维护  2--版本更新') TINYINT(4)"`
	Startdate string    `xorm:"not null comment('开始时间') CHAR(20)"`
	Enddate   string    `xorm:"not null comment('结束时间') CHAR(20)"`
	Content   string    `xorm:"comment('内容') TEXT"`
	AddTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
}
