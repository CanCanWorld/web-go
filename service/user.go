package service

import (
	"github.com/jinzhu/gorm"
	"web-go/model"
	"web-go/model/table"
	"web-go/pkg/utils"
	"web-go/serializer"
)

type UserService struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user table.User
	var count int
	model.DB.Model(&table.User{}).Where("username=?", service.Username).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经有这个人了",
		}
	}
	user.Username = service.Username
	//密码加密
	var err = user.SetPassword(service.Password)
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	//创建用户
	err = model.DB.Create(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user table.User
	var err = model.DB.Where("username=?", service.Username).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先注册",
			}
		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//发token，用于其他接口的身份验证
	var token, e = utils.GenerateToken(user.ID, service.Username, service.Password)
	if e != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}
}
