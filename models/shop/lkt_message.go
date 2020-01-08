package shop

import (
	"time"
)

type Message struct {
	Id           int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId      int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name         string    `xorm:"not null comment('短信模板名称') VARCHAR(50)"`
	Type         int       `xorm:"not null default 0 comment('类型 0:验证码 1:自定义') INT(2)"`
	Type1        int       `xorm:"not null default 0 comment('类别 0:通用 1:申请通过 2:申请拒绝 3:订单提现 4:发货提现 5:收货提现') INT(2)"`
	Templatecode string    `xorm:"not null comment('短信模板Code') VARCHAR(50)"`
	Content      string    `xorm:"not null default '' comment('内容') VARCHAR(50)"`
	AddTime      time.Time `xorm:"not null comment('创建时间') DATETIME"`
}
