package api

import (
	"github.com/gin-gonic/gin"
	"web-go/service"
)

func UserRegister(c *gin.Context) {
	var userService service.UserService
	err := c.ShouldBind(&userService)
	if err == nil {
		res := userService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func UserLogin(c *gin.Context) {
	var userService service.UserService
	err := c.ShouldBind(&userService)
	if err == nil {
		res := userService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func UserLoginOrRegister(c *gin.Context) {
	var userService service.UserService
	err := c.ShouldBind(&userService)
	if err == nil {
		res := userService.LoginOrRegister()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
