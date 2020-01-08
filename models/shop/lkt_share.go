package shop

import (
	"time"
)

type Share struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId    string    `xorm:"comment('用户id') CHAR(15)"`
	WxId      string    `xorm:"comment('微信id') VARCHAR(50)"`
	WxName    string    `xorm:"comment('微信昵称') VARCHAR(150)"`
	Sex       int       `xorm:"comment('性别 0:未知 1:男 2:女') INT(11)"`
	Type      int       `xorm:"not null default 0 comment('类别 0：新闻 1：文章') TINYINT(4)"`
	ArticleId int       `xorm:"not null default 0 comment('新闻id') INT(11)"`
	ShareAdd  time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('分享时间') TIMESTAMP"`
	Coupon    string    `xorm:"not null default 0.00 comment('礼券') DECIMAL(12,2)"`
}
