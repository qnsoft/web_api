package shop

import (
	"time"
)

type MingxiLog struct {
	Id      int       `xorm:"not null pk autoincr INT(11)"`
	Type    int       `xorm:"not null default 1 comment('1.积分明细 2.臻珠明细 3.余额明细') TINYINT(5)"`
	Value   string    `xorm:"not null comment('明细金额') DECIMAL(10)"`
	AddTime time.Time `xorm:"not null comment('添加时间') DATETIME"`
	FormId  int       `xorm:"comment('来自于某人的转让用户id') INT(11)"`
	FormImg string    `xorm:"comment('来自于某人的转让图片') VARCHAR(255)"`
	Status  int       `xorm:"comment('1转出 2 提现 3 转入') TINYINT(11)"`
}
