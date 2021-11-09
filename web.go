package main

/*
	Usage:
		webRouter, httpRouter := Web.Router()
			webRouter. {
				Domains(domains...string)
				Methods(methods...string)
				PathPrefix(string)
				SupportParameters(parameter)
				Route([]Web.RouteConfig)
				RouteByController([]string, BaseController)
			}
*/

import (
	"fmt"
	"net/http"
	"time"
	"github.com/ahjiat/gomvc/controller"
	"github.com/ahjiat/gomvc/parameter"
	"github.com/ahjiat/gomvclib"
)

type RC = Web.RouteConfig

func main() {
	webRouter, httpRouter := Web.Router()
	webRouter = webRouter.SetViewDir("view").SetControllerDir("controller")

	routeDomain01 := webRouter.Domains("test.grannygame.io")
	routeDomain01.Route(RC{"/mytest", "Index"}, new(controller.Test))

	route02 := webRouter.SupportParameters(
		new(parameter.Username),
		new(parameter.Password))
	route02.Route([]RC{
		{Path:"/testinfo", Action:"Info"},
		{Path:"/testinfo2", Action:"Info2"},
	}, new(controller.Info))
	route02.Route([]RC{
		{"/testgetpost", "TestGetPost"},
		{"/matchtest/{n:.*}", "Index"},
	}, new(controller.Motor))
	route02.RouteByController("/allinone", new(controller.Motor))
	route02.RouteByController("/info/{path}", new(controller.Info))
	route02.Route(RC{"/{n:.*}", "PageNotFound"}, new(controller.Page))

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        httpRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog: nil,
	}
	fmt.Println("Http Server...")
	server.ListenAndServe()
}
