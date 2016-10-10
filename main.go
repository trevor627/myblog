package main

import (
	"beeblog/controllers"
	"beeblog/models"
	_ "beeblog/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.Router("/login", &controllers.LoginController{})

	os.Mkdir("attachment", os.ModePerm)
	//作为静态文件
	//beego.SetStaticPath("/attachment", "attachment")
	//作为一个控制器
	beego.RESTRouter("/attachment/:all", &controllers.AttachController{})

	beego.Run()
}
