package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"mime/multipart"
	"strings"
)
const defaultMultipartMemory = 32 << 20 // 32 MB
type IRequest interface {
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	ParamStringSlice(key string, def []string) ([]string, bool)
	Param(key string) interface{}
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	Form(key string) interface{}
	FormAll() map[string][]string
	FormFile(key string) (*multipart.FileHeader, error)
	BindJson(obj interface{}) error
	BindXml(obj interface{}) error
	Uri() string
	Host() string
	Method() string
	ClientIp() string
	Headers() map[string][]string
	Header(string) (string,bool)
	Cookies() map[string]string
	Cookie(string) (string,bool)

}

var _ IRequest = (*Context)(nil)

func (c *Context) QueryAll() map[string][]string {
	if c.request != nil {
		return c.request.URL.Query()
	}
	return map[string][]string{}
}

func (c *Context) QueryInt(key string, def int) (int, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToInt(items[0]), true
		}
	}
	return def, false
}

func (c *Context) QueryInt64(key string, def int64) (int64, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToInt64(items[0]), true
		}
	}
	return def, false
}

func (c *Context) QueryFloat32(key string, def float32) (float32, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToFloat32(items[0]), true
		}
	}
	return def, false
}

func (c *Context) QueryFloat64(key string, def float64) (float64, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToFloat64(items[0]), true
		}
	}
	return def, false
}

func (c *Context) QueryBool(key string, def bool) (bool, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToBool(items[0]), true
		}
	}
	return def, false
}

func (c *Context) QueryString(key string, def string) (string, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToString(items[0]), true
		}
	}
	return def, false
}

func (c *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return items, true
		}
	}
	return def, false
}

func (c *Context) Query(key string) interface{} {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return items[0]
		}
	}
	return nil
}

func (c *Context) ParamInt(key string, def int) (int, bool) {
	if val := c.Param(key); val != nil {
		return cast.ToInt(val), true
	}
	return def, false
}

func (c *Context) ParamInt64(key string, def int64) (int64, bool) {
	if val := c.Param(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (c *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if val := c.Param(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (c *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if val := c.Param(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (c *Context) ParamBool(key string, def bool) (bool, bool) {
	if val := c.Param(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (c *Context) ParamString(key string, def string) (string, bool) {
	if val := c.Param(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

func (c *Context) ParamStringSlice(key string, def []string) ([]string, bool) {
	if val := c.Param(key); val != nil {
		return cast.ToStringSlice(val), true
	}
	return def, false
}

func (c *Context) Param(key string) interface{} {
	key = strings.ToUpper(key)
	if c.params != nil {
		if val, ok := c.params[key]; ok {
			return val
		}
	}
	return nil
}

func (c *Context) FormInt(key string, def int) (int, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) FormInt64(key string, def int64) (int64, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) FormFloat32(key string, def float32) (float32, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) FormFloat64(key string, def float64) (float64, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) FormBool(key string, def bool) (bool, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) FormString(key string, def string) (string, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToString(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) FormStringSlice(key string, def []string) ([]string, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToStringSlice(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) Form(key string) interface{} {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}

func (c *Context) FormAll() map[string][]string {
	if c.request != nil {
		c.request.ParseForm()
		return c.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormFile(key string) (*multipart.FileHeader, error) {
	if ctx.request.MultipartForm == nil {
		if err := ctx.request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := ctx.request.FormFile(key)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

func (c *Context) BindJson(obj interface{}) error{
	if c.request != nil{
		body,err := ioutil.ReadAll(c.request.Body)
		if err != nil{
			return err
		}
		c.request.Body = ioutil.NopCloser(bytes.NewReader(body))
		err = json.Unmarshal(body,obj)
		if err != nil{
			return err
		}
		return nil
	}
	return errors.New("ctx.request empty")
}

func (c *Context) BindXml(obj interface{}) error{
	if c.request != nil{
		body,err := ioutil.ReadAll(c.request.Body)
		if err != nil{
			return err
		}
		c.request.Body = ioutil.NopCloser(bytes.NewReader(body))
		err = xml.Unmarshal(body,obj)
		if err != nil{
			return err
		}
		return nil
	}
	return errors.New("ctx.request empty")
}

func (c *Context) GetRawData() ([]byte, error) {
	if c.request != nil {
		body, err := ioutil.ReadAll(c.request.Body)
		if err != nil {
			return nil,err
		}
		c.request.Body = ioutil.NopCloser(bytes.NewReader(body))
		return body,nil
	}
	return nil,nil
}

func (c *Context) Uri() string{
	var uri string
	if c.request != nil{
		uri = c.request.RequestURI
	}
	return uri
}
func (c *Context) Host() string{
	var host string
	if c.request != nil{
		host = c.request.Host
	}
	return host
}
func (c *Context) Method() string{
	var mtname string
	if c.request != nil{
		mtname = c.request.Method
	}
	return mtname
}

func (c *Context) ClientIp() string{
	var ipAddress string
	if r := c.request;r != nil{
		ipAddress = r.Header.Get("X-Real-Ip")
		if ipAddress == "" {
			ipAddress = r.Header.Get("X-Forwarded-For")
		}
		if ipAddress == "" {
			ipAddress = r.RemoteAddr
		}
	}
	return ipAddress
}

func (c *Context) Headers() map[string][]string {
	if c.request != nil{
		return c.request.Header
	}
	return map[string][]string{}
}

func (c *Context) Header(key string) (string, bool) {
	if c.request != nil{
		if vals,ok:=c.request.Header[key];ok{
			return vals[0],true
		}
	}
	return "",false
}

func (c *Context) Cookies() map[string]string {
	ans := map[string]string{}
	if c.request != nil{
		arr := c.request.Cookies()
		for i:=0;i<len(arr);i++{
			ans[arr[i].Name] = arr[i].Value
		}
	}
	return ans
}

func (c *Context) Cookie(key string) (string,bool) {
	cMap := c.Cookies()
	if val,ok:=cMap[key];ok{
		return val,true
	}
	return "",false
}