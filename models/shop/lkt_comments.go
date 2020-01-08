package shop

import (
	"time"
)

type Comments struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId     int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Oid         string    `xorm:"comment('订单ID') VARCHAR(32)"`
	Uid         string    `xorm:"comment('用户id') VARCHAR(32)"`
	Pid         string    `xorm:"comment('商品id') VARCHAR(32)"`
	AttributeId int       `xorm:"comment('属性id') INT(11)"`
	Size        string    `xorm:"default '默认' comment('商品规格') CHAR(50)"`
	Content     string    `xorm:"comment('评价内容') TEXT"`
	Commenttype string    `xorm:"comment('评价类型') CHAR(11)"`
	Anonymous   string    `xorm:"default '0' comment('匿名 ') VARCHAR(32)"`
	AddTime     time.Time `xorm:"default '0000-00-00 00:00:00' TIMESTAMP"`
	Review      string    `xorm:"comment('追评') TEXT"`
	ReviewTime  time.Time `xorm:"default '0000-00-00 00:00:00' comment('追评时间') TIMESTAMP"`
	Type        string    `xorm:"not null default '' comment('评价类型-KJ 砍价评论-PT 拼团评论') VARCHAR(10)"`
}
