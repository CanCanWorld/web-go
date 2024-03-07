package service

import (
	"web-go/model"
	"web-go/model/table"
	"web-go/serializer"
)

type CreateArticleService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Mid     uint   `json:"mid" form:"mid"`
}

func (service CreateArticleService) Create(uid uint) serializer.Response {
	var article = table.Article{
		Uid:     uid,
		Mid:     service.Mid,
		Title:   service.Title,
		Content: service.Content,
	}
	model.DB.Create(&article)
	return serializer.Response{
		Status: 200,
		Data:   article,
		Msg:    "添加成功",
	}
}

func (service CreateArticleService) Query() serializer.Response {
	var articles []table.Article
	return serializer.Response{
		Status: 200,
		Data:   articles,
		Msg:    "查询成功",
	}
}
