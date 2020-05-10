package dbnode

import (
	"dbms/controllers"
	. "dbms/models/dbconfig"
	. "dbms/models/dbnode"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-oci8"
)

//节点管理
type ManageNodeConfigController struct {
	controllers.BaseController
}

func (this *ManageNodeConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-node-manage") {
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

	ostype := this.GetString("ostype")
	host := this.GetString("host")
	nodename := this.GetString("nodename")
	condArr := make(map[string]string)
	condArr["ostype"] = ostype
	condArr["host"] = host
	condArr["nodename"] = nodename

	countOs := CountNodeConfig(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countOs)
	_, _, osconf := ListNodeConfig(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["osconf"] = osconf
	this.Data["countOs"] = countOs

	this.TplName = "dbnode/dbnode-index.tpl"
}

//添加节点信息
type AddNodeConfigController struct {
	controllers.BaseController
}

func (this *AddNodeConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-node-add") {
		this.Abort("401")
	}

	var osconf Nodeconfig
	this.Data["osconf"] = osconf

	dbconf := ListAllDBconfig()
	this.Data["dbconf"] = dbconf

	this.TplName = "dbnode/dbnode-form.tpl"
}

func (this *AddNodeConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-node-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权新增"}
		this.ServeJSON()
		return
	}

	db_id, _ := this.GetInt("db_id")

	os_type, _ := this.GetInt("os_type")
	if os_type <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择主机类型"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
		this.ServeJSON()
		return
	}

	protocol := this.GetString("protocol")
	if "" == protocol {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择协议"}
		this.ServeJSON()
		return
	}

	port, _ := this.GetInt("port")
	if port <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写端口"}
		this.ServeJSON()
		return
	}

	nodename := this.GetString("nodename")
	if "" == nodename {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机名"}
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

	var conf Nodeconfig

	conf.Ostype = os_type
	conf.Host = host
	conf.Protocol = protocol
	conf.Port = port
	conf.Nodename = nodename
	conf.Username = username
	conf.Password = password
	conf.DbId = db_id

	err := AddNodeConfig(conf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "主机信息添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "主机信息添加失败"}
	}
	this.ServeJSON()
}

//修改节点信息
type EditNodeConfigController struct {
	controllers.BaseController
}

func (this *EditNodeConfigController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-node-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	osconf, err := GetNodeConfig(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["osconf"] = osconf

	dbconf := ListAllDBconfig()
	this.Data["dbconf"] = dbconf

	this.TplName = "dbnode/dbnode-form.tpl"
}

func (this *EditNodeConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-node-edit") {
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

	os_type, _ := this.GetInt("os_type")
	if os_type <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择主机类型"}
		this.ServeJSON()
		return
	}

	host := this.GetString("host")
	if "" == host {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机IP"}
		this.ServeJSON()
		return
	}

	protocol := this.GetString("protocol")
	if "" == protocol {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择协议"}
		this.ServeJSON()
		return
	}

	port, _ := this.GetInt("port")
	if port <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写端口"}
		this.ServeJSON()
		return
	}

	nodename := this.GetString("nodename")
	if "" == nodename {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写主机名"}
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

	db_id, _ := this.GetInt("db_id")

	var conf Nodeconfig

	conf.Ostype = os_type
	conf.Host = host
	conf.Protocol = protocol
	conf.Port = port
	conf.Nodename = nodename
	conf.Username = username
	conf.Password = password
	conf.DbId = db_id

	err := UpdateNodeConfig(id, conf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "节点配置信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "节点配置信息修改失败"}
	}
	this.ServeJSON()
}

//节点状态更改异步操作
type AjaxStatusNodeConfigController struct {
	controllers.BaseController
}

func (this *AjaxStatusNodeConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-node-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	id, _ := this.GetInt("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择节点"}
		this.ServeJSON()
		return
	}
	status, _ := this.GetInt("status")
	if status <= 0 || status >= 3 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择操作状态"}
		this.ServeJSON()
		return
	}

	err := ChangeNodeConfigStatus(id, status)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "节点状态更改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "节点状态更改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteNodeConfigController struct {
	controllers.BaseController
}

func (this *AjaxDeleteNodeConfigController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-node-delete") {
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

	err := DeleteNodeConfig(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
