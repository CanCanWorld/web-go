package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"web-go/api"
	"web-go/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("my_session", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		v1.POST("user/loginOrRegister", api.UserLoginOrRegister)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("queryModule", api.QueryModule)
			authed.POST("createModule", api.CreateModule)
			authed.POST("createArticle", api.CreateArticle)
			authed.POST("queryArticle", api.QueryArticle)
			authed.POST("createComment", api.CreateComment)
		}
	}
	return r
}
