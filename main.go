package main

import (
	"github.com/ken-aio/go-echo-xorm/app/routes"
	_ "github.com/ken-aio/go-echo-xorm/docs"
)

// @title Go Echo SQLBoiler sample project
// @version 1.0
// @description This is a sample server.

// @contact.name API Support
// @contact.url dummy
// @contact.email hoge

// @license.name ken-aio
// @license.url dummmy

// @host localhost:1314
// @BasePath /
func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1313"))
}
