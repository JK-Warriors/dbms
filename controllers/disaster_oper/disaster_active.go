package disaster_oper

import (
	"opms/controllers"
	. "opms/models/business"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//业务系统管理
type ManageDisasterActiveController struct {
	controllers.BaseController
}

func (this *ManageDisasterActiveController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "business-manage") {
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

	this.TplName = "business/business-index.tpl"
}
