package shop

import (
	"time"
)

type Third struct {
	Id           int       `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Ticket       string    `xorm:"comment('获取token的凭证') VARCHAR(255)"`
	TicketTime   time.Time `xorm:"comment('凭证更新时间') TIMESTAMP"`
	Token        string    `xorm:"comment('授权token') VARCHAR(255)"`
	TokenExpires int       `xorm:"comment('token过期时间戳') INT(11)"`
	Appid        string    `xorm:"comment('第三方平台appid') VARCHAR(50)"`
	Appsecret    string    `xorm:"comment('第三方平台秘钥') VARCHAR(100)"`
	CheckToken   string    `xorm:"comment('消息检验Token，第三方平台设置') VARCHAR(255)"`
	EncryptKey   string    `xorm:"comment('消息加减密key') VARCHAR(255)"`
	ServeDomain  string    `xorm:"comment('服务器域名') TEXT"`
	WorkDomain   string    `xorm:"comment('业务域名') TEXT"`
}
