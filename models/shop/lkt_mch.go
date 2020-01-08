package shop

import (
	"time"
)

type Mch struct {
	Id              int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId         int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId          string    `xorm:"comment('用户id') CHAR(15)"`
	Name            string    `xorm:"comment('店铺名称') VARCHAR(255)"`
	ShopInformation string    `xorm:"comment('店铺信息') TEXT"`
	ShopRange       string    `xorm:"comment('经营范围') TEXT"`
	Realname        string    `xorm:"comment('真实姓名') VARCHAR(255)"`
	IdNumber        string    `xorm:"comment('身份证号码') VARCHAR(255)"`
	Tel             string    `xorm:"comment('联系电话') VARCHAR(255)"`
	Address         string    `xorm:"comment('联系地址') VARCHAR(255)"`
	Logo            string    `xorm:"comment('店铺Logo') VARCHAR(255)"`
	ShopNature      int       `xorm:"not null default 0 comment('店铺性质：0.个人 1.企业') TINYINT(3)"`
	BusinessLicense string    `xorm:"comment('营业执照') VARCHAR(255)"`
	AddTime         time.Time `xorm:"comment('申请时间') TIMESTAMP"`
	ReviewTime      time.Time `xorm:"comment('审核时间') TIMESTAMP"`
	ReviewStatus    int       `xorm:"not null default 0 comment('审核状态：0.待审核 1.审核通过 2.审核不通过') TINYINT(3)"`
	ReviewResult    string    `xorm:"comment('拒绝理由') VARCHAR(255)"`
	AccountMoney    string    `xorm:"not null default 0.00 comment('商户余额') DECIMAL(12,2)"`
	CollectionNum   int       `xorm:"default 0 comment('收藏数量') INT(11)"`
	IsOpen          int       `xorm:"not null default 0 comment('是否营业：0.否 1.是') TINYINT(1)"`
	IsLock          int       `xorm:"not null default 0 comment('是否被系统关闭：0.否 1.是') TINYINT(1)"`
	Confines        string    `xorm:"comment('店铺经营范围') VARCHAR(255)"`
	BusinessHours   string    `xorm:"comment('营业时间') VARCHAR(255)"`
	Lon             float32   `xorm:"comment('经度') FLOAT"`
	Lat             float32   `xorm:"comment('纬度') FLOAT"`
}
