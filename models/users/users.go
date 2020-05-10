package users

import (
	crypt_rand "crypto/rand"
	"dbms/models"
	"dbms/utils"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       int64         `orm:"pk;column(userid);"`
	Profile  *UsersProfile `orm:"rel(one);"`
	Username string
	Password string
	Avatar   string
	Status   int
}

type UsersProfile struct {
	Id          int64 `orm:"pk;column(userid);"`
	Realname    string
	Sex         int
	Birth       string
	Email       string
	Webchat     string
	Qq          string
	Phone       string
	Tel         string
	Address     string
	Emercontact string
	Emerphone   string
	Lognum      int
	Ip          string
	Lasted      int64
}

func (this *Users) TableName() string {
	return models.TableName("users")
}

func init() {
	//orm.RegisterModel(new(Users), new(UsersProfile))
	orm.RegisterModel(new(Users))
	//orm.RegisterModelWithPrefix("pms_", new(Users))
	orm.RegisterModelWithPrefix("pms_", new(UsersProfile))
}

//登录
func LoginUser(username, password string) (err error, user Users) {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("users"))
	cond := orm.NewCondition()

	cond = cond.And("username", username)
	pwdmd5 := utils.Md5(password)
	cond = cond.And("password", pwdmd5)
	//cond = cond.And("status", 1)

	qs = qs.SetCond(cond)
	var users Users
	err = qs.Limit(1).One(&users, "userid", "username", "avatar", "status")
	fmt.Println(err)
	if err == nil {
		if 1 == users.Status {
			o.Raw("UPDATE pms_users_profile SET lasted = ?,lognum=lognum+? WHERE userid = ?", time.Now().Unix(), 1, users.Id).Exec()
		}
	}
	return err, users
}

//得到用户信息
func GetUser(id int64) (Users, error) {
	var user Users
	var err error
	o := orm.NewOrm()

	user = Users{Id: id}
	err = o.Read(&user)

	if err == orm.ErrNoRows {
		return user, nil
	}
	return user, err
}

func GetRealname(id int64) string {
	var err error
	var realname string

	err = utils.GetCache("GetRealname.id."+fmt.Sprintf("%d", id), &realname)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var user UsersProfile
		o := orm.NewOrm()
		o.QueryTable(models.TableName("users_profile")).Filter("userid", id).One(&user, "realname")
		realname = user.Realname
		utils.SetCache("GetRealname.id."+fmt.Sprintf("%d", id), realname, cache_expire)
	}
	return realname
}

func GetUserEmail(id int64) string {
	var err error
	var email string

	err = utils.GetCache("GetUserEmail.id."+fmt.Sprintf("%d", id), &email)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var user UsersProfile
		o := orm.NewOrm()
		o.QueryTable(models.TableName("users_profile")).Filter("userid", id).One(&user, "email")
		email = user.Email
		utils.SetCache("GetUserEmail.id."+fmt.Sprintf("%d", id), email, cache_expire)
	}
	return email
}

func GetAvatarUserid(id int64) string {
	var err error
	var avatar string

	err = utils.GetCache("GetAvatarUserid.id."+fmt.Sprintf("%d", id), &avatar)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var user Users
		o := orm.NewOrm()
		o.QueryTable(models.TableName("users")).Filter("userid", id).One(&user, "avatar")
		avatar = user.Avatar
		utils.SetCache("GetAvatarUserid.id."+fmt.Sprintf("%d", id), avatar, cache_expire)
	}
	if "" == avatar {
		return fmt.Sprintf("/static/img/avatar/%d.jpg", rand.Intn(5))
	}
	return avatar
}

//得到用户详情信息
func GetProfile(id int64) (UsersProfile, error) {
	var pro UsersProfile
	var err error
	o := orm.NewOrm()

	pro = UsersProfile{Id: id}
	err = o.Read(&pro)

	if err == orm.ErrNoRows {
		return pro, nil
	}
	return pro, err
}

//修改个人信息
func UpdateProfile(id int64, updPro UsersProfile) error {
	var pro UsersProfile
	o := orm.NewOrm()
	pro = UsersProfile{Id: id}

	pro.Realname = updPro.Realname
	pro.Sex = updPro.Sex
	pro.Birth = updPro.Birth
	pro.Email = updPro.Email
	pro.Webchat = updPro.Webchat
	pro.Qq = updPro.Qq
	pro.Phone = updPro.Phone
	pro.Tel = updPro.Tel
	pro.Address = updPro.Address
	pro.Emercontact = updPro.Emercontact
	pro.Emerphone = updPro.Emerphone

	_, err := o.Update(&pro, "realname", "sex", "birth", "email", "webchat", "qq", "phone", "tel", "address", "emercontact", "emerphone")
	return err

}

//修改用户
func UpdateUser(id int64, updUser Users) error {
	var user Users
	o := orm.NewOrm()
	user = Users{Id: id}

	user.Username = updUser.Username
	if updUser.Password != "" {
		user.Password = utils.Md5(updUser.Password)
		_, err := o.Update(&user, "username", "password")
		return err
	} else {
		_, err := o.Update(&user, "username")
		return err
	}
}

//修改密码
func UpdatePassword(id int64, oldPawd string, newPwd string) error {
	o := orm.NewOrm()

	user := Users{Id: id}
	err := o.Read(&user)
	if nil != err {
		return err
	} else {
		if user.Password == utils.Md5(oldPawd) {
			user.Password = utils.Md5(newPwd)
			_, err := o.Update(&user)
			return err
		} else {
			return fmt.Errorf("验证出错")
		}
	}
}

//添加个人信息
func AddProfile(updPro UsersProfile) error {
	o := orm.NewOrm()
	pro := new(UsersProfile)

	pro.Id = updPro.Id
	pro.Realname = updPro.Realname
	pro.Sex = updPro.Sex
	pro.Birth = updPro.Birth
	pro.Email = updPro.Email
	pro.Webchat = updPro.Webchat
	pro.Qq = updPro.Qq
	pro.Phone = updPro.Phone
	pro.Tel = updPro.Tel
	pro.Address = updPro.Address
	pro.Emercontact = updPro.Emercontact
	pro.Emerphone = updPro.Emerphone
	pro.Lognum = 1
	pro.Ip = updPro.Ip
	pro.Lasted = time.Now().Unix()
	_, err := o.Insert(pro)
	return err
}

//添加用户
func AddUserProfile(updUser Users, updPro UsersProfile) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)

	pro := new(UsersProfile)

	//pro.Id = updPro.Id
	pro.Realname = updPro.Realname
	pro.Sex = updPro.Sex
	pro.Birth = updPro.Birth
	pro.Email = updPro.Email
	pro.Tel = updPro.Tel
	pro.Webchat = updPro.Webchat
	pro.Qq = updPro.Qq
	pro.Phone = updPro.Phone
	pro.Address = updPro.Address
	pro.Emercontact = updPro.Emercontact
	pro.Emerphone = updPro.Emerphone
	pro.Lognum = 1
	pro.Ip = updPro.Ip
	pro.Lasted = time.Now().Unix()
	id, err := o.Insert(pro)

	pro.Id = id
	user.Id = id
	user.Profile = pro
	user.Username = updUser.Username
	user.Password = utils.Md5(updUser.Password)
	user.Avatar = utils.GetAvatar("")
	user.Status = 1
	_, err = o.Insert(user)

	return id, err
}

//用户列表
func ListUser(condArr map[string]string, page int, offset int) (num int64, err error, user []Users) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("users"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("username__icontains", condArr["keywords"]).Or("profile__realname__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()

	var users []Users
	num, err1 := qs.Limit(offset, start).All(&users)
	return num, err1, users
}

//统计数量
func CountUser(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("users"))
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("username__icontains", condArr["keywords"]).Or("profile__realname__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}

//更改用户状态
func ChangeUserStatus(id int64, status int) error {
	o := orm.NewOrm()

	user := Users{Id: id}
	err := o.Read(&user, "userid")
	if nil != err {
		return err
	} else {
		user.Status = status
		_, err := o.Update(&user)
		return err
	}
}

//更改用户头像
func ChangeUserAvatar(id int64, avatar string) error {
	o := orm.NewOrm()

	user := Users{Id: id}
	err := o.Read(&user, "userid")
	if nil != err {
		return err
	} else {
		user.Avatar = avatar
		_, err := o.Update(&user)
		return err
	}
}

type UsersFind struct {
	Userid   int64
	Realname string
	Avatar   string
	Position string
}

//显示所有用户
func ListUserFind() (num int64, err error, user []UsersFind) {
	var users []UsersFind
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("upr.userid", "upr.realname", "p.name AS position", "u.avatar").From("pms_users AS u").
		LeftJoin("pms_users_profile AS upr").On("upr.userid = u.userid").
		LeftJoin("pms_positions AS p").On("p.positionid = upr.positionid").
		Where("u.status=1").
		OrderBy("p.name").
		Desc()
	sql := qb.String()
	o := orm.NewOrm()
	nums, err := o.Raw(sql).QueryRows(&users)
	return nums, err, users
}

//删除用户
func DeleteUsers(ids string) error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM " + models.TableName("users") + " WHERE userid IN(" + ids + ")").Exec()
	_, err = o.Raw("DELETE FROM " + models.TableName("role_user") + " WHERE user_id IN(" + ids + ")").Exec()
	_, err = o.Raw("DELETE FROM " + models.TableName("users_profile") + " WHERE userid IN(" + ids + ")").Exec()
	return err
}

//重置用户密码
func ResetUserPassword(id string) (error, string) {
	var newPasswd string
	b := make([]byte, 8)
	//ReadFull从rand.Reader精确地读取len(b)字节数据填充进b
	//rand.Reader是一个全局、共享的密码用强随机数生成器
	if _, err := io.ReadFull(crypt_rand.Reader, b); err != nil {
		newPasswd = "123456"
	} else {
		//fmt.Println(b) //[62 186 123 16 209 19 130 218]
		newPasswd = base64.URLEncoding.EncodeToString(b)
	}

	pwdmd5 := utils.Md5(newPasswd)

	o := orm.NewOrm()
	_, err := o.Raw("UPDATE pms_users SET password = ? WHERE userid = ?", pwdmd5, id).Exec()
	return err, newPasswd
}

type UsersPermissionsAll struct {
	Permission string
	Roleid     string
}

func GetPermissionsAll(id int64) (UsersPermissionsAll, error) {
	var pers []UsersPermissionsAll
	var err error

	sql := `select p.ename as permission, ru.role_id as roleid 
			from pms_role_user ru, pms_role_permission rp, pms_permissions p
			where rp.role_id = ru.role_id 
			and p.id = rp.permission_id
			and ru.user_id = ?`
	//utils.LogDebug(sql)
	o := orm.NewOrm()
	_, err = o.Raw(sql, id).QueryRows(&pers)

	var permissionstring = ""
	var roleidstring = ""
	for _, v := range pers {
		permissionstring += v.Permission + ","
		roleidstring += v.Roleid + ","
	}
	roleidstrings := strings.Split(strings.Trim(roleidstring, ","), ",")
	roleMap := utils.RemoveDuplicatesAndEmpty(roleidstrings)
	roleidstring = ""

	for _, v := range roleMap {
		roleidstring += v + ","
	}
	var per UsersPermissionsAll
	per.Permission = strings.Trim(permissionstring, ",")
	per.Roleid = strings.Trim(roleidstring, ",")

	return per, err
}
