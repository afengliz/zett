package middlewares

import (
	"context"
	"fmt"
	"github.com/afengliz/zett/framework/gin"
	"time"
)

func TimeoutMiddleware(t time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context)  {
		panicChan := make(chan interface{}, 1)
		finishChan := make(chan struct{}, 1)
		toutCtx, cancel := context.WithTimeout(ctx.BaseContext(), t)
		defer cancel()
		go func() {
			defer func() {
				if err := recover(); err != nil {
					panicChan <- err
				}
			}()
			ctx.Next()
			finishChan <- struct {}{}
		}()
		select {
		case <-finishChan:
			fmt.Println("finish")
		case err := <-panicChan:
			fmt.Println("panic:", err)
			ctx.IJson(500, "panic")
		case <-toutCtx.Done():
			ctx.IJson(500, "time out")
		}
	}
}
