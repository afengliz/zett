package user

import "github.com/afengliz/zett/framework/gin"

func UserListControllerHandler(ctx *gin.Context)  {
	ctx.Writer.Write([]byte("I am /user/list"))

}
