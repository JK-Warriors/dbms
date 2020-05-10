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
type ManageRolePermissionController struct {
	controllers.BaseController
}

func (this *ManageRolePermissionController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-permission") {
		this.Abort("401")
	}
	idstr := this.Ctx.Input.Param(":id")

	roleid, _ := strconv.Atoi(idstr)

	role, _ := roles.GetRole(int64(roleid))
	this.Data["role"] = role

	condArr := make(map[string]string)
	_, _, permissions := roles.ListPermission(condArr, 1, 1000)

	var pstring = ""
	var cstring = ""
	for _, value := range permissions {
		if value.Parentid == 1 {
			pstring += fmt.Sprintf("%d", value.Id) + "||" + fmt.Sprintf("%d", value.Parentid) + "||" + value.Name + "||" + value.Url + ","
		} else {
			cstring += fmt.Sprintf("%d", value.Id) + "||" + fmt.Sprintf("%d", value.Parentid) + "||" + value.Name + "||" + value.Url + ","
		}
	}
	this.Data["pstring"] = pstring
	this.Data["cstring"] = cstring
	//fmt.Println(pstring)
	//fmt.Println(cstring)

	rolespermissions := roles.ListRolePermission(int64(roleid))
	var rolestring = ""
	for _, v := range rolespermissions {
		rolestring += fmt.Sprintf("%d", v.Permissionid) + ","
	}

	this.Data["rolespermissions"] = rolestring
	//fmt.Println(rolestring)

	this.TplName = "roles/role-permission.tpl"
}

func (this *ManageRolePermissionController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-permission") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
	}
	roleid, _ := this.GetInt64("roleid")
	if roleid <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请选择组"}
		this.ServeJSON()
		return
	}

	//permission := make([]string, 0, 2)
	//this.Ctx.Input.Bind(&permission, "permission") //ul ==[str array]

	permission := this.GetString("permission")

	//fmt.Println("hello")
	//fmt.Println(permission)

	var rolePermission roles.RolePermission
	var err error

	rolePermission.Roleid = roleid

	//先删除,再添加
	roles.DeleteRolePermissionForRoleid(roleid)

	names := strings.Split(permission, ",")
	for _, v := range names {
		pid, _ := strconv.Atoi(v)
		rolePermission.Id = utils.SnowFlakeId()
		rolePermission.Permissionid = int64(pid)
		err = roles.AddRolePermission(rolePermission)
	}

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "操作成功", "id": fmt.Sprintf("%d", roleid)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
	}
	this.ServeJSON()
}

type AjaxDeleteRolePermissionController struct {
	controllers.BaseController
}

func (this *AjaxDeleteRolePermissionController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "role-permission") {
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

	err := roles.DeleteRolePermission(id)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}
