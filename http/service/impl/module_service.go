package impl

import (
	"context"
	"github.com/xilanhuaer/http-client/common/response"
	"github.com/xilanhuaer/http-client/dal/model"
	"github.com/xilanhuaer/http-client/dal/query"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/utils"
)

type ModuleService struct{}

func (moduleService *ModuleService) Create(module model.Module) error {
	if module.SystemName == "" {
		system, err := query.System.WithContext(context.Background()).
			Where(query.System.ID.Eq(module.SystemID)).
			First()
		if err != nil {
			return err
		}
		module.SystemName = system.Name
	}
	return query.Module.WithContext(context.Background()).Create(&module)
}
func (moduleService *ModuleService) List(params string) (response.Page, error) {
	var (
		module []model.Module
		count  int64
	)
	q, err := utils.ParseCondition(global.DB.Model(&model.Module{}), params)
	if err != nil {
		return response.Page{}, err
	}
	if err = q.Count(&count).Error; err != nil {
		return response.Page{}, err
	}
	if err = q.Find(&module).Error; err != nil {
		return response.Page{}, err
	}
	return response.Page{List: module, Total: count}, nil
}
func (moduleService *ModuleService) Find(id string) (model.Module, error) {
	idInt, err := utils.StringToInt32(id)
	if err != nil {
		return model.Module{}, err
	}
	module, err := query.Module.WithContext(context.Background()).Where(query.Module.ID.Eq(idInt)).First()
	if err != nil {
		return model.Module{}, err
	}
	return *module, nil
}
func (moduleService *ModuleService) Update(id string, module model.Module) error {
	idInt, err := utils.StringToInt32(id)
	if err != nil {
		return err
	}
	data := utils.StructToMap(module)
	_, err = query.Module.WithContext(context.Background()).Where(query.Module.ID.Eq(idInt)).Updates(data)
	return err
}
