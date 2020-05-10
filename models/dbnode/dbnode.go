package dbnode

import (
	//"fmt"
	"dbms/models"

	//"dbms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Nodeconfig struct {
	Id        int    `orm:"pk;column(id);"`
	Ostype    int    `orm:"column(os_type);"`
	OsVersion string `orm:"column(os_version);"`
	Host      string `orm:"column(host);"`
	Protocol  string `orm:"column(protocol);"`
	Port      int    `orm:"column(port);"`
	Nodename  string `orm:"column(nodename);"`
	Username  string `orm:"column(username);"`
	Password  string `orm:"column(password);"`
	DbId      int    `orm:"column(db_id);"`
	Status    int    `orm:"column(status);"`
	IsDelete  int    `orm:"column(is_delete);"`
	Created   int64  `orm:"column(created);"`
	Updated   int64  `orm:"column(updated);"`
}

func (this *Nodeconfig) TableName() string {
	return models.TableName("node_config")
}
func init() {
	orm.RegisterModel(new(Nodeconfig))
}

//添加节点
func AddNodeConfig(upd Nodeconfig) error {
	o := orm.NewOrm()
	o.Using("default")
	conf := new(Nodeconfig)

	conf.Ostype = upd.Ostype
	conf.OsVersion = upd.OsVersion
	conf.Host = upd.Host
	conf.Protocol = upd.Protocol
	conf.Port = upd.Port
	conf.Nodename = upd.Nodename
	conf.Username = upd.Username
	conf.Password = upd.Password
	conf.DbId = upd.DbId
	conf.Status = 1
	conf.IsDelete = 0
	conf.Created = time.Now().Unix()
	_, err := o.Insert(conf)
	return err
}

//修改节点信息
func UpdateNodeConfig(id int, upd Nodeconfig) error {
	var conf Nodeconfig
	o := orm.NewOrm()
	conf, err := GetNodeConfig(id)
	if err == nil {
		conf.Ostype = upd.Ostype
		conf.OsVersion = upd.OsVersion
		conf.Host = upd.Host
		conf.Protocol = upd.Protocol
		conf.Port = upd.Port
		conf.Nodename = upd.Nodename
		conf.Username = upd.Username
		conf.Password = upd.Password
		conf.DbId = upd.DbId
		conf.Updated = time.Now().Unix()

		_, err = o.Update(&conf)
	}
	return err
}

//获取节点信息
func GetNodeConfig(id int) (Nodeconfig, error) {
	var conf Nodeconfig
	var err error
	o := orm.NewOrm()

	conf = Nodeconfig{Id: id}
	err = o.Read(&conf)

	if err == orm.ErrNoRows {
		return conf, nil
	}
	return conf, err
}

//获取节点列表
func ListNodeConfig(condArr map[string]string, page int, offset int) (num int64, err error, conf []Nodeconfig) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("node_config"))
	cond := orm.NewCondition()

	if condArr["ostype"] != "" {
		cond = cond.And("os_type", condArr["ostype"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["nodename"] != "" {
		cond = cond.And("nodename__icontains", condArr["nodename"])
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
	nums, errs := qs.Limit(offset, start).All(&conf)
	return nums, errs, conf
}

func ListAllNodeConfig() (conf []Nodeconfig) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("node_config"))
	cond := orm.NewCondition()

	cond = cond.And("is_delete", 0)
	qs = qs.SetCond(cond)

	_, _ = qs.OrderBy("id").All(&conf)
	return conf
}

//统计数量
func CountNodeConfig(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("node_config"))
	cond := orm.NewCondition()

	if condArr["ostype"] != "" {
		cond = cond.And("os_type", condArr["ostype"])
	}
	if condArr["host"] != "" {
		cond = cond.And("host__icontains", condArr["host"])
	}
	if condArr["nodename"] != "" {
		cond = cond.And("nodename__icontains", condArr["nodename"])
	}
	cond = cond.And("status", 1)
	cond = cond.And("is_delete", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

func DeleteNodeConfig(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("node_config") + " WHERE id IN(" + ids + ")").Exec()
	return err
}

//更改节点状态
func ChangeNodeConfigStatus(id int, status int) error {
	o := orm.NewOrm()

	conf := Nodeconfig{Id: id}
	err := o.Read(&conf, "id")
	if nil != err {
		return err
	} else {
		conf.Status = status
		_, err := o.Update(&conf)
		return err
	}
}
