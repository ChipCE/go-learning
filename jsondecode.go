package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

type SampleData struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}

func GenericPostHandle(ctx iris.Context) {
	var request map[string]interface{}
	err := ctx.ReadJSON(&request)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.Writef("Received: %#+v\n", request)
	fmt.Println(request)
}

func PostHandleFunc(ctx iris.Context) {
	var sampleData SampleData
	err := ctx.ReadJSON(&sampleData)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.Writef("Received: %#+v\n", sampleData)
	fmt.Println(sampleData)

}

func main() {
	app := iris.New()
	// Load all templates from the "./views" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package.
	app.RegisterView(iris.HTML("./views", ".html"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Post("/", PostHandleFunc)
	app.Post("/generic", GenericPostHandle)

	// Method:    GET
	// Resource:  http://localhost:8080/user/42
	//
	// Need to use a custom regexp instead?
	// Easy;
	// Just mark the parameter's type to 'string'
	// which accepts anything and make use of
	// its `regexp` macro function, i.e:
	// app.Get("/user/{id:string regexp(^[0-9]+$)}")
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	// Start the server using a network address.
	app.Listen(":8080")
}
