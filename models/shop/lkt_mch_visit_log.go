package shop

import (
	"time"
)

type MchVisitLog struct {
	Id        int64     `xorm:"BIGINT(20)"`
	UserId    int       `xorm:"not null INT(11)"`
	MchId     int       `xorm:"not null INT(11)"`
	Addtime   int       `xorm:"not null INT(11)"`
	VisitDate time.Time `xorm:"not null DATE"`
}
