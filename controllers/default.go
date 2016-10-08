package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"

	c.Data["IsLogin"] = checkAccount(c.Ctx)

<<<<<<< HEAD
	topics, err := models.GetAllTopics(true)
=======
	topics, err := models.GetAllTopics(c.Input().Get("cate"), c.Input().Get("label"), true)
>>>>>>> newbranch
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
<<<<<<< HEAD
=======

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories

>>>>>>> newbranch
}
