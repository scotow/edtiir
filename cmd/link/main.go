package main

import (
	"fmt"
	"github.com/scotow/edtiir"
)

func main() {
	year := edtiir.Year2018
	fmt.Println(year.CurrentWeek().Link())
}
