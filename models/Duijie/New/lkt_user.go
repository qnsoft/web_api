package New

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
商城用户信息
*/
type User struct {
	// Id               int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	// StoreId          int       `xorm:"not null default 0 comment('商城id') INT(11)"` //对接必填
	// UserId           string    `xorm:"comment('用户id') CHAR(15)"`                   //对接必填
	// UserName         []byte    `xorm:"comment('用户昵称') VARBINARY(100)"`             //对接必填(注意这里的用户名不是原数据库的username)
	// AccessId         string    `xorm:"comment('授权id') VARCHAR(255)"`
	// AccessKey        string    `xorm:"comment('授权密钥') VARCHAR(32)"`
	// WxId             string    `xorm:"comment('微信id') VARCHAR(50)"`
	// WxName           []byte    `xorm:"comment('微信昵称') VARBINARY(150)"`
	// GzhId            string    `xorm:"default '' comment('公众号id') VARCHAR(50)"`
	// ZfbId            string    `xorm:"comment('支付宝id') VARCHAR(50)"`
	// BdId             string    `xorm:"comment('百度id') VARCHAR(50)"`
	// TtId             string    `xorm:"comment('头条id') VARCHAR(50)"`
	// Sex              int       `xorm:"comment('性别 0:未知 1:男 2:女') INT(11)"`
	// Headimgurl       string    `xorm:"comment('微信头像') MEDIUMTEXT"`
	// Province         string    `xorm:"default '' comment('省') VARCHAR(50)"`
	// City             string    `xorm:"default '' comment('市') VARCHAR(50)"`
	// County           string    `xorm:"default '' comment('县') VARCHAR(50)"`
	// DetailedAddress  string    `xorm:"comment('详细地址') VARCHAR(100)"`
	// Money            string    `xorm:"default 0.00 comment('金额') DECIMAL(12,2)"` //对接必填（后期操作信用卡与商城，同变）
	// Zhenzhu          string    `xorm:"default 0.00 comment('珍珠') DECIMAL(12,2)"`
	// Score            int       `xorm:"default 0 comment('积分') INT(11)"`
	// Password         string    `xorm:"comment('支付密码') CHAR(32)"`
	// RegisterData     time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('注册时间') TIMESTAMP"` //对接必填
	// EMail            string    `xorm:"comment('邮箱') VARCHAR(30)"`                             //对接必填
	// RealName         string    `xorm:"comment('真实姓名') VARCHAR(100)"`                          //对接必填
	// Mobile           string    `xorm:"comment('手机') VARCHAR(20)"`                             //对接必填
	// Birthday         string    `xorm:"default '' comment('生日') VARCHAR(32)"`
	// WechatId         string    `xorm:"default '' comment('微信号') VARCHAR(50)"`
	// Address          string    `xorm:"comment('地址') VARCHAR(300)"`
	// BankName         string    `xorm:"comment('银行名称') VARCHAR(30)"`
	// Cardholder       string    `xorm:"comment('持卡人') VARCHAR(30)"`
	// BankCardNumber   string    `xorm:"comment('银行卡号') VARCHAR(30)"`
	// ShareNum         int       `xorm:"default 0 comment('分享次数') INT(11)"`
	// Referee          string    `xorm:"comment('推荐人') CHAR(15)"`
	// AccessToken      string    `xorm:"default '' comment('访问令牌') VARCHAR(32)"`
	// ConsumerMoney    string    `xorm:"default 0.00 comment('消费金') DECIMAL(12,2)"`
	// ImgToken         string    `xorm:"comment('分享图片id') VARCHAR(32)"`
	// Zhanghao         string    `xorm:"comment('账号') VARCHAR(32)"`                               //对接必填
	// Mima             string    `xorm:"comment('密码') VARCHAR(32)"`                               //对接必填
	// Source           int       `xorm:"not null default 1 comment('来源 1.小程序 2.app') TINYINT(4)"` //对接必填
	// LoginNum         int       `xorm:"default 0 comment('登录次数') INT(11)"`                       //对接必填
	// VerificationTime time.Time `xorm:"comment('验证支付密码时间') TIMESTAMP"`
	// Parameter        string    `xorm:"comment('参数') TEXT"`
	// IsVip            int       `xorm:"default 0 comment('是否VIP 0-不是 1-是') TINYINT(2)"`  //是否vip
	// IsLock           int       `xorm:"default 0 comment('是否冻结 0-不冻结 1-冻结') TINYINT(2)"` //对接必填

	Id               int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId          int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId           string    `xorm:"comment('用户id') CHAR(15)"`
	UserName         []byte    `xorm:"default 我 comment('用户昵称') VARBINARY(100)"`
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
	Zhanghao         string    `xorm:"comment('账号') VARCHAR(32)"`
	Mima             string    `xorm:"comment('密码') VARCHAR(32)"`
	Source           int       `xorm:"not null default 1 comment('来源 1.小程序 2.app') TINYINT(4)"`
	LoginNum         int       `xorm:"default 0 comment('登录次数') INT(11)"`
	VerificationTime time.Time `xorm:"comment('验证支付密码时间') TIMESTAMP"`
	Parameter        string    `xorm:"comment('参数') TEXT"`
	IsVip            int       `xorm:"not null default 0 comment('是否vip  0升级城主 1城主 2 合伙人 3 股东 ') INT(11)"`
	IsLock           int       `xorm:"default 0 comment('是否冻结 0-不冻结 1-冻结') TINYINT(2)"`
}

type User_level struct {
	Id                int     `xorm:"not null index INT(10)"`
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
	ShopDirectPoint   float32 `xorm:"not null default 0.4 comment('商城直推返佣比例') FLOAT"`
	ShopIndirecPoint  float32 `xorm:"not null default 0.2 comment('商城间推返佣比例') FLOAT"`
	SellerDirectPoint float32 `xorm:"not null default 0.1 comment('线下直推销售返佣比例') FLOAT"`
	ChangeDirectPoint float32 `xorm:"not null default 0.1 comment('积分兑换直推返佣比例') FLOAT"`
	UserProfit        string  `xorm:"default 0 comment('推人总返利') DECIMAL(10)"`
	CardProfit        string  `xorm:"default 0 comment('信用卡总返利') DECIMAL(10)"`
	ShopProfit        string  `xorm:"default 0 comment('商城消费总返利') DECIMAL(10)"`
	SellerProfit      string  `xorm:"default 0 comment('商家销售总返利') DECIMAL(10)"`
	ChangeProfit      string  `xorm:"default 0 comment('积分兑换总返利') DECIMAL(10)"`
}
