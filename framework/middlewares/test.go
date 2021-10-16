package middlewares

import (
	"fmt"
	"github.com/afengliz/zett/framework/gin"
)

func Test3() gin.HandlerFunc {
	return func(ctx *gin.Context)  {
		fmt.Println("middleware pre test3")
		ctx.Next()
		fmt.Println("middleware post test3")
	}
}
