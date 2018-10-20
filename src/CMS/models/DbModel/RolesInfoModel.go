package DbModel

import (
	"time"
)

type RoleInfo struct {
	Id        int
	RoleName  string      `orm:"size(32)"`
	Remark    string
	DelFlag   int
	AddDate   time.Time
	ModifDate time.Time
	Users     []*Userinfo `orm:"reverse(many)"`
}
