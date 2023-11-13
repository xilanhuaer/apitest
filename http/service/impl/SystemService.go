package impl

import (
	"context"

	"github.com/xilanhuaer/http-client/common/response"
	"github.com/xilanhuaer/http-client/dal/model"
	"github.com/xilanhuaer/http-client/dal/query"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/utils"
)

type SystemService struct{}

func (systemService *SystemService) Create(system model.System) error {
	return query.System.WithContext(context.Background()).Create(&system)
}

func (systemService *SystemService) List(params string) (response.Page, error) {
	var (
		count  int64
		system []model.System
	)
	query, err := utils.ParseCondition(global.DB.Model(&model.System{}), params)
	if err != nil {
		return response.Page{}, err
	}
	if err = query.Count(&count).Error; err != nil {
		return response.Page{}, err
	}
	if err = query.Find(&system).Error; err != nil {
		return response.Page{}, err
	}
	return response.Page{List: system, Total: count}, nil
}

func (systemService *SystemService) Find(id string) (model.System, error) {
	idInt, err := exchange.IdExchange(id)
	if err != nil {
		return model.System{}, err
	}
	system, err := query.System.WithContext(context.Background()).Where(query.System.ID.Eq(idInt)).First()
	if err != nil {
		return model.System{}, err
	}
	return *system, nil
}

func (systemService *SystemService) Update(id string, system model.System) error {
	idInt, err := exchange.IdExchange(id)
	if err != nil {
		return err
	}
	data := exchange.StructToMap(system)
	_, err = query.System.WithContext(context.Background()).Where(query.System.ID.Eq(idInt)).Updates(data)
	return err
}
