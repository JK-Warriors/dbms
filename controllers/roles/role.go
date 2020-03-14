package roles

import (
	"fmt"
	"opms/controllers"
	"opms/utils"
	"opms/models/roles"
	"strconv"
	"strings"
	//"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//角色管理
type ManageRoleController struct {
	controllers.BaseController
}

func (this *ManageRoleController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-manage") {
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
	keywords := this.GetString("keywords")

	condArr := make(map[string]string)
	condArr["keywords"] = keywords

	countRole := roles.CountRole(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countRole)
	_, _, roles := roles.ListRole(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["roles"] = roles
	this.Data["countRole"] = countRole

	this.TplName = "roles/index.tpl"
}

type FormRoleController struct {
	controllers.BaseController
}

func (this *FormRoleController) Get() {

	idstr := this.Ctx.Input.Param(":id")
	if "" != idstr {
		//权限检测
		if !strings.Contains(this.GetSession("userPermission").(string), "role-edit") {
			this.Abort("401")
		}
		id, _ := strconv.Atoi(idstr)
		role, _ := roles.GetRole(int64(id))
		this.Data["role"] = role
	} else {
		//权限检测
		if !strings.Contains(this.GetSession("userPermission").(string), "role-add") {
			this.Abort("401")
		}
	}

	this.TplName = "roles/form.tpl"
}

func (this *FormRoleController) Post() {
	//权限检测
	name := this.GetString("name")
	if "" == name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写角色名称"}
		this.ServeJSON()
		return
	}
	summary := this.GetString("summary")

	var role roles.Roles
	role.Name = name
	role.Summary = summary

	roleid, _ := this.GetInt64("id")
	var err error
	if roleid <= 0 {
		roleid = utils.SnowFlakeId()
		role.Id = roleid
		err = roles.AddRole(role)
	} else {
		err = roles.UpdateRole(roleid, role)
	}

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "操作成功", "id": fmt.Sprintf("%d", roleid)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteRoleController struct {
	controllers.BaseController
}

func (this *AjaxDeleteRoleController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	ids := this.GetString("ids")
	if "" == ids {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择要删除的选项"}
		this.ServeJSON()
		return
	}

	err := roles.DeleteRole(ids)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
