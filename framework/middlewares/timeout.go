package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/afengliz/zett/framework"
	"time"
)

func TimeoutMiddleware(t time.Duration) framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		panicChan := make(chan interface{}, 1)
		finishChan := make(chan error, 1)
		toutCtx, cancel := context.WithTimeout(ctx.BaseContext(), t)
		defer cancel()
		go func() {
			defer func() {
				if err := recover(); err != nil {
					panicChan <- err
				}
			}()
			ans := ctx.Next()
			finishChan <- ans
		}()
		var anserr error
		select {
		case <-finishChan:
			//fmt.Println("normal finish")
		case err := <-panicChan:
			fmt.Println("panic:", err)
			ctx.Json(500, "panic")
			anserr = errors.New(fmt.Sprintf("panic err:%+v", err))
		case <-toutCtx.Done():
			ctx.Json(500, "time out")
			ctx.SetHasTimeOut()
			anserr = errors.New("time out")
		}
		return anserr
	}
}
