package shop

import (
	"time"
)

type ProductImg struct {
	Id         int       `xorm:"not null pk autoincr comment('图片id') INT(11)"`
	ProductUrl string    `xorm:"not null comment('产品图片') VARCHAR(100)"`
	ProductId  int       `xorm:"not null comment('所属产品id') INT(11)"`
	SellerId   string    `xorm:"not null default '' comment('用户id') CHAR(15)"`
	AddDate    time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Sort       int       `xorm:"not null default 0 comment('是否主图 默认0 不是主图，1为主图') TINYINT(1)"`
}
