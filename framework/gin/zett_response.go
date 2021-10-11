package gin


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
	IJson(status int, data interface{}) error
	// jsonp
	IJsonP(obj interface{}) error
	//xml 输出
	IXml(obj interface{}) error
	// html 输出
	IHtml(template string, obj interface{}) error
	// string
	IText(format string, values ...interface{}) error
	// 重定向
	IRedirect(path string) IResponse
	// header
	ISetHeader(key string, val string) IResponse
	// Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse
	// 设置状态码
	ISetStatus(code int) IResponse
	// 设置 200 状态
	ISetOkStatus() IResponse
}
var _ IResponse = (*Context)(nil)
func (c *Context) JsonOk(obj interface{}) error {
	return c.ISetOkStatus().IJson(200, obj)
}
func (c *Context) IJson(status int, data interface{}) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	if data != nil {
		byt, err := json.Marshal(data)
		if err != nil {
			return err
		}
		c.Writer.Write(byt)
	}
	return nil
}

func (c *Context) IJsonP(obj interface{}) error {
	callbackfunc, _ := c.DefaultQueryString("callback", "callback_function")
	c.ISetHeader("Content-Type", "application/javascript")
	callback := template.JSEscapeString(callbackfunc)
	c.Writer.Write([]byte(callback))
	_, err := c.Writer.Write([]byte("("))
	if err != nil {
		return err
	}
	rawData, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = c.Writer.Write(rawData)
	if err != nil {
		return err
	}
	_, err = c.Writer.Write([]byte(")"))
	if err != nil {
		return err
	}
	return err
}

func (c *Context) IXml(obj interface{}) error {
	bArr, err := xml.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = c.Writer.Write(bArr)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) IHtml(fileName string, obj interface{}) error {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	err = tmpl.Execute(c.Writer, obj)
	if err != nil {
		return err
	}
	c.ISetHeader("Content-Type", "application/html")
	return nil
}

func (c *Context) IText(format string, obj ...interface{}) error {
	tstr := fmt.Sprintf(format, obj...)
	c.ISetHeader("Content-Type", "application/text")
	_, err := c.Writer.Write([]byte(tstr))
	return err
}

func (c *Context) ISetOkStatus() IResponse {
	c.Writer.WriteHeader(http.StatusOK)
	return c
}

func (c *Context) IRedirect(path string) IResponse {
	http.Redirect(c.Writer,c.Request,path,http.StatusMovedPermanently)
	return c
}

func (c *Context) ISetHeader(key string, val string) IResponse {
	c.Writer.Header().Set(key, val)
	return c
}

func (c *Context) ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	http.SetCookie(c.Writer,&http.Cookie{
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

func (c *Context) ISetStatus(code int) IResponse {
	c.Writer.WriteHeader(code)
	return c
}
