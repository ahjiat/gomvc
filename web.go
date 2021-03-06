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
	"context"
	"fmt"
	"net"
	"net/http"
	"syscall"
	"time"
	"log"
	"os"

	"github.com/ahjiat/gomvc/controller"
	"github.com/ahjiat/gomvc/parameter"
	"github.com/ahjiat/gomvclib"
	"github.com/gorilla/mux"
	"golang.org/x/sys/unix"
)

type C = Web.RouteConfig

func domainRoute(route *Web.Route) {
	route.Route(C{"/super", "Index"}, new(controller.Test))
}

func loginRoute(route *Web.Route) {
	route.Route(C{"/login/{n:.*}", "Index"}, new(controller.Test))
}

func defaultRoute(route *Web.Route) {
	route.Route(C{"/template", "Index"}, new(controller.Template))
	route.Route(C{"/api/process", "SetAry"}, new(controller.Template))
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
	// enable log
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666); if err != nil { panic(err) }
	defer logFile.Close()
	log.SetOutput(logFile)

	webRouter, httpRouter := Web.Router()
	webRouter = webRouter.SetViewDir("view")
	webRouter = webRouter.SupportParameters(new(parameter.Username), new(parameter.Password))

	domainRoute(webRouter.Domains("test.grannygame.io"))
	loginRoute(webRouter.Use("Check", new(controller.Login)).Use("Check2", new(controller.Login)))
	defaultRoute(webRouter)

	httpRouter.Use(RequestLogger)
	normal_webserver(httpRouter)
	//reuseport_webserver(httpRouter)
}

func normal_webserver(r *mux.Router) {
	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog: nil,
	}
	fmt.Println("normal_webserver...")
	err := server.ListenAndServe(); if err != nil { panic(err) }
}

func reuseport_webserver(r *mux.Router) {
	lc := net.ListenConfig{
        Control: func(network, address string, conn syscall.RawConn) error {
            var operr error
            if err := conn.Control(func(fd uintptr) {
                operr = syscall.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
            }); err != nil {
                return err
            }
            return operr
        },
    }

    ln, err := lc.Listen(context.Background(), "tcp", "0.0.0.0:8080")
    if err != nil {
        panic(err)
    }

	server := &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog: nil,
	}

	fmt.Println("reuseport_webserver...")
	err = server.Serve(ln); if err != nil { panic(err) }
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		targetMux.ServeHTTP(w, r)
		requesterIP := r.RemoteAddr
		log.Printf(
			 "%s\t\t%s\t\t%s\t\t%v",
			 r.Method,
			 r.RequestURI,
			 requesterIP,
			 time.Since(start),
		)
	})
 }

