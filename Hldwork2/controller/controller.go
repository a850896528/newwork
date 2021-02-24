package controller

import (
	"Hldwork/dao"
	"Hldwork/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
//界面
func Register(c *gin.Context){
	c.HTML(200,"try.html",nil)
}
func Login(c *gin.Context){
	c.HTML(200,"login.html",nil)
}
func Registe(c *gin.Context){
	c.HTML(200,"register.html",nil)
}
func Index(c *gin.Context){
	c.HTML(200,"index.html",nil)
}
//登录
func Signin(c *gin.Context){
	Id:=c.PostForm("id")
	Psword:=c.PostForm("psword")
	sqlStr:="select id , name ,psword from user where id=?"
	var u models.User
	err:=dao.DB.QueryRow(sqlStr,Id).Scan(&u.Id,&u.Name,&u.Psword)
	if err!=nil{
		fmt.Printf("scan fail:")
	}
	if u.Id == Id {
		if u.Psword==Psword{
			sqlStr1:="SELECT picture , fan , action , attention , lv , id FROM userinfo where id=?"
			var u models.UserMessage
			err:=dao.DB.QueryRow(sqlStr1,Id).Scan(&u.Picture,&u.Fan,&u.Action,&u.Attention,&u.Lv,&u.Id)
			if err!=nil{
				fmt.Printf("Sign get userinfo woring:%v\n",err)
			}
			c.JSON(1,gin.H{
				"picture":u.Picture,
				"id":u.Id,
				"fan":u.Fan,
				"acton":u.Action,
				"attention":u.Attention,
				"lv":u.Lv,
			})
			//设置cookie
			c.SetCookie("status","1",3600,"/", "localhost", false, true)
			return
		}
	}else{
		c.JSON(200,gin.H{
			"code":0,
			"登陆状态":"用户名与密码错误",
		})
	}
}
//创建用户
func CreatUser(c *gin.Context){
	var u models.NewUser
	if err:=c.ShouldBind(&u);err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"error": "Login information is not complete",
		})
	}
	sqlstr:="insert into user(name,tel,psword)values(?,?,?)"
	dao.DB.Exec(sqlstr,u.Name,u.Tel,u.Psword)
}
//新用户头像，关注，粉丝，动态，等级
func NewUser(c *gin.Context){
	sqlStr:="insert into userinfo(id)values(?)"
	sqlstr1:="select id from user where id =(select max(id) from user )"
	sqlstr2:="select picture ,attention , fan , action , lv , id from userinfo where id=?"
	var a int
	//得到最后的id
	err:=dao.DB.QueryRow(sqlstr1).Scan(a);if err!=nil{
		fmt.Printf("get newuser woring:%v\n",err)
	}
	//插入数据
	ret,err1:=dao.DB.Exec(sqlStr,a)
	if err1!=nil{
		fmt.Printf("put in newid woring:%v\n",err1)
		return
	}
	theID,err2:=ret.LastInsertId()
	if err2!=nil{
		fmt.Printf("get last id fail :%v\n",err2)
		return
	}
	var u models.UserMessage
	err3:=dao.DB.QueryRow(sqlstr2,theID).Scan(&u.Picture,&u.Attention,&u.Fan,&u.Action,&u.Lv,&u.Id)
	if err3!=nil{
		fmt.Printf("get all infomation woring:%d\n",err3)
		return
	}
	c.JSON(200,gin.H{
		"status":1,
		"picture":u.Picture,
		"attention":u.Attention,
		"fan":u.Fan,
		"action":u.Action,
		"id":u.Id,
		"lv":u.Lv,
	})
}
//得到粉丝关注头像id等级
func Information(c *gin.Context){
	var u models.UserMessage
	sqlstr:="select picture,attention, action ,fan,id,lv from userinfo where tel=?"
	tel:=c.Query("tel")
	dao.DB.QueryRow(sqlstr,tel).Scan(&u.Picture,&u.Attention,&u.Action,&u.Fan,&u.Id,&u.Lv)
	c.JSON(200,gin.H{
		"picture":u.Picture,
		"attention":u.Attention,
		"action":u.Action,
		"fan":u.Fan,
		"id":u.Id,
		"lv":u.Lv,
	})
}
//主页(数据库完成)
func Allmovie(c *gin.Context){
	sqlStr:="select name,broadcast,address,point,picture,upname,number from movie where id>?"
	rows,err:=dao.DB.Query(sqlStr,0)
	if err!=nil{
		fmt.Printf("main get allmovie fail:%v\n",err)
	}
	defer rows.Close()
	for rows.Next(){
		var m models.AllMovie
		err:=rows.Scan(&m.Name,&m.Broadcast,&m.Address,&m.Point,&m.Picture,&m.Upname)
		if err!=nil{
			fmt.Printf("allmoviie scan fail:%v\n",err)
			return
		}
		c.JSON(200,gin.H{
			"code":1,
			"name":m.Name,
			"broadcast":m.Broadcast,
			"address":m.Address,
			"point":m.Point,
			"picture":m.Picture,
			"upname":m.Upname,
		})
	}
}
//上传视频(前端没做割了)
func Upload(c *gin.Context){
	n:=c.PostForm("name")
	f,err:=c.FormFile("upload")
	time:=time.Now()
	picture,err:=c.FormFile("uploadPicture")
	if err!=nil{
		c.JSON(200,gin.H{
			"error":err.Error(),
		})
	}else {
		//将读取到的文件储存在本地
		dst:=fmt.Sprintf("./movie/%s",f.Filename)
		dst1:=fmt.Sprintf("./picture/%s",picture.Filename)
		c.SaveUploadedFile(f,dst)
		sqlStr:="insert into movie(address,name,time,picture) values(?,?,?,?)"
		ret,err:=dao.DB.Exec(sqlStr,dst,n,time,dst1)
		if err!=nil{
			fmt.Printf("upload woring:%d\n",err)
		}
		theNum,err:=ret.LastInsertId()
		if err!=nil{
			fmt.Printf("get upload num woring:%d\n",err)
			return
		}
		c.JSON(200,gin.H{
			"code":1,
			"number":theNum,
			"path":dst,
		})
	}
}
//进入视频页面（数据库完成）
func InMovie(c *gin.Context){
	sqlStr:="select like ,collect,coin,address,from movie1 where number=?"
	var u models.Info
	var m models.Movie1
	var address string
	err1:=c.ShouldBind(&m)
	if err1!=nil{
		c.JSON(400,gin.H{
			"error":err1.Error(),
		})
	}
	err:=dao.DB.QueryRow(sqlStr,m.Number).Scan(&u.Like,&u.Collect,&u.Coin,address)
	if err!=nil{
		fmt.Printf("inmovie woring:%d\n",err)
		return
	}
	c.JSON(200,gin.H{
		"code":1,
		"number":m.Number,
		"like":u.Like,
		"coin":u.Coin,
		"collect":u.Collect,
		"address":address,
	})
	fmt.Printf("%d",m.Number)
}
//点赞
func Like(c *gin.Context){
	sqlStr:="select like, collect, coin from movie1 where number=?"
	var u models.Info
	number:=c.Query("number")
	err:=dao.DB.QueryRow(sqlStr,number).Scan(&u.Like,&u.Collect,&u.Coin)
	sqlstr:="update movie1 set like=? where number=?"
	dao.DB.Exec(sqlstr,u.Like+1,number)
	if err!=nil{
		fmt.Printf("san fail like :%d\n",err)
		return
	}
	dao.DB.QueryRow(sqlStr,number).Scan(&u.Like,&u.Collect,&u.Coin)
	c.JSON(200,gin.H{
		"code":1,
		"number":number,
		"coin":u.Coin,
		"like":u.Like,
		"collect":u.Collect,
	})
}
//收藏
func Collect(c *gin.Context){
	sqlStr:="select like, collect, coin from movie1 where number=?"
	var u models.Info
	number:=c.Query("number")
	err:=dao.DB.QueryRow(sqlStr,number).Scan(&u.Like,&u.Collect,&u.Coin)
	sqlstr:="update movie1 set collect=? where number=?"
	dao.DB.Exec(sqlstr,u.Collect+1,number)
	if err!=nil{
		fmt.Printf("san fail like :%d\n",err)
		return
	}
	dao.DB.QueryRow(sqlStr,number).Scan(&u.Like,&u.Collect,&u.Coin)
	c.JSON(200,gin.H{
		"code":1,
		"number":number,
		"coin":u.Coin,
		"like":u.Like,
		"collect":u.Collect,
	})
}
//投币
func Coin(c *gin.Context){
	sqlStr:="select like, collect, coin from movie1 where number=?"
	var u models.Info
	number:=c.Query("number")
	err:=dao.DB.QueryRow(sqlStr,number).Scan(&u.Like,&u.Collect,&u.Coin)
	sqlstr:="update movie1 set coin=? where number=?"
	dao.DB.Exec(sqlstr,u.Coin+1,number)
	if err!=nil{
		fmt.Printf("san fail like :%d\n",err)
		return
	}
	dao.DB.QueryRow(sqlStr,number).Scan(&u.Like,&u.Collect,&u.Coin)
	c.JSON(200,gin.H{
		"code":1,
		"number":number,
		"coin":u.Coin,
		"like":u.Like,
		"collect":u.Collect,
	})
}
//三连
func Triple(c *gin.Context){
	sqlStr:="select like, collect, coin from movie1 where number=?"
	var u models.Info
	number:=c.Query("number")
	err:=dao.DB.QueryRow(sqlStr,number).Scan(&u.Like,&u.Collect,&u.Coin)
	sqlstr:="update movie set like=? , set coin=? ,set collect=?,where number=?"
	dao.DB.Exec(sqlstr,u.Like+1,u.Coin+1,u.Collect+1,number)
	if err!=nil{
		fmt.Printf("san fail like :%d\n",err)
		return
	}
	dao.DB.QueryRow(sqlStr,number).Scan(&u.Like,&u.Collect,&u.Coin)
	c.JSON(200,gin.H{
		"code":1,
		"number":number,
		"coin":u.Coin,
		"like":u.Like,
		"collect":u.Collect,
	})
}
//留言
func GiveMessage(c *gin.Context){
	//当前时间
	time:=time.Now()
	//读取视频的编号
	var m models.Movie1
	number:=c.ShouldBind(&m)
	//读取用户名
	var a models.Movie
	var n string//用户名
	err1:=c.ShouldBind(&a)
	if err1!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err1.Error(),
		})
	}
	sqlstr:="select name from user where id=?"
	err:=dao.DB.QueryRow(sqlstr,a.Id).Scan(n)
	if err!=nil{
		fmt.Printf("scan fail,err:=%d\n",err)
	}
	//留言储存
	message:=c.PostForm("message")
	//存入数据库
	sqlstr1:="insert into message(name,time,number,message)values (?,?,?,?)"
	dao.DB.Exec(sqlstr1,n,time,number,message)
}
//留言展示
func Allmessage(c *gin.Context){
	var b models.Movie1
	if err:=c.ShouldBind(&b);err!=nil{
		fmt.Printf("Allmessage woring %d\n",err)
		return
	}
	sqlstr:="select name,message,time form message where number=?"
	rows,err:=dao.DB.Query(sqlstr,b.Number)
	if err!=nil{
		fmt.Printf("query woring (allmessage)%d\n",err)
	}
	defer rows.Close()
	for rows.Next(){
		var m models.Message
		err:=rows.Scan(&m.Name,&m.Message,&m.Time)
		if err!=nil{
			fmt.Printf(" allmeeage scan fail:%d\n",err)
			return
		}
		c.JSON(200,gin.H{
			"name":m.Name,
			"message":m.Message,
			"time":m.Time,
		})
	}
}