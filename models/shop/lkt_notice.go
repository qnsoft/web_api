package shop

import (
	"time"
)

type Notice struct {
	Id              int       `xorm:"not null pk autoincr comment('配置id') INT(11)"`
	StoreId         int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	PaySuccess      string    `xorm:"comment('购买成功通知') VARCHAR(100)"`
	OrderDelivery   string    `xorm:"comment('订单发货提醒') VARCHAR(100)"`
	OrderSuccess    string    `xorm:"comment('订单支付成功通知') VARCHAR(100)"`
	GroupPaySuccess string    `xorm:"comment('开团成功提醒') VARCHAR(100)"`
	GroupPending    string    `xorm:"comment('拼团待成团提醒') VARCHAR(100)"`
	GroupSuccess    string    `xorm:"comment('拼团成功通知') VARCHAR(100)"`
	GroupFail       string    `xorm:"comment('拼团失败通知') VARCHAR(100)"`
	LotteryRes      string    `xorm:"comment('抽奖结果通知') VARCHAR(100)"`
	RefundSuccess   string    `xorm:"comment('退款成功通知') VARCHAR(100)"`
	RefundRes       string    `xorm:"comment('退款通知') VARCHAR(100)"`
	Receive         string    `xorm:"comment('领取通知') VARCHAR(100)"`
	UpdateTime      time.Time `xorm:"comment('更新时间') DATETIME"`
}
