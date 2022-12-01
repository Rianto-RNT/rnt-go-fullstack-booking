package main

import (
	"errors"
	"fmt"
)

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Caonnot divide by 0")
	}

	result = x / y
	return result, nil
}

func main() {
	result, err := divide(100.0, 5) // change value here
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("result of division is", result)
}