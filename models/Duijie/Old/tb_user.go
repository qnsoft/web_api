package Old

import (
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

/*
用户表
*/
type User struct {
	Id        int       `xorm:"not null pk autoincr INT(11)" json:"id" orm:"column(id)"`
	Username  string    `xorm:"VARCHAR(50)" json:"username" orm:"column(username)"`
	Password  string    `xorm:"VARCHAR(500)" json:"password" orm:"column(password)"`
	Salt      string    `xorm:"VARCHAR(36)" json:"salt" orm:"column(salt)"`
	Status    int       `xorm:"INT(11)" json:"status" orm:"column(status)"`
	Addtime   time.Time `xorm:"DATETIME" json:"addtime" orm:"column(addtime)"`
	User_type int       `xorm:"INT(11)" json:"user_type" orm:"column(user_type)"`
}

/*
用户表扩展表
*/
type User_ext struct {
	Id         int     `xorm:"INT(11)" json:"id" orm:"column(id)"`
	Userid     int     `xorm:"INT(11)" json:"userid" orm:"column(userid)"`
	Parentid   int     `xorm:"INT(11)" json:"parentid" orm:"column(parentid)"`
	Layout     string  `xorm:"VARCHAR(8000)" json:"layout" orm:"column(layout)"`
	Regcode    string  `xorm:"VARCHAR(50)" json:"regcode" orm:"column(regcode)"`
	User_type  int     `xorm:"INT(11)" json:"user_type" orm:"column(user_type)"`
	Truename   string  `xorm:"VARCHAR(50)" json:"truename" orm:"column(truename)"`
	Telephone  string  `xorm:"VARCHAR(50)" json:"telephone" orm:"column(telephone)"`
	Qq         string  `xorm:"VARCHAR(50)" json:"qq" orm:"column(qq)"`
	Email      string  `xorm:"VARCHAR(50)" json:"email" orm:"column(email)"`
	User_money float64 `xorm:"DECIMAL(18)" json:"user_money" orm:"column(user_money)"`
	Smallimage string  `xorm:"VARCHAR(8000)" json:"smallimage" orm:"column(smallimage)"`
	Point      int     `xorm:"INT(11)" json:"point" orm:"column(point)"`
}

/*
用户VIP激活信息表
*/
type User_layout struct {
	Id      int  `xorm:"INT(11)" json:"id" orm:"column(id)"`
	Userid  int  `xorm:"INT(11)" json:"userid" orm:"column(userid)"`
	Layout  int  `xorm:"INT(11)" json:"layout" orm:"column(layout)"`
	Vipflag bool `xorm:"INT(11)" json:"vipflag" orm:"column(vipflag)"`
}
