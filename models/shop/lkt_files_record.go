package shop

import (
	"time"
)

type FilesRecord struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	StoreId    string    `xorm:"not null default '0' comment('商城id') CHAR(11)"`
	StoreType  string    `xorm:"not null default '0' comment('平台类型') CHAR(11)"`
	Group      string    `xorm:"not null default '' comment('分组') VARCHAR(50)"`
	UploadMode string    `xorm:"not null comment('上传方式 1:本地 2:阿里云 3:腾讯云 4:七牛云') VARCHAR(50)"`
	ImageName  string    `xorm:"not null comment('图片名称') LONGTEXT"`
	AddTime    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	ShopId     int       `xorm:"INT(11)"`
}
