package service

import (
	"fmt"
	"mime/multipart"
	"web-go/model"
	"web-go/model/table"
	"web-go/serializer"
)

type CreateArticleService struct {
	Title   string                 `json:"title" form:"title"`
	Content string                 `json:"content" form:"content"`
	Mid     uint                   `json:"mid" form:"mid"`
	Images  []multipart.FileHeader `json:"images" form:"images"`
}

func (service CreateArticleService) Create(uid uint) serializer.Response {
	var article = table.Article{
		Uid:     uid,
		Mid:     service.Mid,
		Title:   service.Title,
		Content: service.Content,
	}
	for _, image := range service.Images {
		fmt.Println("file: ", image.Filename)
	}
	model.DB.Create(&article)
	return serializer.Response{
		Status: 200,
		Data:   article,
		Msg:    "添加成功",
	}
}

type QueryArticleService struct {
	Title    string `json:"title" form:"title"`
	Mid      uint   `json:"mid" form:"mid"`
	Uid      uint   `json:"uid" form:"uid"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

func (service QueryArticleService) Query() serializer.Response {
	var articles []table.Article
	db := model.DB
	if service.Mid != 0 {
		fmt.Println("mid != 0")
		db = db.Where("mid=?", service.Mid)
	}
	db.Find(&articles)
	return serializer.Response{
		Status: 200,
		Data:   articles,
		Msg:    "查询成功",
	}
}
