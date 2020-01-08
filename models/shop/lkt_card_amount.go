package shop

type CardAmount struct {
	Id       int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	UserId   string `xorm:"not null comment('用户id') VARCHAR(20)"`
	UserT    int    `xorm:"not null comment('手续费') INT(11)"`
	UserEdfl int    `xorm:"not null comment('大额费率') INT(11)"`
	UserSmfl int    `xorm:"not null comment('小额费率') INT(11)"`
	UserLvok int    `xorm:"not null comment('是否启用') INT(11)"`
}
