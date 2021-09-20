package middlewares

import (
	"fmt"
	"github.com/afengliz/zett/framework"
)

func Test3() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		fmt.Println("middleware pre test3")
		ctx.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}
