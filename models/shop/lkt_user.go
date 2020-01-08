package shop

import (
	"time"
)

type User struct {
	Id               int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId          int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId           string    `xorm:"comment('用户id') CHAR(15)"`
	UserName         []byte    `xorm:"comment('用户昵称') VARBINARY(100)"`
	AccessId         string    `xorm:"comment('授权id') VARCHAR(255)"`
	AccessKey        string    `xorm:"comment('授权密钥') VARCHAR(32)"`
	WxId             string    `xorm:"comment('微信id') VARCHAR(50)"`
	WxName           []byte    `xorm:"comment('微信昵称') VARBINARY(150)"`
	GzhId            string    `xorm:"default '' comment('公众号id') VARCHAR(50)"`
	ZfbId            string    `xorm:"comment('支付宝id') VARCHAR(50)"`
	BdId             string    `xorm:"comment('百度id') VARCHAR(50)"`
	TtId             string    `xorm:"comment('头条id') VARCHAR(50)"`
	Sex              int       `xorm:"comment('性别 0:未知 1:男 2:女') INT(11)"`
	Headimgurl       string    `xorm:"comment('微信头像') MEDIUMTEXT"`
	Province         string    `xorm:"default '' comment('省') VARCHAR(50)"`
	City             string    `xorm:"default '' comment('市') VARCHAR(50)"`
	County           string    `xorm:"default '' comment('县') VARCHAR(50)"`
	DetailedAddress  string    `xorm:"comment('详细地址') VARCHAR(100)"`
	Money            string    `xorm:"default 0.00 comment('金额') DECIMAL(12,2)"`
	Zhenzhu          string    `xorm:"default 0.00 comment('臻珠') DECIMAL(12,2)"`
	Score            string    `xorm:"default 0.00 comment('积分') DECIMAL(12,2)"`
	Password         string    `xorm:"comment('支付密码') CHAR(32)"`
	RegisterData     time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('注册时间') TIMESTAMP"`
	EMail            string    `xorm:"comment('邮箱') VARCHAR(30)"`
	RealName         string    `xorm:"comment('真实姓名') VARCHAR(100)"`
	Mobile           string    `xorm:"comment('手机') VARCHAR(20)"`
	Birthday         string    `xorm:"default '' comment('生日') VARCHAR(32)"`
	WechatId         string    `xorm:"default '' comment('微信号') VARCHAR(50)"`
	Address          string    `xorm:"comment('地址') VARCHAR(300)"`
	BankName         string    `xorm:"comment('银行名称') VARCHAR(30)"`
	Cardholder       string    `xorm:"comment('持卡人') VARCHAR(30)"`
	BankCardNumber   string    `xorm:"comment('银行卡号') VARCHAR(30)"`
	ShareNum         int       `xorm:"default 0 comment('分享次数') INT(11)"`
	Referee          string    `xorm:"comment('推荐人') CHAR(15)"`
	AccessToken      string    `xorm:"default '' comment('访问令牌') VARCHAR(32)"`
	ConsumerMoney    string    `xorm:"default 0.00 comment('消费金') DECIMAL(12,2)"`
	ImgToken         string    `xorm:"comment('分享图片id') VARCHAR(32)"`
	RelationId       string    `xorm:"default '' comment('淘宝渠道id') VARCHAR(32)"`
	Zhanghao         string    `xorm:"comment('账号') VARCHAR(32)"`
	Mima             string    `xorm:"comment('密码') VARCHAR(32)"`
	Source           int       `xorm:"not null default 1 comment('来源 1.小程序 2.app') TINYINT(4)"`
	LoginNum         int       `xorm:"default 0 comment('登录次数') INT(11)"`
	VerificationTime time.Time `xorm:"comment('验证支付密码时间') TIMESTAMP"`
	Parameter        string    `xorm:"comment('参数') TEXT"`
	IsVip            int       `xorm:"not null default 0 comment('是否vip  0升级城主 1城主 2 合伙人 3 股东 ') INT(11)"`
	IsLock           int       `xorm:"default 0 comment('是否冻结 0-不冻结 1-冻结') TINYINT(2)"`
}
