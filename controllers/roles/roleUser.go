package roles

import (
	"dbms/controllers"
	"dbms/models/roles"
	"dbms/utils"
	"fmt"
	"strconv"
	"strings"
	//"time"
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/utils/pagination"
)

//角色成员管理
type ManageRoleUserController struct {
	controllers.BaseController
}

func (this *ManageRoleUserController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-user") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")

	roleid, _ := strconv.Atoi(idstr)

	role, _ := roles.GetRole(int64(roleid))
	this.Data["role"] = role

	_, _, users := roles.ListRolesUserAndName(int64(roleid))
	fmt.Println(users)
	this.Data["users"] = users

	this.TplName = "roles/user.tpl"
}

type FormRoleUserController struct {
	controllers.BaseController
}

func (this *FormRoleUserController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-user-add") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")
	//fmt.Println("hello" + idstr)

	if "" != idstr {
		id, _ := strconv.Atoi(idstr)
		role, _ := roles.GetRole(int64(id))
		this.Data["role"] = role
	}
	this.TplName = "roles/user-form.tpl"
}

func (this *FormRoleUserController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-user-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	roleid, _ := this.GetInt64("roleid")
	if roleid <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择组"}
		this.ServeJSON()
		return
	}
	userid, _ := this.GetInt64("userid")
	if userid <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写成员"}
		this.ServeJSON()
		return
	}

	var roleUser roles.RolesUser
	var err error
	roleUser.Id = utils.SnowFlakeId()
	roleUser.Roleid = roleid
	roleUser.Userid = userid
	err = roles.AddRolesUser(roleUser)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "操作成功", "id": fmt.Sprintf("%d", roleid)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteRoleUserController struct {
	controllers.BaseController
}

func (this *AjaxDeleteRoleUserController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-user-delete") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	id, _ := this.GetInt64("id")
	if id <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择要删除的选项"}
		this.ServeJSON()
		return
	}

	err := roles.DeleteRolesUser(id)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
