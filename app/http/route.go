package http

import (
	"github.com/afengliz/zett/app/http/module/user"
	"github.com/afengliz/zett/framework/gin"
)

func Route(engine *gin.Engine){
	engine.GET("/list", user.UserListControllerHandler)
}
