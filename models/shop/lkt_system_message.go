package shop

import (
	"time"
)

type SystemMessage struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId     int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Senderid    string    `xorm:"not null comment('发送人ID') VARCHAR(30)"`
	Recipientid string    `xorm:"not null comment('接收人ID') VARCHAR(30)"`
	Title       string    `xorm:"comment('标题') TEXT"`
	Content     string    `xorm:"comment('内容') TEXT"`
	Time        time.Time `xorm:"comment('时间') DATETIME"`
	Type        int       `xorm:"not null default 1 comment('1未读  2 已读') INT(2)"`
}
