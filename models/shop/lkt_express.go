package shop

type Express struct {
	Id         int    `xorm:"not null comment('id') INT(11)"`
	KuaidiName string `xorm:"comment('快递名称') VARCHAR(255)"`
	Type       string `xorm:"comment('简称') VARCHAR(255)"`
}
