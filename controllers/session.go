package controllers

import (
	"fmt"
  	"encoding/json"
    	"github.com/astaxie/beego"
    	"github.com/astaxie/beego/orm"
    	"zufang/models"
)

type SessionController struct {
    beego.Controller
}


//TODO 获取Session
func (c *SessionController) GetSessionData()  {
	fmt.Println("connect success")
	resp :=make(map[string]interface{})
	defer c.RetData(resp)
	user := models.User{}

	resp["errno"] = 4001
	resp["errmsg"] = "数据获取失败"
	name := c.GetSession("name")
	if name!=nil {
		user.Name = name.(string)
		resp["errno"] = 0
		resp["mrrmsg"] = "ok"
		resp["data"] = user
	}
}

//TODO 删除对应的Session
func (c *SessionController) DeleteSessionData()  {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	c.DelSession("name")
	resp["errno"]=0
	resp["errmsg"]="ok"
}
// TODO 登录
func (c *SessionController) Login()  {

	resp := make(map[string]interface{})
	defer c.RetData(resp)

	//得到用户信息获取前端传递过来的json数据
	json.Unmarshal(c.Ctx.Input.RequestBody,&resp)
	fmt.Println(&resp)
	//判断是否合法
	if resp["mobile"] == nil || resp["password"] ==nil{
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
		return
	}
	//与数据库匹配账号密码是否正确
	o := orm.NewOrm()
	user := models.User{Name:resp["mobile"].(string)}
	moble := resp["mobile"].(string)
	qs := o.QueryTable("user")
	err := qs.Filter("mobile",moble).One(&user)
	if err !=nil {
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
		fmt.Println("2222name=",resp["mobile"],"========password====",resp["password"])

		return
	}
	if user.Password_hash != resp["password"] {
		resp["errno"]=models.RECODE_DATAERR
		resp["errmsg"]=models.RecodeText(models.RECODE_DATAERR)
		fmt.Println("3333name=",resp["mobile"],"========password====",resp["password"])
		return
	}
	//添加Session
	c.SetSession("name",resp["mobile"])
	c.SetSession("mobile",resp["mobile"])
	c.SetSession("user_id",user.Id)
	//返回json数据给前端
	resp["errno"]=models.RECODE_OK
	resp["errmsg"]=models.RecodeText(models.RECODE_OK)


}
func (c *SessionController) RetData(resp map[string]interface{}) {
	c.Data["json"] =resp
	c.ServeJSON()
}