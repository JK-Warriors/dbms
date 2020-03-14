package roles

import (
	"fmt"
	"opms/models"
	"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Roles struct {
	Id      int64 `orm:"pk;column(id);"`
	Name    string
	Summary string
	Created int64
	Changed int64
}

func (this *Roles) TableName() string {
	return models.TableName("roles")
}

func init() {
	orm.RegisterModel(new(Roles))
}

func AddRole(upd Roles) error {
	o := orm.NewOrm()
	role := new(Roles)

	role.Id = upd.Id
	role.Name = upd.Name
	role.Summary = upd.Summary
	role.Created = time.Now().Unix()
	role.Changed = time.Now().Unix()
	_, err := o.Insert(role)
	return err
}

func UpdateRole(id int64, upd Roles) error {
	var role Roles
	o := orm.NewOrm()
	role = Roles{Id: id}

	role.Name = upd.Name
	role.Summary = upd.Summary
	role.Changed = time.Now().Unix()
	var err error
	_, err = o.Update(&role, "name", "summary", "changed")

	return err
}

func ListRole(condArr map[string]string, page int, offset int) (num int64, err error, ops []Roles) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("roles"))
	cond := orm.NewCondition()

	if condArr["keywords"] != "" {
		cond = cond.And("name__icontains", condArr["keywords"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qs = qs.OrderBy("id")
	var role []Roles
	num, errs := qs.Limit(offset, start).All(&role)
	return num, errs, role
}

func ListRoleAll()(num int64, err error, ops []Roles) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From(models.TableName("roles")).
		OrderBy("id").
		Asc()
	sql := qb.String()

	var roles []Roles
	nums, errs := o.Raw(sql).QueryRows(&roles)
	return nums, errs, roles
}

func CountRole(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("roles"))
	cond := orm.NewCondition()

	if condArr["keywords"] != "" {
		cond = cond.And("name__icontains", condArr["keywords"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}

func GetRole(id int64) (Roles, error) {
	var role Roles
	var err error

	err = utils.GetCache("GetRole.id."+fmt.Sprintf("%d", id), &role)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		o := orm.NewOrm()
		role = Roles{Id: id}
		err = o.Read(&role)
		utils.SetCache("GetRole.id."+fmt.Sprintf("%d", id), role, cache_expire)
	}
	return role, err
}

func DeleteRole(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("roles") + " WHERE id IN(" + ids + ")").Exec()
	return err
}
