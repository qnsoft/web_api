package shop

import (
	"time"
)

type Config struct {
	Id              int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId         int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	IsRegister      int       `xorm:"not null default 0 comment('是否需要注册 1.注册 2.免注册') TINYINT(4)"`
	Logo            string    `xorm:"comment('公司logo') VARCHAR(100)"`
	Logo1           string    `xorm:"comment('首页logo') VARCHAR(100)"`
	Company         string    `xorm:"not null default '' comment('公司名称') VARCHAR(100)"`
	Appid           string    `xorm:"not null default '' comment('小程序id') CHAR(18)"`
	Appsecret       string    `xorm:"not null default '' comment('小程序密钥') CHAR(32)"`
	Domain          string    `xorm:"not null default '' comment('小程序域名') VARCHAR(100)"`
	AppDomainName   string    `xorm:"not null comment('APP域名') VARCHAR(255)"`
	MchId           int       `xorm:"comment('商户id') INT(11)"`
	Ip              string    `xorm:"comment('ip地址') VARCHAR(25)"`
	UploadimgDomain string    `xorm:"not null default '' comment('图片上传域名') VARCHAR(100)"`
	Upserver        int       `xorm:"not null default 2 comment('上传服务器:1,本地　2,阿里云 3,腾讯云 4,七牛云') TINYINT(4)"`
	Uploadimg       string    `xorm:"not null default '' comment('图片上传位置') VARCHAR(100)"`
	UploadFile      string    `xorm:"not null default '' comment('文件上传位置') VARCHAR(100)"`
	ModifyDate      time.Time `xorm:"comment('修改时间') TIMESTAMP"`
	MchKey          string    `xorm:"not null comment('微信支付key') VARCHAR(100)"`
	MchCert         string    `xorm:"not null comment('支付证书文件地址') VARCHAR(100)"`
	UserId          string    `xorm:"default '' comment('用户ID默认前缀') VARCHAR(20)"`
	WxName          string    `xorm:"comment('微信默认用户名称') VARCHAR(255)"`
	WxHeadimgurl    string    `xorm:"comment('微信默认用户头像') VARCHAR(255)"`
	CustomerService string    `xorm:"comment('客服') TEXT"`
	Agreement       string    `xorm:"comment('用户协议') TEXT"`
	Aboutus         string    `xorm:"comment('关于我们') TEXT"`
}
