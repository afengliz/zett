package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type IResponse interface {
	// json
	Json(status int, data interface{}) error
	// jsonp
	JsonP(obj interface{}) error
	//xml 输出
	Xml(obj interface{}) error
	// html 输出
	Html(template string, obj interface{}) error
	// string
	Text(format string, values ...interface{}) error
	// 重定向
	Redirect(path string) IResponse
	// header
	SetHeader(key string, val string) IResponse
	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse
	// 设置状态码
	SetStatus(code int) IResponse
	// 设置 200 状态
	SetOkStatus() IResponse
}
var _ IResponse = (*Context)(nil)
func (c *Context) JsonOk(obj interface{}) error {
	return c.SetOkStatus().Json(200, obj)
}
func (c *Context) Json(status int, data interface{}) error {
	if c.HasTimeOut() {
		return nil
	}
	c.responseWriter.Header().Set("Content-Type", "application/json")
	c.responseWriter.WriteHeader(status)
	if data != nil {
		byt, err := json.Marshal(data)
		if err != nil {
			return err
		}
		c.responseWriter.Write(byt)
	}
	return nil
}

func (c *Context) JsonP(obj interface{}) error {
	callbackfunc, _ := c.QueryString("callback", "callback_function")
	c.SetHeader("Content-Type", "application/javascript")
	callback := template.JSEscapeString(callbackfunc)
	c.responseWriter.Write([]byte(callback))
	_, err := c.responseWriter.Write([]byte("("))
	if err != nil {
		return err
	}
	rawData, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = c.responseWriter.Write(rawData)
	if err != nil {
		return err
	}
	_, err = c.responseWriter.Write([]byte(")"))
	if err != nil {
		return err
	}
	return err
}

func (c *Context) Xml(obj interface{}) error {
	bArr, err := xml.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = c.responseWriter.Write(bArr)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) Html(fileName string, obj interface{}) error {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	err = tmpl.Execute(c.responseWriter, obj)
	if err != nil {
		return err
	}
	c.SetHeader("Content-Type", "application/html")
	return nil
}

func (c *Context) Text(format string, obj ...interface{}) error {
	tstr := fmt.Sprintf(format, obj...)
	c.SetHeader("Content-Type", "application/text")
	_, err := c.responseWriter.Write([]byte(tstr))
	return err
}

func (c *Context) SetOkStatus() IResponse {
	c.responseWriter.WriteHeader(http.StatusOK)
	return c
}

func (c *Context) Redirect(path string) IResponse {
	http.Redirect(c.responseWriter,c.request,path,http.StatusMovedPermanently)
	return c
}

func (c *Context) SetHeader(key string, val string) IResponse {
	c.responseWriter.Header().Set(key, val)
	return c
}

func (c *Context) SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	http.SetCookie(c.responseWriter,&http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return c
}

func (c *Context) SetStatus(code int) IResponse {
	c.responseWriter.WriteHeader(code)
	return c
}
