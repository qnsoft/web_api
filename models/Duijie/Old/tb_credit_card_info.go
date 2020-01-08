package Old

import (
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

/*
银行信息表(快付通)
*/
type Credit_card_info struct {
	Id         int       `xorm:"not null pk autoincr INT(11)" json:"id" orm:"column(id)"`
	Memberid   string    `xorm:"int(11)" json:"memberid" orm:"column(memberid)"`
	Name       string    `xorm:"VARCHAR(50)" json:"name" orm:"column(name)"`
	Certno     string    `xorm:"VARCHAR(50)" json:"certno" orm:"column(certno)"`
	Bankname   string    `xorm:"VARCHAR(50)" json:"bankname" orm:"column(bankname)"`
	Cardno     string    `xorm:"VARCHAR(50)" json:"cardno" orm:"column(cardno)"`
	Phone      string    `xorm:"VARCHAR(50)" json:"phone" orm:"column(phone)"`
	Addtime    time.Time `xorm:"DATETIME" json:"addtime" orm:"column(addtime)"`
	Cardtype   string    `xorm:"VARCHAR(50)" json:"cardtype" orm:"column(cardtype)"`
	Token      string    `xorm:"VARCHAR(50)" json:"token" orm:"column(token)"`
	Expiretime string    `xorm:"VARCHAR(50)" json:"expiretime" orm:"column(expiretime)"`
	Cvv2       string    `xorm:"VARCHAR(50)" json:"cvv2" orm:"column(cvv2)"`
	Iscashonly int       `xorm:"int(11)" json:"iscashonly" orm:"column(iscashonly)"`
}
