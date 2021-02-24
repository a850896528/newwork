package main

import (
	"Hldwork/dao"
	"Hldwork/routers"
	"fmt"
)

func main() {
	err:=dao.InitDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	defer dao.Close()
	//模型绑定

	//注册路由
	r:=routers.SetupRouter()
	r.Run(":9090")//启动server
}
