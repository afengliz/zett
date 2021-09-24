package framework

import (
	"encoding/json"
	"html/template"
)

type IResponse interface {
	JsonP(obj interface{}) error
}
func (c *Context) JsonP(obj interface{}) error{
	callbackfunc,_ :=c.QueryString("callback","callback_function")
	c.SetHeader("Content-Type","application/javascript")
	callback := template.JSEscapeString(callbackfunc)
	c.responseWriter.Write([]byte(callback))
	_,err := c.responseWriter.Write([]byte("("))
	if err != nil{
		return err
	}
	rawData,err :=json.Marshal(obj)
	if err != nil{
		return err
	}
	_,err =c.responseWriter.Write(rawData)
	if err != nil{
		return err
	}
	_,err =c.responseWriter.Write([]byte(")"))
	if err != nil{
		return err
	}
	return err
}

