package main

/*
	Usage:
		webRouter, httpRouter := Web.Router()
			webRouter. {
				Domains(domains...string)
				Methods(methods...string)
				PathPrefix(string)
				SupportParameters(parameter)
				RouteChain([]Web.RouteChainConfig)
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

type C = Web.RouteConfig
type RC = Web.RouteChainConfig

func main() {
	webRouter, httpRouter := Web.Router()
	webRouter = webRouter.SetViewDir("view").SetControllerDir("controller")
	webRouter = webRouter.SupportParameters(new(parameter.Username), new(parameter.Password))

	domain1 := webRouter.Domains("test.grannygame.io")
	{
		c := domain1.RouteChain(RC{"Check", new(controller.Login)})
		c.Route(C{"/super", "Index"}, new(controller.Test))
	}
	domain1.Route(C{"/mytest", "Index"}, new(controller.Test))

	route02 := webRouter
	route02.Route([]C{
		{Path:"/testinfo", Action:"Info"},
		{Path:"/testinfo2", Action:"Info2"},
	}, new(controller.Info))
	route02.Route([]C{
		{"/testgetpost", "TestGetPost"},
		{"/matchtest/{n:.*}", "Index"},
	}, new(controller.Motor))
	route02.RouteByController("/allinone", new(controller.Motor))
	route02.RouteByController("/info/{path}", new(controller.Info))
	route02.Route(C{"/{n:.*}", "PageNotFound"}, new(controller.Page))

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
