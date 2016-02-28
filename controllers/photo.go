package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/parser/response"
	"github.com/instagram-beego/repository"
	"strconv"
)

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) GetAll() {
	photoRepository := repository.PhotoRepository{}
	photos, err := photoRepository.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	this.Data["json"] = &photos

	this.ServeJSON()
}

func (this *PhotoController) GetByUserId() {
	photoRepository := repository.PhotoRepository{}
	userId, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	photos, err := photoRepository.GetByUserId(userId)

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = &response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	} else {
		this.Data["json"] = &photos
	}

	this.ServeJSON()
}
