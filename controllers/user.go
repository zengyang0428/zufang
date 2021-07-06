package controllers

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"zufang/models"
	//"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserController) Reg() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	//获取前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	/*
		mobile:"111"
		password:"111"
		sms_code:"111"

		fmt.Println(`resp["mobile"] =`,resp["mobile"])
		fmt.Println(`resp["password"] =`,resp["password"])
		fmt.Println(`resp["sms_code"] =`,resp["sms_code"])
	*/

	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string)
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)

	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = 4002
		resp["errmsg"] = "注册失败"
		return
	}

	fmt.Println("reg success ,id = ", id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"

	this.SetSession("name", user.Name)

}
func (this *UserController)Postavatar() {
	file, information, err := this.GetFile("file")  //返回文件，文件信息头，错误信息
	if err != nil {
	    this.Ctx.WriteString("File retrieval failure")
	    return
	} else {
	    filename := information.Filename
	    picture := strings.Split(filename,".")      //读取到字符串，并以.符号分隔开
	    layout := strings.ToLower(picture[len(picture)-1])          //把字母字符转换成小写，非字母字符不做出处理,返回此字符串转换为小写形式的副本。
    
	    if layout != "jpg" && layout != "png" && layout != "gif" {
		this.Ctx.WriteString("请上传符合格式的图片（png、jpg、gif）")
		return      //结束整个程序，不执行保存文件
	    }
    
	    err = this.SaveToFile("file",path.Join("static/upload",filename))
	    if err != nil {
		this.Ctx.WriteString("File upload failed！")
	    } else {
		this.Ctx.WriteString("File upload succeed!")
	    }
	}
    
	defer file.Close()    //关闭上传的文件，否则出现零食文件不清除的情况
	this.TplName = "upload.html"
}