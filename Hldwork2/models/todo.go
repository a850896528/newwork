package models
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//登录信息
type User struct {
	Id string `json:"id" form:"id"`
	Psword string `json:"pwd" form:"pwd"`
	Name string `json:"uname" form:"uname"`
	Tel string `json:"tel" form:"tel"`
}
type NewUser struct {
	Psword string `json:"pwd" form:"pwd" binding:"required" `
	Name string `json:"uname" form:"uname" binding:"required"`
	Tel string `json:"tel" form:"tel" binding:"required"`
}
//账户信息
type UserMessage struct {
	Picture string `json:"picture"`
	Id string   `json:"id"`
	Action int `json:"action"`
	Attention int `json:"attention"`
	Lv int `json:"lv"`
	Fan int `json:"fan"`
}
//视频
type AllMovie struct {
	Name string
	Broadcast string
	Address string
	Point int
	Picture string
	Upname string
}
//视频信息
type Info struct {
	Like int `json:"like"`
	Coin int `json:"coin"`
	Collect int `json:"name"`
}
//留言板
type Todo struct{
	Name string `json:"name"`
	Message string `json:"message"`
	Time string `json:"time"`
}
//视频number
type Movie1 struct {
	Number int `json:"number"`
}
//评论人的id
type Movie struct {
	Id string `json:"id"`
}
//留言
type Message struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Message string `json:"message"`
}
func SetCookie()gin.HandlerFunc{
	return func(c *gin.Context){
		if cookie,err:=c.Cookie("status");err==nil{
			if cookie =="1"{
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
