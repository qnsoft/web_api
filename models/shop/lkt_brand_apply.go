package shop

import (
	"time"
)

type BrandApply struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Brandname string    `xorm:"comment('品牌名称') VARCHAR(200)"`
	StoreId   int       `xorm:"INT(11)"`
	ShopId    int       `xorm:"INT(11)"`
	EditUser  string    `xorm:"VARCHAR(50)"`
	Logourl   string    `xorm:"comment('logo 图片名') VARCHAR(100)"`
	AddTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Status    int       `xorm:"default 0 comment('申请状态 0待申，1通过，2拒绝') TINYINT(4)"`
	Mark      string    `xorm:"VARCHAR(500)"`
}
