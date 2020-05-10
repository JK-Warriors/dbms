package controllers

import (
	//"dbms/initial"

	. "dbms/models/messages"
	. "dbms/models/roles"
	"fmt"

	//"dbms/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	//UserInfo string
	UserId     int64
	Username   string
	UserAvatar string
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
		tmp := strings.Split((this.GetSession("userLogin")).(string), "||")

		userid, _ := strconv.Atoi(tmp[0])
		longid := int64(userid)
		this.Data["LoginUserid"] = longid
		this.Data["LoginUsername"] = tmp[1]
		this.Data["LoginAvatar"] = tmp[2]

		this.UserId = longid
		this.Username = tmp[1]
		this.UserAvatar = tmp[2]

		//this.Data["PermissionModel"] = this.GetSession("userPermissionModel")
		//this.Data["PermissionModelc"] = this.GetSession("userPermissionModelc")

		//消息
		msgcondArr := make(map[string]string)
		msgcondArr["touserid"] = fmt.Sprintf("%d", longid)
		msgcondArr["view"] = "1"
		countTopMessage := CountMessages(msgcondArr)
		_, _, topMessages := ListMessages(msgcondArr, 1, 10)
		this.Data["topMessages"] = topMessages
		this.Data["countTopMessage"] = countTopMessage

		//fmt.Println(this.GetSession("userGroupid").(string))
		//左侧导航
		url := this.Ctx.Request.RequestURI
		n_url := ""
		url_first_part := ""
		url_sec_part := ""
		//utils.LogDebug(string(url[1:]))
		if strings.Index(url, "/") >= 0 {
			if strings.Index(url, "?") >= 0 {
				pos := strings.Index(url, "?")
				url = url[0:pos]
				//utils.LogDebug(url)
			}

			sec_pos := strings.Index(url[1:], "/")
			if sec_pos >= 0 {
				url_first_part = url[0 : sec_pos+1]
				//utils.LogDebug(url_first_part)
				n_url = url[sec_pos+2:]

				third_pos := strings.Index(n_url, "/")

				if third_pos >= 0 {
					url_sec_part = n_url[0:third_pos]
					//utils.LogDebug(url_sec_part)
				}
			}
		}
		if url_sec_part != "" {
			n_url = url_first_part + "/" + url_sec_part
		} else {
			n_url = url_first_part
		}

		//点首页的时候，只有一个 / 或者为空时
		if n_url == "/" || n_url == "" {
			n_url = "---"
		}

		this.Data["url_first_part"] = url_first_part
		//_, _, leftNav := ListRoleUserPermission(this.GetSession("userRoleid").(string))
		_, _, leftNav := GetLeftNav(this.GetSession("userRoleid").(string), n_url)
		this.Data["leftNav"] = leftNav

	}
	this.Data["IsLogin"] = this.IsLogin
}
