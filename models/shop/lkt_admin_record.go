package shop

import (
	"time"
)

type AdminRecord struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"default 0 comment('商城id') INT(11)"`
	AdminName string    `xorm:"comment('管理员账号') VARCHAR(50)"`
	Event     string    `xorm:"comment('事件') VARCHAR(200)"`
	Type      int       `xorm:"comment('类型 0:登录/退出 1:添加 2:修改 3:删除 4:导出 5:启用/禁用 6:通过/拒绝') TINYINT(4)"`
	AddDate   time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('操作时间') TIMESTAMP"`
}
