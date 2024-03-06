package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"web-go/service"
)

func CreateModule(c *gin.Context) {
	var s service.CreateModuleService
	err := c.ShouldBind(&s)
	if err == nil {
		var res = s.Create()
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(400, err)
	}
}
