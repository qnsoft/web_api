package shop

type Option struct {
	Id      int64  `xorm:"pk autoincr comment('id') BIGINT(20)"`
	StoreId int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	Group   string `xorm:"not null default '' comment('类型') VARCHAR(255)"`
	Name    string `xorm:"not null comment('软件显示类型') VARCHAR(255)"`
	Value   string `xorm:"not null comment('操作值') LONGTEXT"`
}
