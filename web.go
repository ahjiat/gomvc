package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/ahjiat/gomvc/controller"
	"github.com/ahjiat/gomvc/parameter"
	Web "github.com/ahjiat/gomvclib"
)

type RC = Web.RouteConfig

func main() {
	webRouter, httpRouter := Web.Router()
	root := webRouter.SupportParameters(
		new(parameter.Username),
		new(parameter.Password))

	ipDomain := root.Domains("52.77.146.102")
	ipDomain.Route([]RC{
		{"/testinfo", "Info"},
		{"/testinfo2", "Info2"},
	}, new(controller.Info))
	ipDomain.Route([]RC{
		{"/testgetpost", "TestGetPost"},
		{"/matchtest/{n:.*}", "Index"},
	}, new(controller.Motor))

	root.RouteByController("/allinone", new(controller.Motor))
	root.RouteByController("/info/{path}", new(controller.Info))

	root.Route(RC{"/{n:.*}", "PageNotFound"}, new(controller.Page))

	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        httpRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Http Server...")
	server.ListenAndServe()
}
