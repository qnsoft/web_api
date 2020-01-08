package shop

import (
	"time"
)

type DistributionRecord struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId  string    `xorm:"comment('购买员') CHAR(15)"`
	FromId  string    `xorm:"comment('分销员') CHAR(15)"`
	Money   string    `xorm:"not null default 0.00 comment('佣金金额') DECIMAL(12,2)"`
	Sno     string    `xorm:"comment('订单号') VARCHAR(255)"`
	Level   int       `xorm:"comment('层级') INT(5)"`
	Event   string    `xorm:"comment('事件') VARCHAR(255)"`
	Type    int       `xorm:"not null default 0 comment('类型 1:转入(收入) 2:提现 3:管理佣金 4:使用消费金 5:转入消费金 6:系统扣除(消费金) 7:消费金解封 8:充值积分 9:使用积分 10:使用消费金抽奖 11:中奖获得消费金 12:购买赠送 13:系统充值(消费金) 14:转出消费金') TINYINT(4)"`
	AddDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Status  int       `xorm:"not null default 0 comment('是否发放') INT(10)"`
}
