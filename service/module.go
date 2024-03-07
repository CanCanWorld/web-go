package service

import (
	"github.com/jinzhu/gorm"
	"web-go/model"
	"web-go/model/table"
	"web-go/serializer"
)

type CreateModuleService struct {
	Title string `json:"title" form:"title"`
}

func (service *CreateModuleService) Create() serializer.Response {
	code := 200
	module := table.Module{
		Model: gorm.Model{},
		Title: service.Title,
	}
	model.DB.Create(&module)
	return serializer.Response{
		Status: code,
		Data:   module,
		Msg:    "创建成功",
	}
}

func QueryModule() serializer.Response {
	var modules []table.Module
	model.DB.Find(&modules)
	return serializer.Response{
		Status: 200,
		Data:   modules,
		Msg:    "查询成功",
	}
}
