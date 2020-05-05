package disaster_oper

import (
	"opms/controllers"
	. "opms/models/business"
	. "opms/models/users"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//业务切换管理
type ManageDisasterSwitchController struct {
	controllers.BaseController
}

func (this *ManageDisasterSwitchController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch") {
		this.Abort("401")
	}

	page, err := this.GetInt("p")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	search_name := this.GetString("search_name")
	condArr := make(map[string]string)
	condArr["search_name"] = search_name

	countBs := CountBusiness(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countBs)
	_, _, bsconf := ListBusiness(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["bsconf"] = bsconf
	this.Data["countBs"] = countBs

	this.TplName = "disaster_oper/operation-index.tpl"
}

//业务查看
type ViewDisasterSwitchController struct {
	controllers.BaseController
}

func (this *ViewDisasterSwitchController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
		this.Abort("401")
	}

	user, _ := GetUser(1)
	this.Data["user"] = user

	this.TplName = "disaster_oper/operation-detail.tpl"
}

type AjaxDisasterSwitchoverController struct {
	controllers.BaseController
}

func (this *AjaxDisasterSwitchoverController) Post() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
	// 	this.Abort("401")
	// }

	user, _ := GetUser(1)
	this.Data["user"] = user

}

type AjaxDisasterFailoverController struct {
	controllers.BaseController
}

func (this *AjaxDisasterFailoverController) Post() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
	// 	this.Abort("401")
	// }

	user, _ := GetUser(1)
	this.Data["user"] = user

}

type AjaxDisasterProcessController struct {
	controllers.BaseController
}

func (this *AjaxDisasterProcessController) Post() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "oper-switch-view") {
	// 	this.Abort("401")
	// }

	curr_time := time.Now().Format("2006-01-02 15:04:05")
	this.Data["json"] = map[string]interface{}{"on_process": 1, "op_type": "SWITCHOVER", "op_reason": "null", "process_time": curr_time, "process_desc": "testdesc"}
	this.ServeJSON()

}
