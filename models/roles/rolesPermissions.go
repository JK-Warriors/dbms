package roles

import (
	"dbms/models"
	"dbms/utils"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RolePermission struct {
	Id           int64 `orm:"pk;"`
	Roleid       int64 `orm:"column(role_id)"`
	Permissionid int64 `orm:"column(permission_id)"`
}

func (this *RolePermission) TableName() string {
	return models.TableName("role_permission")
}

func init() {
	orm.RegisterModel(new(RolePermission))
}

func AddRolePermission(upd RolePermission) error {
	o := orm.NewOrm()
	permission := new(RolePermission)

	permission.Id = upd.Id
	permission.Roleid = upd.Roleid
	permission.Permissionid = upd.Permissionid
	_, err := o.Insert(permission)
	return err
}

func DeleteRolePermission(id int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM "+models.TableName("role_permission")+" WHERE id = ?", id).Exec()
	return err
}
func DeleteRolePermissionForRoleid(roleid int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM "+models.TableName("role_permission")+" WHERE role_id = ?", roleid).Exec()
	return err
}

func ListRolePermission(roleid int64) (ops []RolePermission) {
	var permissions []RolePermission

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("role_permission"))
	cond := orm.NewCondition()
	cond = cond.And("roleid", roleid)
	qs = qs.SetCond(cond)
	qs.All(&permissions)

	return permissions
}

func ListRoleUserPermission(roleid string) (num int64, err error, ops []Permissions) {
	var users []Permissions
	err = utils.GetCache("ListRoleUserPermission.id."+fmt.Sprintf("%d", roleid), &users)
	var nums int64
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		sql := `select p.*
				from pms_role_permission AS rp, pms_permissions AS p
				where rp.permission_id = p.id
				and rp.role_id IN ( ? )
				order by p.id`
		o := orm.NewOrm()
		nums, err = o.Raw(sql, roleid).QueryRows(&users)
		utils.SetCache("ListRoleUserPermission.id."+fmt.Sprintf("%d", roleid), users, cache_expire)
	}
	return nums, err, users
}

type Urls struct {
	Id       int64
	Parentid int64
	Name     string
	EName    string
	Url      string
	Icon     string
	Nav      int
	IsShow   int64
	Sort     int
	IsActive int
}

func GetLeftNav(roleid string, url string) (num int64, err error, ops []Urls) {
	var urls []Urls
	var nums int64
	sql := `select p1.*, instr(url,'` + url + `') as is_active
			from pms_role_permission AS rp, pms_permissions AS p1
			where rp.permission_id = p1.id
			and rp.role_id IN ( ? )
			and p1.parent_id in (select p2.parent_id from pms_permissions AS p2 where p2.url like '` + url + `%')
			order by p1.id`
	o := orm.NewOrm()
	nums, err = o.Raw(sql, roleid).QueryRows(&urls)

	return nums, err, urls
}
