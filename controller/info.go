package controller

import (
	_ "strconv"
	"github.com/ahjiat/gomvc/parameter"
	. "github.com/ahjiat/gomvclib/basecontroller"
)

type Info struct{ BaseController }

func (self *Info) Info(arg struct {
	GET_name      *parameter.Username
	POST_name     *parameter.Username
	POST_password *parameter.Password
}) {
	self.Base.Echo("view path: " + self.Base.ViewBasePath + "\n")
	if arg.GET_name != nil {
		if arg.GET_name.Valid {
			self.Base.Echo("GET name => " + arg.GET_name.Value + "\n")
		} else {
			self.Base.Echo("Get name is invalid\n")
		}
	}
	if arg.POST_name != nil {
		if arg.POST_name.Valid {
			self.Base.Echo("POST name => " + arg.POST_name.Value + "\n")
		} else {
			self.Base.Echo("Get name is invalid\n")
		}
	}
}

func (self *Info) Info2(arg struct {
	GET_name  parameter.Username
	POST_name parameter.Username
}) {
	self.Base.Echo("GET name => " + arg.GET_name.Value + "\n")
	self.Base.Echo("POST name => " + arg.POST_name.Value + "\n")
}
