package main

import (
	"github.com/gowebf/goslim"
	"github.com/gowebf/goslim/Example/routers"
	"github.com/gowebf/goslim/middleware"
)

func main() {
	g := goslim.New()

	routers.RegistRout(g.Router)

	g.Router.Use(middleware.LogMiddleware)

	g.Run()
}
