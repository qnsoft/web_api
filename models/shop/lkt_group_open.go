package shop

import (
	"time"
)

type GroupOpen struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Uid       string    `xorm:"not null default '' comment('微信id') CHAR(40)"`
	PtgoodsId int       `xorm:"comment('拼团商品id') INT(11)"`
	Ptcode    string    `xorm:"comment('拼团编号') VARCHAR(50)"`
	Groupman  int       `xorm:"not null comment('几人团') SMALLINT(6)"`
	Ptnumber  int       `xorm:"comment('拼团人数') INT(11)"`
	Addtime   time.Time `xorm:"comment('创建日期') DATETIME"`
	Endtime   time.Time `xorm:"comment('结束时间') DATETIME"`
	Ptstatus  int       `xorm:"default 0 comment('0:未付款 1:拼团中，2:拼团成功, 3：拼团失败, ') TINYINT(1)"`
}
