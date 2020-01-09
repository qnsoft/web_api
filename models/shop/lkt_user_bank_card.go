package shop

import (
	"time"
)

/*
用户绑卡信息
*/
type UserBankCard struct {
	Id               int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId          int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId           string    `xorm:"not null comment('用户ID') VARCHAR(15)"`
	Cardholder       string    `xorm:"comment('持卡人') VARCHAR(30)"`
	IdCard           string    `xorm:"comment('身份证') VARCHAR(30)"`
	BankCode         string    `xorm:"comment('银行行别代码') VARCHAR(50)"`
	BankName         string    `xorm:"comment('银行名称') VARCHAR(50)"`
	Branch           string    `xorm:"comment('支行名称') VARCHAR(250)"`
	BankCardNumber   string    `xorm:"comment('银行卡号') VARCHAR(30)"`
	Mobile           string    `xorm:"comment('手机') VARCHAR(20)"`
	AddDate          time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	IsDefault        int       `xorm:"not null default 0 comment('是否默认地址 1默认') TINYINT(4)"`
	MchId            int       `xorm:"default 0 comment('店铺ID') INT(11)"`
	CardType         int       `xorm:"comment('卡类型0、储蓄卡1、信用卡') INT(11)"`
	Token            string    `xorm:"comment('与快付通签约识别号') VARCHAR(30)"`
	Expiretime       string    `xorm:"comment('信用卡有效期') VARCHAR(10)"`
	Cvv2             string    `xorm:"comment('cvv2') VARCHAR(10)"`
	Treatyid         string    `xorm:"comment('协议号') VARCHAR(50)"`
	TreatyidSmall    string    `xorm:"comment('小额协议号') VARCHAR(50)"`
	Treatytype       string    `xorm:"comment('协议类型：11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：余额+借记卡扣款 15： 余额+信用卡扣款') CHAR(10)"`
	Treatyenddate    string    `xorm:"comment('协议失效日期') VARCHAR(50)"`
	Iscashonly       int       `xorm:"comment('现金标志') INT(11)"`
	Delflag          int       `xorm:"comment('签约标志') INT(11)"`
	DelflagSmall     int       `xorm:"comment('小额签约标志') INT(11)"`
	Orderid          int       `xorm:"default 255 comment('排序号') INT(11)"`
	State            int       `xorm:"comment('解绑状态') INT(11)"`
	GyfStauts        string    `xorm:"comment('工易付签约状态：BIND:已绑定，其他未绑定') VARCHAR(50)"`
	GyfToken         string    `xorm:"comment('工易付签约token,若为空则未签约') VARCHAR(50)"`
	FfMerordernumber string    `xorm:"comment('丰付-商户签约订单号') VARCHAR(50)"`
	FfIssign         string    `xorm:"comment('丰付-商户签约标识') VARCHAR(50)"`
	FfKjzf           string    `xorm:"comment('丰付-快捷支付签约标识') VARCHAR(50)"`
}
