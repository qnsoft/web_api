package shop

type Hotkeywords struct {
	Id      int    `xorm:"not null pk INT(11)"`
	Keyword string `xorm:"not null default '' CHAR(60)"`
}
