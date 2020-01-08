package shop

type Payment struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	Name         string `xorm:"not null comment('支付名称') VARCHAR(50)"`
	Type         int    `xorm:"not null default 1 comment('1:线上、2:线下') TINYINT(1)"`
	ClassName    string `xorm:"not null comment('支付类名称') VARCHAR(50)"`
	Description  string `xorm:"comment('描述') TEXT"`
	Logo         string `xorm:"not null comment('支付方式logo图片路径') VARCHAR(255)"`
	Status       int    `xorm:"not null default 1 comment('状态 0启用 1禁用') TINYINT(1)"`
	Sort         int    `xorm:"not null default 99 comment('排序') SMALLINT(5)"`
	Note         string `xorm:"comment('支付说明') TEXT"`
	Poundage     string `xorm:"not null default 0.00 comment('手续费') DECIMAL(15,2)"`
	PoundageType int    `xorm:"not null default 1 comment('手续费方式 1百分比 2固定值') TINYINT(1)"`
	ClientType   int    `xorm:"not null default 1 comment('1:PC端 2:移动端 3:通用') TINYINT(1)"`
}
