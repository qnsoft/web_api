package shop

import (
	"time"
)

type Article struct {
	ArticleId     int       `xorm:"not null pk autoincr comment('文章id') INT(11)"`
	StoreId       int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	ArticleTitle  string    `xorm:"not null default '' comment('文章标题') VARCHAR(100)"`
	ArticlePrompt string    `xorm:"not null default '' comment('文章提示') VARCHAR(100)"`
	ArticleImgurl string    `xorm:"not null default '' comment('文章图片') VARCHAR(200)"`
	Content       string    `xorm:"not null comment('文章内容') TEXT"`
	Sort          int       `xorm:"comment('排序') INT(11)"`
	AddDate       time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	ShareNum      int       `xorm:"not null default 0 comment('分享次数') INT(11)"`
	TotalAmount   string    `xorm:"not null default 0.00 comment('红包总金额') DECIMAL(12,2)"`
	TotalNum      int       `xorm:"not null default 0 comment('红包个数') INT(11)"`
	Wishing       string    `xorm:"not null default '' comment('祝福语') VARCHAR(100)"`
}
