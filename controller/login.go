package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
	_"fmt"
)

type Login struct{ BaseController }
func (self *Login) Check() {
	self.Base.Echo("from Check middleware\n")
	self.Base.RouteNext("welcome", "20")
}
