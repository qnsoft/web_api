package shop

import (
	"time"
)

type TwelveDrawProduct struct {
	Id      int       `xorm:"not null pk autoincr comment('产品id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name    string    `xorm:"not null default '' comment('产品名字') VARCHAR(100)"`
	Imgurl  string    `xorm:"not null default '' comment('产品图片') VARCHAR(200)"`
	Sort    string    `xorm:"comment('排序') CHAR(11)"`
	AddDate time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Prate   string    `xorm:"not null comment('产品中奖几率') CHAR(10)"`
	Value   string    `xorm:"default '0' comment('产品值') VARCHAR(11)"`
	Num     int       `xorm:"not null default 0 comment('数量') INT(11)"`
	Fid     int       `xorm:"comment('活动场次id') INT(10)"`
	Type    int       `xorm:"default 0 comment('类型 1:谢谢 2:赠品 3:积分 4:消费金') TINYINT(3)"`
}
