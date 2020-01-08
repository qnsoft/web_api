package shop

import (
	"time"
)

type Extension struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(15)"`
	StoreId   string    `xorm:"not null default '0' comment('商城id') CHAR(11)"`
	Name      string    `xorm:"not null comment('名称') VARCHAR(50)"`
	Image     string    `xorm:"not null comment('图片路径') VARCHAR(50)"`
	Sort      string    `xorm:"not null default '100' comment('排序号') CHAR(5)"`
	X         string    `xorm:"not null default '0' comment('x坐标') CHAR(5)"`
	Y         string    `xorm:"not null default '0' comment('y坐标') CHAR(5)"`
	AddDate   time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Kuan      int       `xorm:"not null default 150 comment('二维码宽度') INT(5)"`
	Type      int       `xorm:"not null comment('海报类型 1.文章海报 2.红包海报 3.商品海报 4.分销海报 5.邀请海报 6.竞拍海报') INT(2)"`
	Keyword   string    `xorm:"not null comment('关键词') VARCHAR(100)"`
	Url       string    `xorm:"comment('链接地址') VARCHAR(255)"`
	Isdefault int       `xorm:"default 0 comment('是否默认') INT(2)"`
	Bg        string    `xorm:"comment('背景图片') VARCHAR(200)"`
	Color     string    `xorm:"comment('颜色') VARCHAR(10)"`
	Waittext  string    `xorm:"comment('等待语') VARCHAR(200)"`
	Data      string    `xorm:"comment('排序的数据') TEXT"`
	StoreType int       `xorm:"default 1 comment('来源 1.小程序 2.app') TINYINT(4)"`
}
