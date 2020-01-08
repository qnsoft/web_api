package shop

import (
	"time"
)

type ThirdMiniInfo struct {
	Id                     int       `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	NickName               string    `xorm:"comment('授权方昵称') VARCHAR(100)"`
	AuthorizerAppid        string    `xorm:"comment('授权小程序appid') VARCHAR(100)"`
	AuthorizerAccessToken  string    `xorm:"comment('授权方接口调用凭据（在授权的公众号或小程序具备API权限时，才有此返回值），也简称为令牌') VARCHAR(255)"`
	AuthorizerExpires      int       `xorm:"comment('有效期（在授权的公众号或小程序具备API权限时，才有此返回值）') INT(11)"`
	AuthorizerRefreshToken string    `xorm:"comment('接口调用凭据刷新令牌') VARCHAR(255)"`
	UpdateTime             time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	FuncInfo               string    `xorm:"comment('授权给开发者的权限集列表') LONGTEXT"`
	ExpiresTime            int       `xorm:"comment('过期时间戳') INT(11)"`
	CompanyId              string    `xorm:"comment('公司id') VARCHAR(100)"`
	HeadImg                string    `xorm:"comment('授权方头像') VARCHAR(255)"`
	VerifyTypeInfo         string    `xorm:"comment('授权方认证类型，-1代表未认证，0代表微信认证') VARCHAR(250)"`
	UserName               string    `xorm:"comment('小程序的原始ID') VARCHAR(255)"`
	Signature              string    `xorm:"comment('帐号介绍') VARCHAR(255)"`
	PrincipalName          string    `xorm:"comment('小程序的主体名称') VARCHAR(255)"`
	BusinessInfo           string    `xorm:"comment('开通状况（0代表未开通，1代表已开通）') VARCHAR(255)"`
	QrcodeUrl              string    `xorm:"comment('二维码图片的URL，开发者最好自行也进行保存') VARCHAR(255)"`
	Miniprograminfo        string    `xorm:"comment('根据这个字段判断是否为小程序类型授权') LONGTEXT"`
	StoreId                int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Auditid                string    `xorm:"comment('审核编号') VARCHAR(100)"`
	ReviewMark             int       `xorm:"default 1 comment('审核状态 1：未审核 2：审核中 3：审核失败 4：审核成功') TINYINT(2)"`
	IssueMark              int       `xorm:"default 1 comment('发布状态 1：未发布  2：发布失败 3：发布成功') TINYINT(2)"`
	SubmitTime             time.Time `xorm:"comment('审核提价时间') DATETIME"`
}
