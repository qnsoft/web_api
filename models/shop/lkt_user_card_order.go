package shop

import (
	"time"
)

type UserCardOrder struct {
	Id                    int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	UserId                string    `xorm:"comment('用户id') VARCHAR(15)"`
	Ordernoa              string    `xorm:"comment('代扣订单号') VARCHAR(50)"`
	Ordernob              string    `xorm:"comment('代付订单号') VARCHAR(50)"`
	Amount                int       `xorm:"comment('订单金额') INT(11)"`
	Rateamounta           int       `xorm:"comment('代扣手续费') INT(11)"`
	Rateamountb           int       `xorm:"comment('代还手续费') INT(11)"`
	Bankcardno            string    `xorm:"comment('出金卡号') VARCHAR(50)"`
	Merchantbankaccountno string    `xorm:"comment('入金卡号') VARCHAR(50)"`
	ReturnA               string    `xorm:"comment('代扣返回信息') TEXT"`
	ReturnB               string    `xorm:"comment('代还返回信息') TEXT"`
	UserParentaId         string    `xorm:"comment('直推人id') VARCHAR(15)"`
	UserParentbId         string    `xorm:"comment('间推人id') VARCHAR(15)"`
	UserParentaProfit     int       `xorm:"comment('直推人返利') INT(11)"`
	UserParentbProfit     int       `xorm:"comment('间推人返利') INT(11)"`
	Isjob                 int       `xorm:"comment('是否计划任务') INT(11)"`
	AddTime               time.Time `xorm:"comment('订单时间') DATETIME"`
}
