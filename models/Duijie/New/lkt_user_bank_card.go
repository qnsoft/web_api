package New

import (
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

/*
银行信息表
*/
type User_bank_card struct {
	Id             int       `xorm:"not null pk autoincr INT(11)" json:"id" orm:"column(id)"`
	StoreId        int       `xorm:"INT(11)" json:"store_id" orm:"column(store_id)"`
	UserId         string    `xorm:"VARCHAR(15)" json:"user_id" orm:"column(user_id)"`
	Cardholder     string    `xorm:"VARCHAR(30)" json:"cardholder" orm:"column(cardholder)"`
	IdCard         string    `xorm:"VARCHAR(30)" json:"id_card" orm:"column(id_card)"`
	BankName       string    `xorm:"VARCHAR(50)" json:"bank_name" orm:"column(bank_name)"`
	BankCode       string    `xorm:"VARCHAR(50)" json:"bank_code" orm:"column(bank_code)"`
	Branch         string    `xorm:"VARCHAR(250)" json:"branch" orm:"column(branch)"`
	BankCardNumber string    `xorm:"VARCHAR(30)" json:"bank_card_number" orm:"column(bank_card_number)"`
	Mobile         string    `xorm:"VARCHAR(20)" json:"mobile" orm:"column(mobile)"`
	AddDate        time.Time `xorm:"timestamp" json:"add_date" orm:"column(add_date)"`
	CardType       int       `xorm:"INT(11)" json:"card_type" orm:"column(card_type)"`
	IsDefault      int       `xorm:"TINYINT(4)" json:"is_default" orm:"column(is_default)"`
	Token          string    `xorm:"VARCHAR(30)" json:"token" orm:"column(token)"`
	Expiretime     string    `xorm:"VARCHAR(10)" json:"expiretime" orm:"column(expiretime)"`
	Cvv2           string    `xorm:"VARCHAR(10)" json:"cvv2" orm:"column(cvv2)"`
	Iscashonly     int       `xorm:"INT(11)" json:"iscashonly" orm:"column(iscashonly)"`
	Delflag        int       `xorm:"INT(11)" json:"delflag" orm:"column(delflag)"`
	Orderid        int       `xorm:"INT(11)" json:"orderid" orm:"column(orderid)"`
}
