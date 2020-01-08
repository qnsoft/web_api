package New

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
银行信息表(快付通)
*/
type User_auth struct {
	Id            int       `xorm:"not null pk autoincr INT(11)" json:"id" orm:"column(id)"`
	StoreId       int       `xorm:"INT(11)" json:"store_id" orm:"column(store_id)"`
	UserId        string    `xorm:"CHAR(15)" json:"user_id" orm:"column(user_id)"`
	Truename      string    `xorm:"VARCHAR(50)" json:"truename" orm:"column(truename)"`
	Idcard        string    `xorm:"VARCHAR(50)" json:"idcard" orm:"column(idcard)"`
	Personimg     string    `xorm:"VARCHAR(255)" json:"personimg" orm:"column(personimg)"`
	Personimgback string    `xorm:"VARCHAR(255)" json:"personimgback" orm:"column(personimgback)"`
	Authflag      bool      `xorm:"BIT(1)" json:"authflag" orm:"column(authflag)"`
	Value         float64   `xorm:"VARCHAR(255)" json:"value" orm:"column(value)"`
	Ysbflag       bool      `xorm:"BIT(1)" json:"ysbflag" orm:"column(ysbflag)"`
	Ordertop      int       `xorm:"INT(11)" json:"ordertop" orm:"column(ordertop)"`
	Addtime       time.Time `xorm:"DATETIME" json:"addtime" orm:"column(addtime)"`
}
