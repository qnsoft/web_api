package New

import (
	_ "github.com/go-sql-driver/mysql"
)

/*
银行信息表(快付通)
*/
type Channel_bank_kft struct {
	Id         int     `xorm:"not null pk autoincr INT(11)" json:"id" orm:"column(id)"`
	Bankcode   string  `xorm:"VARCHAR(255)" json:"bankcode" orm:"column(bankcode)"`
	Bankname   string  `xorm:"VARCHAR(255)" json:"bankname" orm:"column(bankname)"`
	D0freerate float64 `xorm:"DECIMAL(10,0)	" json:"d0freerate" orm:"column(d0freerate)"`
	D0fixrate  float64 `xorm:"INT(11)" json:"d0fixrate" orm:"column(d0fixrate)"`
	T1freerate float64 `xorm:"DECIMAL(10,0)" json:"t1freerate" orm:"column(t1freerate)"`
	T1fixrate  float64 `xorm:"INT(11)" json:"t1fixrate" orm:"column(t1fixrate)"`
	Channelid  float64 `xorm:"INT(11)" json:"channelid" orm:"column(channelid)"`
}
