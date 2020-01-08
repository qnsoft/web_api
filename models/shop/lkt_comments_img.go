package shop

import (
	"time"
)

type CommentsImg struct {
	Id          int       `xorm:"not null pk autoincr comment('图片id') INT(11)"`
	CommentsUrl string    `xorm:"comment('评论图片') VARCHAR(100)"`
	CommentsId  int       `xorm:"comment('评论id') INT(11)"`
	AddTime     time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Type        int       `xorm:"default 0 comment('类型 0:评论  1:追评') TINYINT(3)"`
}
