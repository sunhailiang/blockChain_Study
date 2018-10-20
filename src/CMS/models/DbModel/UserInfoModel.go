package DbModel

import "time"

type Userinfo struct {
	Id        int       //用户编号
	UserName  string    //用户名
	UserPwd   string    //用户密码
	Remark    string    //备注
	AddDate   time.Time //添加日期
	ModifDate time.Time //修改日期
	DelFlag   int       //删除标记
	Roles     []*RoleInfo `orm:"rel(m2m)"`
}
