package logs

import (
	"fmt"
	"opms/controllers"
	. "opms/models/logs"
	//"opms/utils"
	//"strconv"
	"strings"
	//"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//消息管理
type ManageLogController struct {
	controllers.BaseController
}

func (this *ManageLogController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "log-manage") {
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
	username := this.GetString("username")
	ip := this.GetString("ip")

	condArr := make(map[string]string)
	condArr["userid"] = fmt.Sprintf("%d", this.BaseController.UserId)
	condArr["username"] = username
	condArr["ip"] = ip

	countLog := CountLogs(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countLog)
	_, _, syslogs := ListLogs(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["syslogs"] = syslogs
	this.Data["countLog"] = countLog

	//查看列表即更新查看状态标为已查看
	//ChangeMessagesStatusAll(this.BaseController.UserId)

	this.TplName = "logs/index.tpl"
}


type AjaxDeleteLogController struct {
	controllers.BaseController
}

func (this *AjaxDeleteLogController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "log-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权删除"}
		this.ServeJSON()
		return
	}
	ids := this.GetString("ids")
	if "" == ids {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择要删除的日志"}
		this.ServeJSON()
		return
	}

	err := DeleteLogs(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
