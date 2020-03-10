// aaa project main.go
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)

const MasterDSName = "root:123456@tcp(127.0.0.1:3306)/?charset=utf8"

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	app.Run(iris.Addr(":8080"))

}
