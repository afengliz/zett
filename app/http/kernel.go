package http

import "github.com/afengliz/zett/framework/gin"

func NewHttpEngine() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	Route(engine)
	return engine
}
