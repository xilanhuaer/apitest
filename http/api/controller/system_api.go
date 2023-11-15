package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/common/response"
	"github.com/xilanhuaer/http-client/dal/model"
)

type SystemApi struct{}

func (systemApi *SystemApi) Create(context *gin.Context) {
	var system model.System
	if err := context.ShouldBindJSON(&system); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	if err := systemService.Create(system); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}

func (systemApi *SystemApi) List(context *gin.Context) {
	params := context.Query("params")
	page, err := systemService.List(params)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(page, context)
}

func (systemApi *SystemApi) Find(context *gin.Context) {
	id := context.Param("id")
	system, err := systemService.Find(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(system, context)
}

func (systemApi *SystemApi) Update(context *gin.Context) {
	var system model.System
	if err := context.ShouldBindJSON(&system); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	id := context.Param("id")
	if err := systemService.Update(id, system); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}
