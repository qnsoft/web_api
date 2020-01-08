package shop

type BackgroundColor struct {
	Id        int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   string `xorm:"not null default '0' comment('商城id') CHAR(11)"`
	ColorName string `xorm:"comment('颜色名称') CHAR(10)"`
	Color     string `xorm:"comment('颜色') VARCHAR(30)"`
	Sort      int    `xorm:"comment('排序') INT(11)"`
	Status    int    `xorm:"not null default 0 comment('状态 0：未启用 1：启用') TINYINT(4)"`
}
