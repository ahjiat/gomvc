package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
)

type Login struct{ BaseController }
func (self *Login) Check(args struct {
}) {
	header := self.Base.ParseTemplate("Navigator", `Top {{.}}`)
	footer := self.Base.ParseTemplate("Footor", `Bottom {{.}}`)
	v := self.Base.SetMasterView()
	v.AddTemplate("header", header)
	v.AddTemplate("footer", footer)
	self.Base.RouteNext("welcome", "20")
}
