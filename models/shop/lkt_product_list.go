package shop

import (
	"time"
)

type ProductList struct {
	Id                   int       `xorm:"not null pk autoincr comment('商品id') INT(11)"`
	StoreId              int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	ProductNumber        string    `xorm:"not null comment('商品编号') VARCHAR(100)"`
	ProductTitle         string    `xorm:"not null default '' comment('商品名字') VARCHAR(100)"`
	Subtitle             string    `xorm:"comment('副标题') VARCHAR(100)"`
	Label                string    `xorm:"not null comment('商品标签') VARCHAR(100)"`
	Scan                 string    `xorm:"not null comment('条形码') VARCHAR(250)"`
	ProductClass         string    `xorm:"not null comment('产品类别') VARCHAR(32)"`
	Imgurl               string    `xorm:"not null default '' comment('产品图片') VARCHAR(200)"`
	Content              string    `xorm:"not null comment('产品内容') TEXT"`
	Sort                 int       `xorm:"comment('排序') INT(11)"`
	AddDate              time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Volume               int       `xorm:"not null default 0 comment('销量') INT(12)"`
	DisplayPosition      string    `xorm:"comment('显示位置 1:首页 2：购物车') VARCHAR(50)"`
	SType                string    `xorm:"comment('产品值属性 1：新品,2：热销，3：推荐') VARCHAR(20)"`
	Num                  int       `xorm:"not null default 0 comment('数量') INT(11)"`
	MinInventory         int       `xorm:"not null default 10 comment('库存预警') INT(11)"`
	Status               int       `xorm:"not null default 1 comment('状态 0::上架 1:下架') TINYINT(3)"`
	BrandId              int       `xorm:"not null default 0 comment('品牌ID') INT(11)"`
	IsDistribution       int       `xorm:"not null default 0 comment('是否为分销商品') INT(2)"`
	IsDefaultRatio       int       `xorm:"not null default 0 comment('是否默认比例') INT(2)"`
	Keyword              string    `xorm:"comment('关键词') VARCHAR(100)"`
	Weight               string    `xorm:"comment('重量') VARCHAR(100)"`
	DistributorId        int       `xorm:"default 0 comment('分销等级id 购买就升级') INT(5)"`
	Freight              string    `xorm:"not null comment('运费') TEXT"`
	IsZhekou             int       `xorm:"default 0 comment('是否开启会员') INT(2)"`
	SeparateDistribution string    `xorm:"default '0' comment('单独分销') VARCHAR(50)"`
	Recycle              int       `xorm:"not null default 0 comment('回收站 0.显示 1.回收') TINYINT(4)"`
	Gongyingshang        string    `xorm:"not null default '' comment('供应商') VARCHAR(100)"`
	IsHexiao             int       `xorm:"default 0 comment('是否支持线下核销:0--不支持　1--支持') TINYINT(4)"`
	Hxaddress            string    `xorm:"comment('核销地址') VARCHAR(255)"`
	Active               string    `xorm:"default '1' comment('支持活动:1--正价商品 2--支持拼团 3--支持砍价 4--支持竞拍') CHAR(10)"`
	MchId                int       `xorm:"default 0 comment('商户ID') INT(10)"`
	MchStatus            int       `xorm:"not null default 1 comment('审核状态：1.待审核，2.审核通过，3.审核不通过，4.暂不审核') TINYINT(2)"`
	RefuseReasons        string    `xorm:"comment('拒绝原因') VARCHAR(255)"`
	SearchNum            int       `xorm:"not null default 0 comment('搜索次数') INT(11)"`
	ShowAdr              string    `xorm:"not null default '1' comment('展示位置:1.
首页 2.购物车') VARCHAR(100)"`
	Publisher string `xorm:"comment('发布人') VARCHAR(255)"`
	ProType   int    `xorm:"not null default 1 comment('类型：1线上 0 附近，2课堂') TINYINT(1)"`
}
