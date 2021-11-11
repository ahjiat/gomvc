package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
	_"fmt"
)

type Login struct{ BaseController }
func (self *Login) Check() {
	self.Base.Echo("from middleware\n")
	self.Base.Next("welcome", "20")
}
