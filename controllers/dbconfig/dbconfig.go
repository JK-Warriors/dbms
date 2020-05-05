package dbconfig

import (
	"fmt"
	"opms/controllers"
	. "opms/models/business"
	. "opms/models/dbconfig"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//库管理
type ManageDBConfigController struct {
	controllers.BaseController
}

func (this *ManageDBConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-manage") {
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

	this.TplName = "dbconfig/dbconfig-index.tpl"
}

//添加数据库配置信息
type AddDBConfigController struct {
	controllers.BaseController
}

func (this *AddDBConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-add") {
		this.Abort("401")
	}

	var dbconf Dbconfigs
	this.Data["dbconf"] = dbconf

	bsconf := ListAllBusiness()
	this.Data["bsconf"] = bsconf

	this.TplName = "dbconfig/dbconfig-form.tpl"
}

func (this *AddDBConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权新增"}
		this.ServeJSON()
		return
	}

	bs_id, _ := this.GetInt("bs_id")

	db_type, _ := this.GetInt("db_type")
	if db_type <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择数据库类型"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
		this.ServeJSON()
		return
	}

	port, _ := this.GetInt("port")
	if port <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写端口"}
		this.ServeJSON()
		return
	}

	username := this.GetString("username")
	if "" == username {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
		return
	}

	password := this.GetString("password")
	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
		return
	}

	role, _ := this.GetInt("role")
	if role <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择角色"}
		this.ServeJSON()
		return
	}

	var dbconf Dbconfigs

	dbconf.Dbtype = db_type
	dbconf.Host = host
	dbconf.Port = port
	dbconf.Alias = this.GetString("alias")
	dbconf.InstanceName = this.GetString("instance_name")
	dbconf.Dbname = this.GetString("db_name")
	dbconf.Username = username
	dbconf.Password = password
	dbconf.Bs_Id = bs_id
	dbconf.Role = role

	err := AddDBconfig(dbconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "数据库配置信息添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "数据库配置信息添加失败"}
	}
	this.ServeJSON()
}

//修改数据库配置信息
type EditDBConfigController struct {
	controllers.BaseController
}

func (this *EditDBConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	dbconf, err := GetDBconfig(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["dbconf"] = dbconf

	bsconf := ListAllBusiness()
	this.Data["bsconf"] = bsconf

	this.TplName = "dbconfig/dbconfig-form.tpl"
}

func (this *EditDBConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权编辑"}
		this.ServeJSON()
		return
	}

	id, _ := this.GetInt("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户参数出错"}
		this.ServeJSON()
		return
	}

	db_type, _ := this.GetInt("db_type")
	if db_type <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择数据库类型"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
		this.ServeJSON()
		return
	}

	port, _ := this.GetInt("port")
	if port <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写端口"}
		this.ServeJSON()
		return
	}

	username := this.GetString("username")
	if "" == username {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		this.ServeJSON()
		return
	}

	password := this.GetString("password")
	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		this.ServeJSON()
		return
	}

	bs_id, _ := this.GetInt("bs_id")

	role, _ := this.GetInt("role")
	if role <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择角色"}
		this.ServeJSON()
		return
	}

	var dbconf Dbconfigs

	dbconf.Dbtype = db_type
	dbconf.Host = host
	dbconf.Port = port
	dbconf.Alias = this.GetString("alias")
	dbconf.InstanceName = this.GetString("instance_name")
	dbconf.Dbname = this.GetString("db_name")
	dbconf.Username = username
	dbconf.Password = password
	dbconf.Bs_Id = bs_id
	dbconf.Role = role

	err := UpdateDBconfig(id, dbconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "数据库配置信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "数据库配置信息修改失败"}
	}
	this.ServeJSON()
}

//数据库配置状态更改异步操作
type AjaxStatusDBConfigController struct {
	controllers.BaseController
}

func (this *AjaxStatusDBConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	id, _ := this.GetInt("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择数据库"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status >= 3 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeDBconfigStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "数据库状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "数据库状态更改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteDBConfigController struct {
	controllers.BaseController
}

func (this *AjaxDeleteDBConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权删除"}
		this.ServeJSON()
		return
	}
	ids := this.GetString("ids")
	if "" == ids {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择要删除的记录"}
		this.ServeJSON()
		return
	}

	err := DeleteDBconfig(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
