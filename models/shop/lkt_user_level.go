package shop

type UserLevel struct {
	Id                int     `xorm:"not null pk autoincr index INT(10)"`
	StoreId           int     `xorm:"comment('商城id') INT(11)"`
	UserId            string  `xorm:"comment('用户id') CHAR(15)"`
	Referee           string  `xorm:"comment('推荐人') CHAR(15)"`
	SuperiorAll       string  `xorm:"comment('隶属上级') VARCHAR(200)"`
	DirectIdAll       string  `xorm:"comment('直推人') VARCHAR(5000)"`
	IndirectIdAll     string  `xorm:"comment('间推人') VARCHAR(5000)"`
	TeamIdAl          string  `xorm:"comment('团队所有人') VARCHAR(8000)"`
	IsVip             int     `xorm:"not null default 1 comment('用户级别') INT(11)"`
	UserDirectJine    string  `xorm:"not null default 80 comment('用户直推金额') DECIMAL(10)"`
	UserIndirectPoint float32 `xorm:"default 0.5 comment('用户间推返佣比例') FLOAT"`
	CardDirectPoint   float32 `xorm:"not null default 0.09 comment('信用卡直推返佣比例') FLOAT"`
	CardIndirectPoint float32 `xorm:"not null default 0.045 comment('信用卡间推返佣比例') FLOAT"`
	ShopDirectPoint   float32 `xorm:"not null default 0.04 comment('商城直推返佣比例') FLOAT"`
	ShopIndirecPoint  float32 `xorm:"not null default 0.02 comment('商城间推返佣比例') FLOAT"`
	SellerDirectPoint float32 `xorm:"not null default 0.01 comment('线下直推销售返佣比例') FLOAT"`
	ChangeDirectPoint float32 `xorm:"not null default 0.01 comment('积分兑换直推返佣比例') FLOAT"`
	UserProfit        string  `xorm:"default 0 comment('推人总返利') DECIMAL(10)"`
	CardProfit        string  `xorm:"default 0 comment('信用卡总返利') DECIMAL(10)"`
	ShopProfit        string  `xorm:"default 0 comment('商城消费总返利') DECIMAL(10)"`
	SellerProfit      string  `xorm:"default 0 comment('商家销售总返利') DECIMAL(10)"`
	ChangeProfit      string  `xorm:"default 0 comment('积分兑换总返利') DECIMAL(10)"`
}
