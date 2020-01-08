package shop

import (
	"time"
)

type Configure struct {
	Id           int       `xorm:"not null pk autoincr comment('配置id') INT(11)"`
	Name         string    `xorm:"comment('配置名称') VARCHAR(100)"`
	Color        string    `xorm:"not null comment('颜色') CHAR(15)"`
	Size         string    `xorm:"not null default '0' comment('尺码') VARCHAR(100)"`
	Costprice    string    `xorm:"not null comment('成本价') DECIMAL(12,2)"`
	Price        string    `xorm:"not null comment('出售价格') DECIMAL(12,2)"`
	Yprice       string    `xorm:"not null comment('原价格') DECIMAL(12,2)"`
	Url          string    `xorm:"comment('全路径') VARCHAR(500)"`
	Img          string    `xorm:"not null comment('图片') VARCHAR(255)"`
	Pid          int       `xorm:"not null comment('商品id') INT(11)"`
	TotalNum     int       `xorm:"not null default 0 comment('总库存') INT(11)"`
	Num          int       `xorm:"not null default 0 comment('剩余库存') INT(11)"`
	Unit         string    `xorm:"not null comment('单位') VARCHAR(30)"`
	BargainPrice string    `xorm:"not null comment('砍价开始价格') DECIMAL(12,2)"`
	Status       int       `xorm:"not null default 0 comment('状态 0:未开启砍价 1:开启砍价 2 上架 3 缺货 4下架') TINYINT(4)"`
	Attribute    string    `xorm:"comment('属性') TEXT"`
	Recycle      int       `xorm:"not null default 0 comment('回收站 0.不回收 1.回收') TINYINT(4)"`
	MinInventory int       `xorm:"not null default 10 comment('产品预警值') SMALLINT(6)"`
	Ctime        time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
}
