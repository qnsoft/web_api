package shop

import (
	"time"
)

type ThirdTemplate struct {
	Id           int       `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	TemplateId   string    `xorm:"comment('微信端模板id') VARCHAR(50)"`
	WxDesc       string    `xorm:"comment('wx模板描述') VARCHAR(100)"`
	WxVersion    string    `xorm:"comment('wx模板版本号') VARCHAR(100)"`
	WxCreateTime int       `xorm:"comment('添加进开放平台模板库时间戳') INT(11)"`
	LkVersion    string    `xorm:"comment('后台定义的模板编号') VARCHAR(255)"`
	LkDesc       string    `xorm:"comment('后台模板描述') TEXT"`
	ImgUrl       string    `xorm:"comment('模板图片路劲') VARCHAR(150)"`
	StoreId      int       `xorm:"comment('商城id') INT(11)"`
	TradeData    string    `xorm:"comment('数据字典-行业') VARCHAR(50)"`
	IsUse        int       `xorm:"default 1 comment('是否应用 0：不应用 1：应用') TINYINT(2)"`
	UpdateTime   time.Time `xorm:"comment('模板更新时间') DATETIME"`
	Title        string    `xorm:"comment('模板名称') VARCHAR(255)"`
}
