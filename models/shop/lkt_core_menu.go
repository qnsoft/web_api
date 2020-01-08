package shop

import (
	"time"
)

type CoreMenu struct {
	Id       int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	SId      int       `xorm:"comment('上级id') INT(11)"`
	Title    string    `xorm:"not null default '' comment('菜单名称') VARCHAR(60)"`
	Name     string    `xorm:"not null default '' comment('菜单标识') VARCHAR(60)"`
	Image    string    `xorm:"comment('图标') VARCHAR(255)"`
	Image1   string    `xorm:"comment('点击后图标') VARCHAR(255)"`
	Module   string    `xorm:"not null default '' comment('菜单模块标识') VARCHAR(60)"`
	Action   string    `xorm:"not null default '' comment('菜单文件标识') VARCHAR(60)"`
	Url      string    `xorm:"not null default '' comment('路径') VARCHAR(60)"`
	Sort     int       `xorm:"comment('排序') INT(11)"`
	Level    int       `xorm:"comment('第几级') INT(11)"`
	IsCore   int       `xorm:"default 0 comment('是否是核心') TINYINT(4)"`
	IsPlugIn int       `xorm:"default 0 comment('是否是插件') TINYINT(4)"`
	Type     int       `xorm:"default 0 comment('类型 0.后台管理 1.小程序 2.app 3.微信公众号 4.PC 5.生活号 6.报表') TINYINT(4)"`
	AddTime  time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Recycle  int       `xorm:"not null default 0 comment('回收站 0.不回收 1.回收') TINYINT(4)"`
}
