package shop

import (
	"time"
)

type AuctionProduct struct {
	Id           int       `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	StoreId      int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	BrandId      int       `xorm:"comment('品牌id') INT(11)"`
	Title        string    `xorm:"comment('竞拍标题') VARCHAR(200)"`
	Content      string    `xorm:"comment('内容描述') TEXT"`
	Starttime    time.Time `xorm:"comment('开始时间') TIMESTAMP"`
	Endtime      time.Time `xorm:"comment('结束时间') TIMESTAMP"`
	Status       int       `xorm:"default 0 comment('竞拍状态：0-未开始 1-进行中 2-已结束 3-已流拍') TINYINT(2)"`
	Price        string    `xorm:"comment('竞拍起价') DECIMAL(12,2)"`
	AddPrice     string    `xorm:"comment('竞拍加价') DECIMAL(12,2)"`
	CurrentPrice string    `xorm:"comment('当前价格') DECIMAL(12,2)"`
	Imgurl       string    `xorm:"comment('图片路径') VARCHAR(255)"`
	Attribute    string    `xorm:"comment('产品及规格') TEXT"`
	MarketPrice  string    `xorm:"comment('市场价') DECIMAL(12,2)"`
	Days         int       `xorm:"comment('保留显示的天数') INT(11)"`
	InvalidTime  time.Time `xorm:"comment('结束显示的时间') TIMESTAMP"`
	Promise      string    `xorm:"default 99.00 comment('保证金') DECIMAL(12,2)"`
	Pepole       int       `xorm:"default 0 comment('参与人数') INT(11)"`
	UserId       string    `xorm:"comment('最终成交人') CHAR(50)"`
	IsBuy        int       `xorm:"default 0 comment('是否付款，0-未付款，1-付款') TINYINT(2)"`
	TradeNo      string    `xorm:"comment('订单编号') CHAR(32)"`
	LowPepole    int       `xorm:"comment('最低开拍人数') INT(11)"`
	WaitTime     int       `xorm:"comment('出价等待时间（单位是秒）') INT(11)"`
	MchId        int       `xorm:"default 0 comment('店铺id') INT(11)"`
}
