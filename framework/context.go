package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Context struct{
	request *http.Request
	responseWriter http.ResponseWriter
	ctx context.Context
	handler ControllerHandler
	writerMux *sync.Mutex
	hasTimeOut bool
}
var _ context.Context = (*Context)(nil)

func NewContext(r *http.Request,w  http.ResponseWriter) *Context{
	return &Context{request: r,responseWriter: w,ctx: r.Context(),writerMux: &sync.Mutex{}}
}

func (c *Context) GetRequest() *http.Request{
	return c.request
}

func(c *Context) GetResponseWriter() http.ResponseWriter{
	return c.responseWriter
}

func (c *Context) WriterMux() *sync.Mutex{
	return c.writerMux
}

func (c *Context) SetHasTimeOut(){
	c.hasTimeOut = true
}
func (c *Context) HasTimeOut() bool{
	return c.hasTimeOut
}
func (c *Context) SetHandler(handler ControllerHandler){
	c.handler = handler
}

func (c *Context) BaseContext() context.Context{
	return c.request.Context()
}


func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func (c *Context) Json(status int,data interface{}) error{
	if c.HasTimeOut(){
		return nil
	}
	c.responseWriter.Header().Set("Content-Type","application/json")
	c.responseWriter.WriteHeader(status)
	if data != nil{
		byt,err := json.Marshal(data)
		if err != nil{
			c.responseWriter.WriteHeader(500)
			return err
		}
		c.responseWriter.Write(byt)
	}
	return nil
}

