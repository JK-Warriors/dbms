package roles

import (
	"dbms/models"
	"dbms/utils"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RolesUser struct {
	Id     int64 `orm:"pk;auto"`
	Roleid int64 `orm:"column(role_id)"`
	Userid int64 `orm:"column(user_id)"`
}

type RolesUserName struct {
	Id       int64
	Userid   int64
	Realname string
	Roleid   int64
}

func (this *RolesUser) TableName() string {
	return models.TableName("role_user")
}

func init() {
	orm.RegisterModel(new(RolesUser))
}

func AddRolesUser(upd RolesUser) error {
	o := orm.NewOrm()
	user := new(RolesUser)

	user.Roleid = upd.Roleid
	user.Userid = upd.Userid
	_, err := o.Insert(user)
	return err
}

func UpdateRolesUser(upd RolesUser) error {
	o := orm.NewOrm()
	user := new(RolesUser)

	_, err := o.Raw("DELETE FROM "+models.TableName("role_user")+" WHERE user_id = ?", upd.Userid).Exec()
	if err == nil {
		user.Roleid = upd.Roleid
		user.Userid = upd.Userid
		_, err = o.Insert(user)
	}
	return err
}

func DeleteRolesUser(id int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM "+models.TableName("role_user")+" WHERE id = ?", id).Exec()
	return err
}

func ListRolesUserAndName(roleid int64) (num int64, err error, user []RolesUserName) {
	var users []RolesUserName
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("upr.userid", "upr.realname", "gu.roleid", "gu.id").
		From("pms_role_user AS gu").
		LeftJoin("pms_users_profile AS upr").On("upr.userid = gu.userid").
		Where("gu.roleid=?").
		OrderBy("gu.id").
		Asc()
	sql := qb.String()
	o := orm.NewOrm()
	nums, err := o.Raw(sql, roleid).QueryRows(&users)
	return nums, err, users
}

func ListRolesUser(roleid int64, page, offset int) (ops []RolesUser) {
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	var users []RolesUser
	var err error
	err = utils.GetCache("ListRolesUser.id."+fmt.Sprintf("%d", roleid), &users)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		o := orm.NewOrm()
		o.Using("default")
		qs := o.QueryTable(models.TableName("role_user"))
		cond := orm.NewCondition()
		cond = cond.And("roleid", roleid)
		qs = qs.SetCond(cond)
		qs.Limit(offset, start).All(&users)
		utils.SetCache("ListRolesUser.id."+fmt.Sprintf("%d", roleid), users, cache_expire)
	}
	return users
}

func GetRoleIdByUserId(userid int64) (err error, roleUser RolesUser) {
	o := orm.NewOrm()

	qs := o.QueryTable(models.TableName("role_user"))
	cond := orm.NewCondition()
	cond = cond.And("user_id", userid)
	qs = qs.SetCond(cond)
	errs := qs.Limit(1).One(&roleUser)

	return errs, roleUser
}
