package shop

import (
	"time"
)

type SessionId struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	SessionId string    `xorm:"not null comment('session_id') VARCHAR(100)"`
	Content   string    `xorm:"not null comment('内容') TEXT"`
	AddDate   time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
