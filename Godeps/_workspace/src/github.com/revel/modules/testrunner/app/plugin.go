package app

import (
	"ForeginWWW/Godeps/_workspace/src/github.com/revel/revel"
	"fmt"
)

func init() {
	revel.OnAppStart(func() {
		fmt.Println("Go to /@tests to run the tests.")
	})
}
