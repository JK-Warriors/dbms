package group

import (
	"dbms/controllers"
	. "dbms/models/group"
	"dbms/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//分组管理
type ManageGroupController struct {
	controllers.BaseController
}

func (this *ManageGroupController) Get() {
	utils.LogDebug("config-group-manage")
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-group-manage") {
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

	countGp := CountGroup(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countGp)
	_, _, gpconf := ListGroup(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["gpconf"] = gpconf
	this.Data["countGp"] = countGp

	this.TplName = "group/group-index.tpl"
}

//添加分组
type AddGroupController struct {
	controllers.BaseController
}

func (this *AddGroupController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-group-add") {
		this.Abort("401")
	}

	var gpconf Group
	this.Data["gpconf"] = gpconf

	this.TplName = "group/group-form.tpl"
}

func (this *AddGroupController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-group-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权新增"}
		this.ServeJSON()
		return
	}

	group_name := this.GetString("group_name")
	if "" == group_name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写分组名称"}
		this.ServeJSON()
		return
	}

	var gpconf Group

	gpconf.GroupName = group_name
	is_exist := CheckNameExist(group_name)
	if 1 == is_exist {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "该组名已经存在，请换一个名称"}
	} else {
		err := AddGroup(gpconf)

		if err == nil {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "分组添加成功"}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "分组添加失败"}
		}
	}

	this.ServeJSON()
}

//修改分组
type EditGroupController struct {
	controllers.BaseController
}

func (this *EditGroupController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-group-edit") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	gpconf, err := GetGroup(id)
	if err != nil {
		this.Abort("404")
	}
	this.Data["gpconf"] = gpconf
}

func (this *EditGroupController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-group-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权编辑"}
		this.ServeJSON()
		return
	}
	idstr := this.GetString("group_id")
	//utils.LogDebug(idstr)
	id, _ := strconv.Atoi(idstr)

	group_name := this.GetString("group_name")
	if "" == group_name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写分组名称"}
		this.ServeJSON()
		return
	}

	var gpconf Group

	gpconf.GroupName = group_name

	err := UpdateGroup(id, gpconf)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "分组信息修改成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "分组信息修改失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteGroupController struct {
	controllers.BaseController
}

func (this *AjaxDeleteGroupController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "config-group-delete") {
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

	err := DeleteGroup(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
