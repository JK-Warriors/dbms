package disaster_config

import (
	"fmt"
	"opms/controllers"
	. "opms/models/business"
	. "opms/models/dbconfig"
	. "opms/models/disaster_config"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//业务系统管理
type ManageDisasterController struct {
	controllers.BaseController
}

func (this *ManageDisasterController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-disaster-manage") {
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

	host := this.GetString("host")
	condArr := make(map[string]string)
	condArr["host"] = host

	countDr := CountDisasterConfig(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countDr)
	_, _, drconf := ListDisasterConfig(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["drconf"] = drconf
	this.Data["countDr"] = countDr

	this.TplName = "disaster_config/disaster-index.tpl"
}

//修改业务系统
type EditDisasterController struct {
	controllers.BaseController
}

func (this *EditDisasterController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-disaster-manage") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	bsconf, err := GetBusiness(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["bsconf"] = bsconf

	drconf, err := GetDisasterConfig(id)
	this.Data["drconf"] = drconf

	pridbconf := ListPrimaryDBconfig(id)
	this.Data["pridbconf"] = pridbconf

	stadbconf := ListStandbyDBconfig(id)
	this.Data["stadbconf"] = stadbconf

	this.Data["dest_list"] = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	this.TplName = "disaster_config/disaster-form.tpl"
}

func (this *EditDisasterController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-disaster-manage") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权编辑"}
		this.ServeJSON()
		return
	}
	idstr := this.GetString("bs_id")
	//utils.LogDebug(idstr)
	bs_id, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_id_p")
	db_id_p, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_dest_p")
	db_dest_p, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_id_s")
	db_id_s, _ := strconv.Atoi(idstr)

	idstr = this.GetString("db_dest_s")
	db_dest_s, _ := strconv.Atoi(idstr)

	idstr = this.GetString("fb_retention")
	fb_retention, err := strconv.Atoi(idstr)
	if err != nil {
		fb_retention = 0
	}

	var is_shift int
	idstr = this.GetString("is_shift")
	if idstr == "on" {
		is_shift = 1
	} else {
		is_shift = 0
	}

	shift_vips := this.GetString("shift_vips")
	network_p := this.GetString("network_p")
	network_s := this.GetString("network_s")

	var disasterconf DisasterConfig

	disasterconf.Bs_Id = bs_id
	disasterconf.Db_Id_P = db_id_p
	disasterconf.Db_Dest_P = db_dest_p
	disasterconf.Db_Id_S = db_id_s
	disasterconf.Db_Dest_S = db_dest_s
	disasterconf.Fb_Retention = fb_retention
	disasterconf.Is_Shift = is_shift
	disasterconf.Shift_Vips = shift_vips
	disasterconf.Network_P = network_p
	disasterconf.Network_S = network_s

	ldc, err := GetDisasterConfig(bs_id)
	//utils.LogDebug(ldc.Db_Id_P)

	if ldc.Db_Id_P > 0 {
		err = UpdateDisasterConfig(bs_id, disasterconf)
	} else {
		err = AddDisasterConfig(disasterconf)
	}

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "修改容灾配置成功", "id": fmt.Sprintf("%d", bs_id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "修改容灾配置失败"}
	}
	this.ServeJSON()
}
