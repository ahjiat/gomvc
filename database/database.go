package Database

import (
	exDB "github.com/ahjiat/gomvclib/database"
)

type BaseDB struct {
	*exDB.DB
}

func AdminSession() *BaseDB{
	return &BaseDB{exDB.GetSession("db")}
}
