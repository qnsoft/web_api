package shop

import (
	"time"
)

type MessageList struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Type       int       `xorm:"not null default 0 comment('类型 0:验证码 1:短信通知') INT(2)"`
	Type1      int       `xorm:"not null default 0 comment('类别 0:通用 1:申请通过 2:申请拒绝 3:订单提现 4:发货提现 5:收货提现') INT(2)"`
	TemplateId int       `xorm:"not null default 0 comment('模板id') INT(11)"`
	Content    string    `xorm:"not null comment('code') VARCHAR(50)"`
	Status     int       `xorm:"not null default 0 comment('类型 0:商城 1:店铺') INT(2)"`
	AdminId    string    `xorm:"not null comment('客户id/店主id') VARCHAR(50)"`
	AddTime    time.Time `xorm:"not null comment('创建时间') DATETIME"`
}
