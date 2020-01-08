package shop

import (
	"time"
)

type Customer struct {
	Id             int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	AdminId        int       `xorm:"not null default 0 comment('管理员id') INT(11)"`
	CustomerNumber string    `xorm:"comment('客户编号') VARCHAR(250)"`
	Name           string    `xorm:"not null default '' comment('姓名') VARCHAR(20)"`
	Mobile         string    `xorm:"comment('手机') VARCHAR(20)"`
	Price          string    `xorm:"not null default 0.00 comment('价格') DECIMAL(12,2)"`
	Company        string    `xorm:"not null default '' comment('公司名称') VARCHAR(50)"`
	Function       string    `xorm:"not null comment('功能') TEXT"`
	AddDate        time.Time `xorm:"comment('购买时间') TIMESTAMP"`
	EndDate        time.Time `xorm:"comment('到期时间') TIMESTAMP"`
	Status         int       `xorm:"not null default 0 comment('类型 0:启用 1:到期') TINYINT(4)"`
	Email          string    `xorm:"not null comment('邮箱') VARCHAR(255)"`
	Recycle        int       `xorm:"not null default 0 comment('回收站 0:不回收 1:回收 ') TINYINT(4)"`
}
