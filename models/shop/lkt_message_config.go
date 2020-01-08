package shop

import (
	"time"
)

type MessageConfig struct {
	Id              int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId         int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Accesskeyid     string    `xorm:"not null comment('accessKeyId') VARCHAR(50)"`
	Accesskeysecret string    `xorm:"not null comment('accessKeySecret') VARCHAR(50)"`
	Signname        string    `xorm:"not null comment('短信签名') VARCHAR(50)"`
	AddTime         time.Time `xorm:"not null comment('创建时间') DATETIME"`
}
