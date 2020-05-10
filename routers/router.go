package routers

import (
	"dbms/controllers/dbconfig"
	"dbms/controllers/dbmanage"
	"dbms/controllers/demo"
	"dbms/controllers/group"
	"dbms/controllers/logs"
	"dbms/controllers/messages"
	"dbms/controllers/roles"
	"dbms/controllers/users"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &users.MainController{})

	//用户
	beego.Router("/login", &users.LoginUserController{})
	beego.Router("/logout", &users.LogoutUserController{})
	beego.Router("/user/manage", &users.ManageUserController{})
	beego.Router("/user/ajax/status", &users.AjaxStatusUserController{})
	beego.Router("/user/add", &users.AddUserController{})
	beego.Router("/user/edit/:id", &users.EditUserController{})
	beego.Router("/user/ajax/delete", &users.AjaxDeleteUserController{})
	beego.Router("/user/avatar", &users.AvatarUserController{})
	beego.Router("/user/ajax/search", &users.AjaxSearchUserController{}) //搜索用户名匹配
	beego.Router("/user/show/:id", &users.ShowUserController{})
	beego.Router("/user/profile", &users.EditUserProfileController{})
	beego.Router("/user/password", &users.EditUserPasswordController{})
	beego.Router("/user/ajax/reset_passwd", &users.AjaxResetPasswordController{})

	beego.Router("/my/manage", &users.ShowUserController{})

	//消息
	beego.Router("/message/manage", &messages.ManageMessageController{})
	beego.Router("/message/ajax/delete", &messages.AjaxDeleteMessageController{})
	beego.Router("/message/ajax/status", &messages.AjaxStatusMessageController{})

	//角色
	beego.Router("/role/manage", &roles.ManageRoleController{})
	beego.Router("/role/ajax/delete", &roles.AjaxDeleteRoleController{})
	beego.Router("/role/add", &roles.FormRoleController{})
	beego.Router("/role/edit/:id", &roles.FormRoleController{})
	//角色成员
	beego.Router("/role/user/:id", &roles.ManageRoleUserController{})
	beego.Router("/role/user/add/:id", &roles.FormRoleUserController{})
	beego.Router("/role/user/ajax/delete", &roles.AjaxDeleteRoleUserController{})

	//角色权限
	beego.Router("/role/permission/:id", &roles.ManageRolePermissionController{})
	beego.Router("/role/permission/ajax/delete", &roles.AjaxDeleteRolePermissionController{})

	//权限
	beego.Router("/permission/manage", &roles.ManagePermissionController{})
	beego.Router("/permission/ajax/delete", &roles.AjaxDeletePermissionController{})
	beego.Router("/permission/add", &roles.FormPermissionController{})
	beego.Router("/permission/edit/:id", &roles.FormPermissionController{})

	//日志
	beego.Router("/log/manage", &logs.ManageLogController{})
	beego.Router("/log/ajax/delete", &logs.AjaxDeleteLogController{})

	//分组管理
	beego.Router("/config/group/manage", &group.ManageGroupController{})
	beego.Router("/config/group/add", &group.AddGroupController{})
	beego.Router("/config/group/edit", &group.EditGroupController{})
	beego.Router("/config/group/ajax/delete", &group.AjaxDeleteGroupController{})

	//数据库配置
	beego.Router("/config/db/manage", &dbconfig.ManageDBConfigController{})
	beego.Router("/config/db/add", &dbconfig.AddDBConfigController{})
	beego.Router("/config/db/edit/:id", &dbconfig.EditDBConfigController{})
	beego.Router("/config/db/ajax/status", &dbconfig.AjaxStatusDBConfigController{})
	beego.Router("/config/db/ajax/delete", &dbconfig.AjaxDeleteDBConfigController{})
	beego.Router("/config/db/ajax/connect", &dbconfig.AjaxConnectDBConfigController{})

	//数据库管理
	beego.Router("/maintenance/healthcheck/manage", &dbmanage.ManageHealthCheckController{})

	//操作

	//UI demo
	beego.Router("/demo/index", &demo.DemoController{})
	beego.Router("/demo/form", &demo.FormController{})
	beego.Router("/demo/base", &demo.BaseController{})
	beego.Router("/demo/dashboard", &demo.DashboardController{})
	beego.Router("/demo/dgscreen", &demo.DgscreenController{})
}
