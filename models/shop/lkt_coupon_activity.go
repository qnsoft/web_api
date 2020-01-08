package shop

import (
	"time"
)

type CouponActivity struct {
	Id           int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId      int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name         string `xorm:"not null comment('优惠券名称') VARCHAR(25)"`
	ActivityType int    `xorm:"default 0 comment('优惠券类型 1:满

减券 2:折扣券') INT(11)"`
	Money          string    `xorm:"default 0.00 comment('优惠券面值') DECIMAL(12,2)"`
	Discount       string    `xorm:"default 0.00 comment('折扣值') DECIMAL(12,2)"`
	ZMoney         string    `xorm:"default 0.00 comment('消费满多少') DECIMAL(12,2)"`
	Shopping       int       `xorm:"not null default 0 comment('满多少赠券') INT(11)"`
	FreeMailTask   int       `xorm:"not null default 0 comment('免邮任务 0：完善资料 1：绑定手机') TINYINT(4)"`
	Circulation    int       `xorm:"comment('发行数量') INT(11)"`
	Num            int       `xorm:"comment('剩余数量') INT(11)"`
	Receive        int       `xorm:"comment('领取限制') INT(11)"`
	UseNum         int       `xorm:"comment('使用限制') INT(11)"`
	StartTime      time.Time `xorm:"comment('开始时间') TIMESTAMP"`
	EndTime        time.Time `xorm:"comment('结束时间') TIMESTAMP"`
	Type           int       `xorm:"default 1 comment('优惠券使用范围 1：全部商品 2:指定商品 3：指定分类') TINYINT(4)"`
	ProductClassId string    `xorm:"comment('活动指定商品类型id') TEXT"`
	ProductId      string    `xorm:"comment('活动指定商品id') TEXT"`
	Content        string    `xorm:"comment('备注') TEXT"`
	Status         int       `xorm:"default 0 comment('状态 0：未启用 1：启用 2:禁用 3：已结束') TINYINT(4)"`
	AddTime        time.Time `xorm:"comment('修改时间') TIMESTAMP"`
	Qualifications string    `xorm:"comment('资格') TEXT"`
	Recycle        int       `xorm:"default 0 comment('回收站 0.不回收 1.回收') TINYINT(4)"`
}
