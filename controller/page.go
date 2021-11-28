package controller

import (
	. "github.com/ahjiat/gomvclib/basecontroller"
	"net/http"
)

type Page struct{ BaseController }

func (self *Page) PageNotFound() {
	self.Base.Response.WriteHeader(http.StatusNotFound)
	self.Base.Echo("Page Not Found")
}
