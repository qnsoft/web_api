package shop

type ChannelBankKft struct {
	Id         int    `xorm:"not null pk autoincr comment('主键') INT(10)"`
	Bankcode   string `xorm:"not null comment('银行代码') VARCHAR(255)"`
	Bankname   string `xorm:"not null comment('银行名称') VARCHAR(255)"`
	D0freerate string `xorm:"not null comment('D0银行代扣费率') DECIMAL(10,4)"`
	D0fixrate  int    `xorm:"not null comment('D0固定费') INT(11)"`
	T1freerate string `xorm:"not null DECIMAL(10)"`
	T1fixrate  int    `xorm:"not null INT(11)"`
	Channelid  int    `xorm:"not null comment('通道') INT(11)"`
	D0myrate   string `xorm:"not null comment('臻方便代付费率') DECIMAL(10,4)"`
}
