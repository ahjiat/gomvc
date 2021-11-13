package main

/*
	Usage:
		webRouter, httpRouter := Web.Router()
			webRouter. {
				Domains(domains...string)
				Methods(methods...string)
				PathPrefix(string)
				SupportParameters(parameter)
				RouteChain([]string|string, BaseController)
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

func domainRoute(route *Web.Route) {
	route.Route(C{"/super", "Index"}, new(controller.Test))
	route.Route(C{"/template", "Index"}, new(controller.Template))
}

func loginRoute(route *Web.Route) {
	route.Route(C{"/login/{n:.*}", "Index"}, new(controller.Test))
}

func defaultRoute(route *Web.Route) {
	route.Route([]C{
		{Path:"/testinfo", Action:"Info"},
		{Path:"/testinfo2", Action:"Info2"},
	}, new(controller.Info))
	route.Route([]C{
		{"/testgetpost", "TestGetPost"},
		{"/matchtest/{n:.*}", "Index"},
	}, new(controller.Motor))
	route.RouteByController("/allinone", new(controller.Motor))
	route.RouteByController("/info/{path}", new(controller.Info))
	route.Route(C{"/{n:.*}", "PageNotFound"}, new(controller.Page))
}

func main() {
	webRouter, httpRouter := Web.Router()
	webRouter = webRouter.SetViewDir("view").SetControllerDir("controller")
	webRouter = webRouter.SupportParameters(new(parameter.Username), new(parameter.Password))

	domainRoute(webRouter.Domains("test.grannygame.io"))
	loginRoute(webRouter.RouteChain("Check", new(controller.Login)).RouteChain("Check2", new(controller.Login)))
	defaultRoute(webRouter)

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
