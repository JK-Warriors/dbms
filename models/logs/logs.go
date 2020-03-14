package logs

import (
	//"fmt"
	"opms/models"
	//"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Logs struct {
	Id           int64 		`orm:"pk;column(id);"`
	Userid       int64		`orm:"column(user_id);"`
	Username     string
	Url          string
	Title        string
	Content      string
	Ip           string
	Useragent    string
	Created      int64
}

func (this *Logs) TableName() string {
	return models.TableName("admin_log")
}
func init() {
	orm.RegisterModel(new(Logs))
}

//登录日志
func AddLogs(upd Logs) error {
	o := orm.NewOrm()
	o.Using("default")
	log := new(Logs)

	log.Userid = upd.Userid
	log.Username = upd.Username
	log.Url = upd.Url
	log.Title = upd.Title
	log.Content = upd.Content
	log.Ip = upd.Ip
	log.Useragent =upd.Useragent
	log.Created = time.Now().Unix()
	_, err := o.Insert(log)
	return err
}

func ListLogs(condArr map[string]string, page int, offset int) (num int64, err error, log []Logs) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("admin_log"))
	cond := orm.NewCondition()

	if condArr["username"] != "" {
		cond = cond.And("username__icontains", condArr["username"])
	}
	if condArr["ip"] != "" {
		cond = cond.And("ip__icontains", condArr["ip"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset

	qs = qs.OrderBy("-id")
	nums, errs := qs.Limit(offset, start).All(&log)
	return nums, errs, log
}

//统计数量
func CountLogs(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("admin_log"))
	cond := orm.NewCondition()

	if condArr["username"] != "" {
		cond = cond.And("username", condArr["username"])
	}
	if condArr["ip"] != "" {
		cond = cond.And("ip", condArr["ip"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}


func DeleteLogs(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("admin_log") + " WHERE id IN(" + ids + ")").Exec()
	return err
}
