package demo

import (
	//"fmt"
	//"dbms/utils"
	//"strconv"
	//"strings"
	//"time"
	"dbms/controllers"
)

//消息管理

type DemoController struct {
	controllers.BaseController
}

func (this *DemoController) Get() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-manage") {
	// 	this.Abort("401")
	// }

	this.TplName = "demo/index.tpl"
}

type FormController struct {
	controllers.BaseController
}

func (this *FormController) Get() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-manage") {
	// 	this.Abort("401")
	// }

	this.TplName = "demo/form.tpl"
}

type BaseController struct {
	controllers.BaseController
}

func (this *BaseController) Get() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-manage") {
	// 	this.Abort("401")
	// }

	this.TplName = "demo/base.tpl"
}

type DashboardController struct {
	controllers.BaseController
}

func (this *DashboardController) Get() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-manage") {
	// 	this.Abort("401")
	// }

	this.TplName = "demo/dashboard.tpl"
}

type DgscreenController struct {
	controllers.BaseController
}

func (this *DgscreenController) Get() {
	//权限检测
	// if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-manage") {
	// 	this.Abort("401")
	// }

	this.TplName = "demo/dgscreen.tpl"
}

// type BaseController struct {
// 	controllers.BaseController
// }
// func (this *BaseController) Get() {
// 	//权限检测
// 	// if !strings.Contains(this.GetSession("userPermission").(string), "dbconfig-manage") {
// 	// 	this.Abort("401")
// 	// }

// 	this.TplName = "demo/base.tpl"
// }
