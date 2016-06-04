package routers

import (
	"github.com/gowebf/goslim"
	"github.com/gowebf/goslim/Example/controllers"
)

func RegistRout(r *goslim.Router) {

	// dc := &controllers.DefaultController{}
	r.Get("/", controllers.IndexPage)
	r.Get("/modeltest", controllers.TestModel)
	r.Post("/post/:kkk", controllers.PostTest)
	r.Get("/user/:name", controllers.DefaultHi)

	rc := &controllers.RestController{}
	r.REST("/User", rc)

}
