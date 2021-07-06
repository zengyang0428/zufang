package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zufang/models"
	_ "github.com/astaxie/beego/cache/redis"
	
	"fmt"
)

type AreaController struct {
	beego.Controller
}

func(this*AreaController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	
	fmt.Println("connect success")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.RetData(resp)

	


	//fmt.Println("cache_conn.aa =",cache_conn.Get("aaa"))
	//fmt.Printf("cache_conn ,conn[aa]= %s\n",cache_conn.Get("aaa"))

	//从mysql数据库拿到area数据
	var areas  []models.Area

	o := orm.NewOrm()
	num ,err :=o.QueryTable("area").All(&areas)

	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
 	}
 	if num ==0{
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}

	resp["data"] = areas

	

	//cache_conn.Put("area",json_str,time.Second*3600)


		//打包成json返回给前段
	fmt.Println("query data sucess ,resp =",resp,"num =",num)


}
