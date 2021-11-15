package controller
import (
	. "github.com/ahjiat/gomvclib/basecontroller"
	_"encoding/json"
	_"text/template"
	_"github.com/ahjiat/gomvc/parameter"
)


type Template struct{ BaseController }
func (self *Template) Index() {
	self.Base.View()
	/*
	header := `{{define "Header"}}<div>Header:{{.}}</div>{{end}}`; _ = header
	body := `{{define "Body"}}<div>Body:{{.}}</div>{{end}}`
	footer := `{{define "Footer"}}<div>Footer:{{.}}</div>{{end}}`
	page := `
{{define "Page"}}
<!DOCTYPE html>
<html>
<body>
	{{template "Header" .User}}
	{{template "Body" .Name}}
	{{template "Footer" .Value}}
</body>
</html>
{{end}}`
	body2 := `{{define "Body"}}<div>Body2:{{.}}</div>{{end}}`

	t := template.New("webpage")
	t.Parse(body)
	t.Parse(footer)
	t.Parse(page)
	//t.Parse(header)
	t.Parse(body2)
	type S1 struct { N1 string; N2 string }
	data := struct {
		User S1
		Name string
		Value string
	} {S1{"N1", "N2"}, "Name", "Value"}; _ = data
	/*
	err := t.ExecuteTemplate(self.Base.Response, "Page", nil)
	if err != nil {
		panic("2nd template failed")
	}
	self.Base.Echo(t.DefinedTemplates())
	*/
}

