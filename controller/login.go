package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
)

type Login struct{ BaseController }
func (self *Login) Check(args struct {
	GET_name string
}) {
	headerDat := "happyday"; _ = headerDat
	if args.GET_name == "check" {
		v := self.Base.CreateMasterView(); _ = v
		v.DefineTemplate("header", headerDat, "Header.html")
	} else if args.GET_name == "check1" {
		self.Base.CreateMasterView("Check1.html")
	} else {
		self.Base.RemoveMasterView()
	}

	self.Base.RouteNext("welcome1", "20")
}

func (self *Login) Check2() {
	self.Base.RouteNext("welcome2", "22")
}

