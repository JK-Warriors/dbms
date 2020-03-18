package demo

import (
	//"fmt"
	//"opms/utils"
	//"strconv"
	//"strings"
	//"time"
	"opms/controllers"
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


