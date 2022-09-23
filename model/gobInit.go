package model

import "encoding/gob"

func GobInit() {
	gob.Register(Employee{})
}
