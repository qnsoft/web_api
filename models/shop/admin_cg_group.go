package shop

type AdminCgGroup struct {
	Groupid     int    `xorm:"not null INT(11)"`
	GCname      string `xorm:"VARCHAR(50)"`
	GParentid   int    `xorm:"INT(11)"`
	GShoworder  int    `xorm:"INT(11)"`
	GLevel      int    `xorm:"INT(11)"`
	GChildcount int    `xorm:"INT(11)"`
	GDelete     int    `xorm:"default 0 INT(11)"`
	GNum        int    `xorm:"default 0 INT(11)"`
}
