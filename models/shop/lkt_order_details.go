package shop

import (
	"time"
)

type OrderDetails struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId     int       `xorm:"default 0 comment('商城id') INT(11)"`
	UserId      string    `xorm:"comment('用户id') CHAR(15)"`
	PId         int       `xorm:"comment('产品id') INT(11)"`
	PName       string    `xorm:"comment('产品名称') VARCHAR(200)"`
	PPrice      string    `xorm:"default 0.00 comment('产品价格') DECIMAL(12,2)"`
	Num         int       `xorm:"comment('数量') INT(11)"`
	Unit        string    `xorm:"comment('单位') VARCHAR(50)"`
	RSno        string    `xorm:"comment('订单号') VARCHAR(255)"`
	AddTime     time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	DeliverTime time.Time `xorm:"comment('发货时间') TIMESTAMP"`
	ArriveTime  time.Time `xorm:"comment('到货时间') TIMESTAMP"`
	RStatus     int       `xorm:"default 0 comment(' '状态 0：未付款 1：未发货 2：待收货 3：待评论 4：退货 5:已完成 6 订单关闭 9拼团中 10 拼团失败-未退款 11 拼团失败-已退款'') TINYINT(4)"`
	Content     string    `xorm:"comment('退货原因') TEXT"`
	Reason      string    `xorm:"comment('app退款原因') CHAR(60)"`
	ReType      int       `xorm:"default 0 comment('退款类型 1:退货退款  2:退款') TINYINT(4)"`
	ReMoney     float32   `xorm:"comment('退款金额') FLOAT(12,2)"`
	ReTime      time.Time `xorm:"comment('申请退款时间') DATETIME"`
	RePhoto     string    `xorm:"comment('上传凭证') TEXT"`
	RContent    string    `xorm:"comment('拒绝原因') TEXT"`
	RType       int       `xorm:"default 0 comment('类型 0:审核中 1:同
意并让用户寄回 2:拒绝(退货退款) 3:用户已快递 4:收到寄回商品,同意并退款 5：拒绝
并退回商品 8:拒绝(退款) 9:同意并退款') TINYINT(4)"`
	ExpressId  int    `xorm:"comment('快递公司id') INT(255)"`
	CourierNum string `xorm:"comment('快递单号') VARCHAR(40)"`
	Freight    string `xorm:"default 0.00 comment('运费') DECIMAL(12,2)"`
	Size       string `xorm:"comment('配置名称') VARCHAR(100)"`
	Sid        string `xorm:"CHAR(11)"`
}
