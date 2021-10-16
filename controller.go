package main

import (
	"fmt"
	"github.com/afengliz/zett/app/provider/demo"
	"github.com/afengliz/zett/framework/gin"
	"github.com/spf13/cast"
	"time"
)

func RootControllerHandler(ctx *gin.Context) {
	ctx.Writer.Write([]byte("I am /"))
}
func FooControllerHandler(ctx *gin.Context) {
	ctx.Writer.Write([]byte("I am /foo"))
}

func HelloControllerHandler(ctx *gin.Context) {
	ctx.Writer.Write([]byte("I am /hello"))
}

func UserInfoGetControllerHandler(ctx *gin.Context) {
	time.Sleep(time.Millisecond * 500)
	paramId, _ := ctx.DefaultParamInt("id", -1)
	ctx.IJson(200, fmt.Sprintf("I am /user/:id/info data:%d", paramId))
}

func UserRootControllerHandler(ctx *gin.Context)  {
	ctx.Writer.Write([]byte("I am /user root"))
}
func UserListControllerHandler(ctx *gin.Context)  {
	ctx.Writer.Write([]byte("I am /user/list"))

}

func UserVipVersionControllerHandler(ctx *gin.Context) {
	ctx.Writer.Write([]byte("I am /user/vip/version"))
}

func GetQueryParamControllerHandler(ctx *gin.Context)  {
	userid, _ := ctx.DefaultQueryInt("id", -1)
	time.Sleep(time.Second*15)
	ctx.Writer.Write([]byte(cast.ToString(userid)))
}

func PostFormParamControllerHandler(ctx *gin.Context) {
	userid, _ := ctx.DefaultFormInt("id", -1)
	ctx.Writer.Write([]byte(cast.ToString(userid)))
}

func TestJsonControllerHandler(ctx *gin.Context)  {
	type Student struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}
	s := Student{}
	err := ctx.BindJson(&s)
	if err != nil{
		ctx.IJson(500,err.Error())
	}
	res := fmt.Sprintf("%+v",s)
	ctx.IJson(200,res)
}


func TestXmlControllerHandler(ctx *gin.Context) {
	type Student struct {
		Name string `xml:"name"`
		Age int `xml:"age"`
	}
	s := Student{}
	err := ctx.BindXml(&s)
	if err != nil{
		ctx.IJson(500,err.Error())
	}
	ctx.IXml(s)
}

func TestGetClientAddress(ctx *gin.Context)  {
	ctx.IJson(200,ctx.ClientIp())
}

func TestFormFile(ctx *gin.Context) {
	header,_ := ctx.FormFile("field-name")
	fmt.Printf("%+v",header)
	ctx.IJson(200,ctx.ClientIp())
}

func TestHeaderControllerHandler(ctx *gin.Context) {
	hMap := ctx.GetHeader("Username")
	ctx.IJson(200,hMap)
}

func TestCookieControllerHandler(ctx *gin.Context)  {
	hMap := ctx.Cookies()
	ctx.IJson(200,hMap)
}

func TestJsonPControllerHandler(ctx *gin.Context) {
	type Student struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}
	ctx.IJsonP(Student{"liyanfeng",26})
}

func TestHtmlControllerHandler(ctx *gin.Context) {
	type Student struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}
	 ctx.IHtml("./test_html_template.html",Student{"liyanfeng",26})
}

func TestTextControllerHandler(ctx *gin.Context) {
	format := "my name is %s,age is %d"
	ctx.IText(format,"liyanfeng",26)
}

func TestRedirectControllerHandler(ctx *gin.Context){
	ctx.IRedirect("/user/info")
}


func PostDemoTestControllerHandler(ctx *gin.Context){
	service :=ctx.MustMake(demo.Key).(*demo.DemoService)
	ctx.IJson(200,service.GetFoo())
}

