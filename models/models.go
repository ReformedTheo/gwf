package models

import (
	"fmt"
	//"GWF/orm"
)

func models() {
	fmt.Print("HALLO")
	type user struct {
		name     string
		age      int
		is_admin bool
	}
}
