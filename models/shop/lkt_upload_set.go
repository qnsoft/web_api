package shop

type UploadSet struct {
	Id        int64  `xorm:"pk autoincr comment('id') BIGINT(20)"`
	StoreId   string `xorm:"not null default '0' comment('商城id') CHAR(11)"`
	Type      string `xorm:"comment('类别') CHAR(30)"`
	Attr      string `xorm:"not null comment('属性') CHAR(100)"`
	Attrvalue string `xorm:"not null default '' comment('值') VARCHAR(100)"`
}
