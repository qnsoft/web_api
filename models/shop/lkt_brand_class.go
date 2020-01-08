package shop

import (
	"time"
)

type BrandClass struct {
	BrandId    int       `xorm:"not null pk autoincr comment('品牌id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	BrandPic   string    `xorm:"comment('品牌logo') VARCHAR(255)"`
	BrandName  string    `xorm:"comment('品牌名称') VARCHAR(255)"`
	BrandYName string    `xorm:"comment('英文名') VARCHAR(100)"`
	Producer   string    `xorm:"comment('产地') VARCHAR(255)"`
	Remarks    string    `xorm:"comment('备注') VARCHAR(255)"`
	Status     int       `xorm:"not null default 0 comment('状态 0:启用 1:禁用') TINYINT(4)"`
	BrandTime  time.Time `xorm:"comment('时间') TIMESTAMP"`
	Sort       int       `xorm:"default 100 INT(10)"`
	Recycle    int       `xorm:"not null default 0 comment('回收站 0.不回收 1.回收') TINYINT(4)"`
}
