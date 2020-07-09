package plugin

import (
	"fmt"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

func GetHandle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GET recv"})
}

func PostHandle(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "POST recv"})
}

func QueryGetHandle(ctx *gin.Context) {
	msg := ctx.Query("msg")
	option := ctx.Query("option")
	ctx.JSON(200, gin.H{"message": msg})
	fmt.Println(option)
	fmt.Println(ctx.Request.URL.Query()) //map[string][]string
}

func QueryPostHandle(ctx *gin.Context) {
	body := ctx.Request.Body
	value,err := ioutil.ReadAll(body)
	if err!= nil{
		fmt.Println(err.Error())
	}

	ctx.JSON(200, gin.H{"msg":string(value)})
}

func PathHandle(ctx *gin.Context) {
	id := ctx.Param("id")
	subid := ctx.Param("subid")
	option := ctx.Param("option")
	fmt.Println(id, subid, option)
	if option == ""{
		fmt.Println("option : null")
	} else {
		fmt.Println("option : ",option)
	}
	
	ctx.JSON(200, gin.H{
		"id":     id,
		"subid":  subid,
		"option": option})

}

func startTwo() {
	server := gin.Default()

	server.GET("/test", GetHandle)
	server.POST("/test", PostHandle)
	server.GET("/query", QueryGetHandle)
	server.POST("/query", QueryPostHandle)
	server.GET("/path/:id/:subid", PathHandle)

	server.Run(":8001")
}
