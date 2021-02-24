package routers

import (
	"Hldwork/controller"
	"github.com/gin-gonic/gin"
)
func SetupRouter()*gin.Engine{
	r:=gin.Default()
	r.Static("/static","./static")
	r.LoadHTMLGlob("templates/*")//模板解析
	//模板渲染
	//登录
	/*r.GET("/login",controller.Login)
	r.POST("/login",models.SetCookie(),controller.Signin)*/
	//注册
    r.GET("/try",controller.Register)
	r.POST("/try",controller.CreatUser,controller.NewUser)
    r.GET("/index",controller.Index,controller.Information,controller.Allmovie)
	//主页
	//r.GET("main",controller.Allmovie)
	//视频页
	/*v1Group:=r.Group("v1")
	{
		//查看赞，投币，收藏
		v1Group.GET("/todo")
		//赞，投币，收藏，三连
		v1Group.PUT("/todo",controller.Like)
		v1Group.PUT("/todo",controller.Coin)
		v1Group.PUT("/todo",controller.Collect)
		v1Group.PUT("/todo",controller.Triple)
	}*/
	return r
}