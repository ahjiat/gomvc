package controller

import (
	. "github.com/ahjiat/gomvclib/basecontroller"
	"github.com/ahjiat/gomvc/database"
	"fmt"
)

type Template struct{ BaseController }

func (self *Template) Index() {
	type User struct {
		Id int			`gorm:"column:id"`
		Value01 string	`gorm:"column:value01"`
		Value02 string	`gorm:"column:value02"`
	}
	users := []User{}; _ = users
	Database.AdminSession().Table("si").Scan(&users)
	for _, u := range users {
		m := fmt.Sprintf("%d, %v, %s\n", u.Id, u.Value01, u.Value02)
		self.Base.Echo(m)
	}
	self.Base.View()
}

func (self *Template) SetAry(arg struct {
	POST_data	[]string
	POST_name	[]int
}) {
	for i, v := range arg.POST_data {
		self.Base.Echo(fmt.Sprintf("%v => %v\n", i, v))
	}
	fmt.Println(arg.POST_name)
}
