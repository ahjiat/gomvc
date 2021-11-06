package controller

import (
	. "github.com/ahjiat/gomvclib/basecontroller"
)

type Page struct{ BaseController }

func (self *Page) PageNotFound() {
	self.Echo("Page Not Found, but status code still 200")	
}
