package main

import (
	"fmt"
	"time"
)

var explain string = "Methods are functions that are associated with values of a particular type."

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
	fmt.Println(explain)
}
