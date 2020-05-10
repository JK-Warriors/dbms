package group

import (
	//"fmt"
	"dbms/models"
	//"dbms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Group struct {
	Id        int    `orm:"pk;column(id);"`
	GroupName string `orm:"column(group_name);"`
	IsDelete  int    `orm:"column(is_delete);"`
	Created   int64  `orm:"column(created);"`
	Updated   int64  `orm:"column(updated);"`
}

func (this *Group) TableName() string {
	return models.TableName("group")
}
func init() {
	orm.RegisterModel(new(Group))
}

//添加业务系统
func AddGroup(gp Group) error {
	o := orm.NewOrm()
	o.Using("default")
	gpconf := new(Group)

	gpconf.GroupName = gp.GroupName
	gpconf.IsDelete = 0
	gpconf.Created = time.Now().Unix()
	_, err := o.Insert(gpconf)
	return err
}

//验证系统名称是否已经存在
func CheckNameExist(name string) int {
	o := orm.NewOrm()
	var li_count int32
	_ = o.Raw("select count(1) from pms_group WHERE group_name =?", name).QueryRow(&li_count)

	if li_count == 0 {
		return 0
	}
	return 1
}

//修改分组
func UpdateGroup(id int, gp Group) error {
	var gpconf Group
	o := orm.NewOrm()
	gpconf = Group{Id: id}

	gpconf.GroupName = gp.GroupName
	gpconf.Updated = time.Now().Unix()

	_, err := o.Update(&gpconf)
	return err
}

//获取分组
func GetGroup(id int) (Group, error) {
	var gpconf Group
	var err error
	o := orm.NewOrm()

	gpconf = Group{Id: id}
	err = o.Read(&gpconf)

	if err == orm.ErrNoRows {
		return gpconf, nil
	}
	return gpconf, err
}

func GetGroupName(id int) string {
	var group_name string
	o := orm.NewOrm()

	err := o.Raw("select group_name from pms_group where id = ?", id).QueryRow(&group_name)

	if err == orm.ErrNoRows {
		return ""
	}
	return group_name
}

//获取分组
func ListGroup(condArr map[string]string, page int, offset int) (num int64, err error, gpconf []Group) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("group"))
	cond := orm.NewCondition()

	if condArr["search_name"] != "" {
		cond = cond.And("group_name__icontains", condArr["search_name"])
	}

	cond = cond.And("is_delete", 0)

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("id")
	nums, errs := qs.Limit(offset, start).All(&gpconf)
	return nums, errs, gpconf
}

func ListAllGroup() (gpconf []Group) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("group"))
	cond := orm.NewCondition()

	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&gpconf)
	return gpconf
}

//统计数量
func CountGroup(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("group"))
	cond := orm.NewCondition()

	if condArr["group_name"] != "" {
		cond = cond.And("group_name__icontains", condArr["group_name"])
	}
	cond = cond.And("is_delete", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

func DeleteGroup(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("group") + " WHERE id IN(" + ids + ")").Exec()
	return err
}
