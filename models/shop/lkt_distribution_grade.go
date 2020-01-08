package shop

import (
	"time"
)

type DistributionGrade struct {
	Id                  int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId             int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Sets                string    `xorm:"TEXT"`
	Datetime            time.Time `xorm:"not null comment('创建时间') DATETIME"`
	IsAgent             int       `xorm:"default 0 comment('是否为代理商') INT(2)"`
	IsOrdinary          int       `xorm:"default 0 comment('是否是普通分销商0 不是 1是') INT(2)"`
	TransferBalance     int       `xorm:"default 0 comment('推荐后消费金转余额比例') INT(5)"`
	Sort                int       `xorm:"default 1 comment('排序号') INT(4)"`
	ConsumerMoney       string    `xorm:"default 0.00 comment('消费金') DECIMAL(12,2)"`
	Integral            int       `xorm:"default 0 comment('积分') INT(12)"`
	MemberProportion    string    `xorm:"default '0' comment('会员专区佣金比例') VARCHAR(255)"`
	MemberConsumerMoney string    `xorm:"default '' comment('会员专区消费金') VARCHAR(255)"`
}
