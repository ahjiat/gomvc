package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
	_"encoding/json"
	_"html/template"
)

type Test struct{ BaseController }
func (self Test) Index() {
	js := struct {
		Title string	`json:"title"`
		Age int			`json:"age"`
	} { "welcome", 40 };
	html := struct {
		Name string		`json:"name"`
		Num	int			`json:"num"`
	} { "power", 20 }
	self.ViewBag = struct {
		Js interface{}
		Html interface{}
	} {
		js,
		html,
	}
	self.View()
}
