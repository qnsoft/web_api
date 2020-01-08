package shop

import (
	"time"
)

type ProductClass struct {
	Cid     int       `xorm:"not null pk autoincr comment('分类id') INT(11)"`
	Sid     int       `xorm:"not null comment('上级id') INT(11)"`
	Pname   string    `xorm:"not null comment('分类名称') CHAR(15)"`
	Img     string    `xorm:"default '' comment('分类图片') VARCHAR(200)"`
	Bg      string    `xorm:"comment('小图标') VARCHAR(255)"`
	Level   int       `xorm:"not null comment('级别') INT(11)"`
	Sort    int       `xorm:"default 100 comment('排序') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	AddDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Recycle int       `xorm:"default 0 comment('回收站 0.不回收 1.回收') TINYINT(4)"`
	IsShow  int       `xorm:"not null default 1 comment('是否是线上  1 线上  0附近 2课堂') INT(11)"`
}
