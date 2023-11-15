package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/common/response"
	"github.com/xilanhuaer/http-client/dal/model"
)

type ModuleApi struct{}

func (moduleApi *ModuleApi) Create(context *gin.Context) {
	var module model.Module
	if err := context.ShouldBindJSON(&module); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	if err := moduleService.Create(module); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}

func (moduleApi *ModuleApi) List(context *gin.Context) {
	params := context.Query("params")
	page, err := moduleService.List(params)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(page, context)
}

func (moduleApi *ModuleApi) Find(context *gin.Context) {
	id := context.Param("id")
	module, err := moduleService.Find(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(module, context)
}

func (moduleApi *ModuleApi) Update(context *gin.Context) {
	id := context.Param("id")
	var module model.Module
	if err := context.ShouldBindJSON(&module); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	if err := moduleService.Update(id, module); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}
