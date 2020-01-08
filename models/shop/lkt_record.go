package shop

import (
	"time"
)

type Record struct {
	Id       int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId  int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId   string    `xorm:"comment('用户id') CHAR(15)"`
	Money    string    `xorm:"not null default 0.00 comment('操作金额') DECIMAL(12,2)"`
	Oldmoney string    `xorm:"not null default 0.00 comment('原有金额') DECIMAL(12,2)"`
	AddDate  time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Event    string    `xorm:"comment('事件') TEXT"`
	Type     int       `xorm:"not null default 0 comment('类型 0:登录/退出 1:充值 2:申请提现 3:分享 4:余额消费 5:退款 6:红包提现 7:佣金 8:管理佣金 9:待定 10:消费金 11:系统扣款   12:给好友转余额 13:转入余额 14:系统充值 15:系统充积分 16:系统充消费金 17:系统扣积分 18:系统扣消费金 19:消费金解封 20:抽奖中奖 21:  提现成功 22:提现失败 23.取消订单  24分享获取红包 26 交竞拍押金 27 退竞拍押金') TINYINT(4)"`
	IsMch    int       `xorm:"default 0 comment('是否是店铺提现 0.不是店铺提现 1.是店铺提现') INT(2)"`
}
