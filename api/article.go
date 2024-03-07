package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"web-go/pkg/utils"
	"web-go/service"
)

func CreateArticle(c *gin.Context) {
	var s service.CreateArticleService
	uid := utils.GetUidByContext(c)
	err := c.ShouldBind(&s)
	if err == nil {
		var res = s.Create(uid)
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(400, err)
	}
}

func QueryArticle(c *gin.Context) {
	var s service.QueryArticleService
	err := c.ShouldBind(&s)
	if err == nil {
		var res = s.Query()
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(400, err)
	}
}
