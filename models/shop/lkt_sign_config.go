package shop

import (
	"time"
)

type SignConfig struct {
	Id           int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId      int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	IsStatus     int       `xorm:"not null default 0 comment('签到插件是否启用 0：未启用 1：启用') TINYINT(4)"`
	ScoreNum     int       `xorm:"not null default 0 comment('签到次数') INT(11)"`
	Starttime    string    `xorm:"not null default '' comment('签到有效开始时间') CHAR(20)"`
	Endtime      string    `xorm:"not null default '' comment('签到有效结束时间') CHAR(20)"`
	IsRemind     int       `xorm:"not null default 0 comment('是否提醒 0：不提醒 1：提醒') TINYINT(4)"`
	Reset        string    `xorm:"not null default '00:00' comment('重置时间') VARCHAR(20)"`
	Imgurl       string    `xorm:"not null comment('图片') VARCHAR(200)"`
	Score        int       `xorm:"not null default 0 comment('领取积分') INT(11)"`
	Continuity   string    `xorm:"comment('连续设置') TEXT"`
	Detail       string    `xorm:"comment('签到详情') TEXT"`
	Instructions string    `xorm:"comment('积分使用说明') TEXT"`
	ModifyDate   time.Time `xorm:"comment('修改时间') TIMESTAMP"`
}
