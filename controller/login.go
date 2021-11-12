package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
)

type Login struct{ BaseController }
func (self *Login) Check(args struct {
	GET_name int
}) {
	if args.GET_name == 20 {
		self.Base.Echo("name not found\n")
	} else {
		self.Base.RouteNext("welcome", "20")
	}
}
