package shop

type MchCommonCat struct {
	Id       int    `xorm:"not null INT(11)"`
	StoreId  int    `xorm:"not null INT(11)"`
	Name     string `xorm:"not null VARCHAR(255)"`
	Sort     int    `xorm:"not null default 1000 comment('排序，升序') INT(11)"`
	IsDelete int    `xorm:"not null default 0 TINYINT(1)"`
}
