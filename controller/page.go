package controller

import (
	. "github.com/ahjiat/gomvclib/basecontroller"
)

type Page struct{ BaseController }

func (self *Page) PageNotFound() {
	self.Base.Echo("Page Not Found, but status code still 200")	
}
