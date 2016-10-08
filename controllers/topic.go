package controllers

import (
	"beeblog/models"
	"strings"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"

<<<<<<< HEAD
	topics, err := models.GetAllTopics(false)
=======
	topics, err := models.GetAllTopics("", "", false)
>>>>>>> newbranch
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	label := this.Input().Get("label")

	var err error
	if len(tid) == 0 {
<<<<<<< HEAD
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
=======
		err = models.AddTopic(title, category, label, content)
	} else {
		err = models.ModifyTopic(tid, title, category, label, content)
>>>>>>> newbranch
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
<<<<<<< HEAD
	this.Data["Tid"] = this.Ctx.Input.Param("0")
}

func (this *TopicController) Modify() {
=======
	this.Data["Labels"] = strings.Split(topic.Labels, " ")

	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
>>>>>>> newbranch
	this.TplName = "topic_modify.html"
	tid := this.Input().Get("tid")

	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
<<<<<<< HEAD
=======
	this.Data["IsLogin"] = true
>>>>>>> newbranch
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 302)
}
