package shop

import (
	"time"
)

type Log struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Logmsg     string    `xorm:"VARCHAR(2000)"`
	Createtime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
