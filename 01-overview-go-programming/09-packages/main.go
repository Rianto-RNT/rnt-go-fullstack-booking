package main

import (
	"fmt"
	"rnt-go-booking/01-overview-go-programming/09-packages/helpers"
)

func main() {
	var varKu helpers.TypeKu
	varKu.NamaType = "stuff"
	fmt.Println(varKu.NamaType)
}