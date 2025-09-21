package main

import (
	"fmt"

	"github.com/akrikkk/goconverting"
)

func main() {
	eng := "ghbdtn"
	rus := "привет"

	fmt.Println("To Russian:", goconverting.ToRussian(eng))
	fmt.Println("To English:", goconverting.ToEnglish(rus))
}
