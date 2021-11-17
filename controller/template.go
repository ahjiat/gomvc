package controller

import (
	. "github.com/ahjiat/gomvclib/basecontroller"
)

type Template struct{ BaseController }

func (self *Template) Index() {
	self.Base.View()
}
