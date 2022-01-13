package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
	_"encoding/json"
	_"text/template"
	"github.com/ahjiat/gomvc/parameter"
)


type Test struct{ BaseController }
func (self *Test) Index(arg struct {
	GET_name    parameter.Username
}) {
	js := struct {
		Title string	`json:"title"`
		Age int			`json:"age"`
	} { "welcome", 40 };
	html := struct {
		Name string		`json:"name"`
		Num	int			`json:"num"`
	} { "power", 20 }

	self.Base.ViewBag = struct {
		Js interface{}
		Html interface{}
	} {
		js,
		html,
	}

	//if len(self.Base.ChainArgs) > 0 {
	//	self.Base.Echo(self.Base.ChainArgs[0].(string))
	//}
	//self.Base.Echo(arg.GET_name.Value)

	self.masterView(self.Base.ViewBag)
	self.Base.View()
}
func (self *Test) masterView(data interface{}, fileName ...string) {
	self.Base.MasterViewBag = data
	self.Base.MasterView(fileName...)
}

