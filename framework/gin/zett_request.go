package gin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"mime/multipart"
)
type IRequest interface {
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParamStringSlice(key string, def []string) ([]string, bool)
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
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
	Cookies() map[string]string

}

var _ IRequest = (*Context)(nil)

func (c *Context) QueryAll() map[string][]string {
	if c.Request != nil {
		return c.Request.URL.Query()
	}
	return map[string][]string{}
}

func (c *Context) DefaultQueryInt(key string, def int) (int, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToInt(items[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToInt64(items[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToFloat32(items[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToFloat64(items[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToBool(items[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryString(key string, def string) (string, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return cast.ToString(items[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	if items, ok := c.QueryAll()[key]; ok {
		if len(items) > 0 {
			return items, true
		}
	}
	return def, false
}



func (c *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := c.ZettParam(key); val != nil {
		return cast.ToInt(val), true
	}
	return def, false
}

func (c *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if val := c.ZettParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (c *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := c.ZettParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (c *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := c.ZettParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (c *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if val := c.ZettParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (c *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := c.ZettParam(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

func (c *Context) DefaultParamStringSlice(key string, def []string) ([]string, bool) {
	if val := c.ZettParam(key); val != nil {
		return cast.ToStringSlice(val), true
	}
	return def, false
}

func (c *Context) ZettParam(key string) interface{} {
	if val, ok := c.Params.Get(key); ok {
		return val
	}
	return nil
}

func (c *Context) DefaultFormInt(key string, def int) (int, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) DefaultFormString(key string, def string) (string, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToString(vals[0]),true
		}
	}
	return def,false
}

func (c *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
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
	c.initFormCache()
	return c.formCache
}



func (c *Context) BindJson(obj interface{}) error{
	if c.Request != nil{
		body,err := ioutil.ReadAll(c.Request.Body)
		if err != nil{
			return err
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		err = json.Unmarshal(body,obj)
		if err != nil{
			return err
		}
		return nil
	}
	return errors.New("ctx.request empty")
}

func (c *Context) BindXml(obj interface{}) error{
	if c.Request != nil{
		body,err := ioutil.ReadAll(c.Request.Body)
		if err != nil{
			return err
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		err = xml.Unmarshal(body,obj)
		if err != nil{
			return err
		}
		return nil
	}
	return errors.New("ctx.request empty")
}

func (c *Context) GetRawData() ([]byte, error) {
	if c.Request != nil {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return nil,err
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		return body,nil
	}
	return nil,nil
}

func (c *Context) Uri() string{
	var uri string
	if c.Request != nil{
		uri = c.Request.RequestURI
	}
	return uri
}
func (c *Context) Host() string{
	var host string
	if c.Request != nil{
		host = c.Request.Host
	}
	return host
}
func (c *Context) Method() string{
	var mtname string
	if c.Request != nil{
		mtname = c.Request.Method
	}
	return mtname
}

func (c *Context) ClientIp() string{
	var ipAddress string
	if r := c.Request;r != nil{
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
	if c.Request != nil{
		return c.Request.Header
	}
	return map[string][]string{}
}

func (c *Context) Cookies() map[string]string {
	ans := map[string]string{}
	if c.Request != nil{
		arr := c.Request.Cookies()
		for i:=0;i<len(arr);i++{
			ans[arr[i].Name] = arr[i].Value
		}
	}
	return ans
}

