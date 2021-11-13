package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
)

type Login struct{ BaseController }
func (self *Login) Check(args struct {
}) {
	header := self.Base.ParseTemplate("Navigator", `Top {{.}}`)
	footer := self.Base.ParseTemplate("Footor", `Bottom {{.}}`)
	v := self.Base.CreateMasterView()
	v.DefineTemplate("header", header)
	v.DefineTemplate("footer", footer)
	self.Base.RouteNext("welcome1", "20")
}

func (self *Login) Check2() {
	footer := self.Base.ParseTemplate("Footor2", `Bottom {{.}}`)
	v := self.Base.GetMasterView()
	v.DefineTemplate("footer", footer)
	self.Base.RouteNext("welcome2", "22")
}

