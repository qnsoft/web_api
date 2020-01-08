package shop

import (
	"time"
)

type Order struct {
	Id                 int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId            int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId             string    `xorm:"not null comment('用户id') CHAR(15)"`
	Name               string    `xorm:"default '' comment('姓名') VARCHAR(50)"`
	Mobile             string    `xorm:"comment('联系电话') VARCHAR(20)"`
	Num                int       `xorm:"comment('数量') INT(11)"`
	ZPrice             string    `xorm:"default 0.00 comment('总价') DECIMAL(12,2)"`
	Sno                string    `xorm:"comment('订单号') VARCHAR(255)"`
	Sheng              int       `xorm:"comment('省') INT(11)"`
	Shi                int       `xorm:"comment('市') INT(11)"`
	Xian               int       `xorm:"comment('县') INT(11)"`
	Address            string    `xorm:"comment('详细地址') TEXT"`
	Remark             string    `xorm:"comment('用户备注') VARCHAR(32)"`
	Pay                string    `xorm:"comment('支付方式') CHAR(15)"`
	AddTime            time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	PayTime            time.Time `xorm:"comment('支付时间') TIMESTAMP"`
	ArriveTime         time.Time `xorm:"comment('到货时间') TIMESTAMP"`
	Status             int       `xorm:"default 0 comment('状态 0:未付款 1:未发货 2:待收货 3:待评论 4:退货 5:已完成 6:订单失效 7:订单  关闭 8:删除订单 9:拼团中 10:拼团失败-未退款 11:拼团失败-已退款 ') TINYINT(4)"`
	CouponId           int       `xorm:"default 0 comment('优惠券id') INT(11)"`
	ConsumerMoney      int       `xorm:"default 0 comment('消费金') INT(11)"`
	CouponActivityName string    `xorm:"default '0' comment('满减活动名称') VARCHAR(50)"`
	Drawid             int       `xorm:"default 0 comment('活动ID') INT(11)"`
	Otype              string    `xorm:"default '' comment('订单类型') CHAR(30)"`
	Ptcode             string    `xorm:"default '' comment('拼团编号') CHAR(15)"`
	Pid                string    `xorm:"comment('拼团订单类型:kaituan开团 cantuan参团') CHAR(10)"`
	Ptstatus           int       `xorm:"comment('拼团状态:0,未付款 1,拼团中 2,拼团成功 3,拼团失败') SMALLINT(6)"`
	Groupman           string    `xorm:"comment('拼团人数') CHAR(5)"`
	Refundsno          string    `xorm:"default '' comment('拼团退款单号') CHAR(30)"`
	TradeNo            string    `xorm:"default '' comment('微信支付单号') CHAR(20)"`
	IsAnonymous        int       `xorm:"default 0 comment('是否匿名(0,否  1.是') INT(1)"`
	LogisticsService   int       `xorm:"comment('物流服务') INT(1)"`
	OverallEvaluation  int       `xorm:"comment('总体评价') INT(1)"`
	SpzPrice           string    `xorm:"default 0.00 comment('商品总价') DECIMAL(10,2)"`
	ReducePrice        string    `xorm:"default 0.00 comment('查询出的满减金额') DECIMAL(10,2)"`
	CouponPrice        string    `xorm:"default 0.00 comment('查询出的优惠券金额') DECIMAL(10,2)"`
	RedPacket          string    `xorm:"default '0' comment('红包金额') VARCHAR(255)"`
	Allow              string    `xorm:"default 0.00 comment('积分') DECIMAL(10,2)"`
	Parameter          string    `xorm:"comment('参数') TEXT"`
	Source             int       `xorm:"not null default 1 comment('来源 1.小程序 2.app') TINYINT(4)"`
	DeliveryStatus     int       `xorm:"default 0 comment('提醒发货') INT(1)"`
	Readd              int       `xorm:"not null default 0 comment('是否已读（0，未读  1 已读）') INT(2)"`
	OffsetBalance      string    `xorm:"default 0.00 comment('抵扣余额') DECIMAL(10,2)"`
	MchId              string    `xorm:"default '' comment('店铺ID') VARCHAR(11)"`
	Zhekou             string    `xorm:"not null default 0.00 comment('分销折扣') DECIMAL(10,2)"`
}
