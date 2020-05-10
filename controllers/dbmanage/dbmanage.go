package dbmanage

import (
	"dbms/controllers"
	. "dbms/models/dbconfig"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//库管理
type ManageHealthCheckController struct {
	controllers.BaseController
}

func (this *ManageHealthCheckController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "mt-healthcheck-manage") {
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

	dbtype := this.GetString("dbtype")
	host := this.GetString("host")
	alias := this.GetString("alias")
	condArr := make(map[string]string)
	condArr["dbtype"] = dbtype
	condArr["host"] = host
	condArr["alias"] = alias

	countDb := CountDBconfig(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDb)
	_, _, dbconf := ListDBconfig(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["dbconf"] = dbconf
	this.Data["countDb"] = countDb

	this.TplName = "dbmanage/dbmanage-index.tpl"
}
