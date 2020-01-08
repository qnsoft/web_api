package shop

import (
	"time"
)

type TaobaokeRecord struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId     int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId      string    `xorm:"comment('推荐人') CHAR(15)"`
	Buyuser     string    `xorm:"comment('购买人') CHAR(15)"`
	Sno         string    `xorm:"comment('订单号') VARCHAR(255)"`
	GoodsTitle  string    `xorm:"comment('商品名称') VARCHAR(255)"`
	GoodsImg    string    `xorm:"comment('商品缩略图') TEXT"`
	GoodsPrice  string    `xorm:"not null default 0.00 comment('商品单价') DECIMAL(12,2)"`
	Num         int       `xorm:"not null default 1 comment('购买商品数量') INT(11)"`
	Amount      string    `xorm:"not null default 0.00 comment('实际支付金额') DECIMAL(12,2)"`
	CommRate    string    `xorm:"not null default 0.00 comment('佣金比例') DECIMAL(10,2)"`
	CommAmount  string    `xorm:"not null default 0.00 comment('佣金金额') DECIMAL(12,2)"`
	Status      int       `xorm:"default 0 comment('0未结算 1已结算') INT(5)"`
	CreateTime  time.Time `xorm:"comment('创建订单时间') TIMESTAMP"`
	EarningTime time.Time `xorm:"comment('订单结算时间') TIMESTAMP"`
	ClickTime   time.Time `xorm:"comment('跟踪时间') TIMESTAMP"`
	OStatus     int       `xorm:"default 3 comment('淘客订单状态，3：订单结算，12：订单付款， 13：订单失效，14：订单成功') INT(5)"`
}
