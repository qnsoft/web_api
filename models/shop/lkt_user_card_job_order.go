package shop

import (
	"time"
)

type UserCardJobOrder struct {
	Id                    int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	ParentOrderId         int       `xorm:"comment('父订单id') INT(11)"`
	UserId                string    `xorm:"comment('用户id') VARCHAR(15)"`
	Ordernoa              string    `xorm:"comment('代扣订单号') VARCHAR(50)"`
	Ordernob              string    `xorm:"comment('代付订单号') VARCHAR(50)"`
	Amount                int       `xorm:"comment('订单金额') INT(11)"`
	Rateamounta           int       `xorm:"comment('代扣手续费') INT(11)"`
	Rateamountb           int       `xorm:"comment('代付还手续') INT(11)"`
	Bankcardno            string    `xorm:"comment('出金卡号') VARCHAR(50)"`
	Merchantbankaccountno string    `xorm:"comment('入金卡号') VARCHAR(50)"`
	ReturnA               string    `xorm:"comment('代扣返回信息') TEXT"`
	ReturnB               string    `xorm:"comment('代还返回信息') TEXT"`
	UserParentaId         string    `xorm:"comment('直推人id') VARCHAR(15)"`
	UserParentbId         string    `xorm:"comment('间推人id') VARCHAR(15)"`
	UserParentaProfit     int       `xorm:"comment('直推人返利') INT(11)"`
	UserParentbProfit     int       `xorm:"comment('间推人返利') INT(11)"`
	Treatyid              string    `xorm:"comment('出金卡协议号') VARCHAR(50)"`
	AddTime               time.Time `xorm:"comment('订单时间') DATETIME"`
	TaskName              string    `xorm:"comment('任务名称') VARCHAR(50)"`
	CronSpec              string    `xorm:"comment('时间表达式') VARCHAR(100)"`
	ImplementTime         string    `xorm:"comment('待执行时间') VARCHAR(50)"`
	Status                int       `xorm:"default 1 comment('任务状态：0停用 1启用') INT(4)"`
	IsFinish              int       `xorm:"default 0 comment('订单是否完成：0未完成 1完成') INT(4)"`
}
