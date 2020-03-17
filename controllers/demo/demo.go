package demo

import (
	"fmt"
	"opms/controllers"
	. "opms/models/demo"
	//"opms/utils"
	//"strconv"
	"strings"
	//"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//消息管理
type demoController struct {
	
	controllers.BaseController
}

func (this *BaseController) Prepare() {

}


