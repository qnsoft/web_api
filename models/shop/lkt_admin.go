package shop

import (
	"time"
)

type Admin struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	Sid        int       `xorm:"not null default 0 comment('上级id') INT(11)"`
	Name       string    `xorm:"not null default '' comment('账号') VARCHAR(30)"`
	Password   string    `xorm:"not null default '' comment('密码') CHAR(32)"`
	AdminType  int       `xorm:"comment('管理类型') INT(11)"`
	Permission string    `xorm:"comment('许可') TEXT"`
	Role       string    `xorm:"comment('角色 ') VARCHAR(255)"`
	Type       int       `xorm:"default 0 comment('类型 0:系统管理员 1：客户 2:商城管理员 3:店主') INT(11)"`
	StoreId    int       `xorm:"default 0 comment('商城id') INT(11)"`
	ShopId     int       `xorm:"not null comment('店铺ID') INT(11)"`
	Status     int       `xorm:"not null default 2 comment('状态 1:禁用 2：启用') INT(11)"`
	Nickname   string    `xorm:"comment('昵称') VARCHAR(20)"`
	Birthday   string    `xorm:"comment('生日') VARCHAR(32)"`
	Sex        int       `xorm:"default 1 comment('性别（1 男  2 女）') INT(2)"`
	Tel        string    `xorm:"comment('手机号码') VARCHAR(22)"`
	Token      string    `xorm:"comment('令牌') VARCHAR(32)"`
	Ip         string    `xorm:"comment('登入ip地址') VARCHAR(10)"`
	Recycle    int       `xorm:"comment('回收站 0:不回收 1:回收 ') TINYINT(2)"`
	AddDate    time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	LoginNum   int       `xorm:"not null default 0 comment('登录次数') INT(11)"`
}
