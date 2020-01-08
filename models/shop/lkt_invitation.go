package shop

type Invitation struct {
	Id            int    `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId       int    `xorm:"not null default 0 comment('商城id') INT(11)"`
	StoreType     string `xorm:"not null default '' comment('类型') VARCHAR(255)"`
	Status        int    `xorm:"default 0 comment('状态 0：未启用 1：启用') INT(2)"`
	Invitees      string `xorm:"comment('被邀请人可获得的奖励') LONGTEXT"`
	InvitePeople  string `xorm:"comment('邀请人可获得的奖励') LONGTEXT"`
	Restriction   string `xorm:"comment('领取限制') LONGTEXT"`
	PaymentAmount string `xorm:"default 0.00 comment('实物订单支付金额') DECIMAL(10,2)"`
	OrderTime     int    `xorm:"comment('完成下单时间') INT(3)"`
	Data          string `xorm:"comment('素材') TEXT"`
	Rule          string `xorm:"comment('活动规则') TEXT"`
	ShareTitle    string `xorm:"default '' comment('分享标题') VARCHAR(255)"`
	ShareImg      string `xorm:"default '' comment('分享图片') VARCHAR(255)"`
}
