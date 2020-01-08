package shop

import (
	"time"
)

type ReplyComments struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(15)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Cid     string    `xorm:"not null comment('评论ID') VARCHAR(32)"`
	Uid     string    `xorm:"not null comment('用户id') VARCHAR(32)"`
	Content string    `xorm:"not null comment('评价内容') TEXT"`
	AddTime time.Time `xorm:"not null default '0000-00-00 00:00:00' TIMESTAMP"`
}
