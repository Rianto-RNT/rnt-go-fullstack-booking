package main

import "fmt"

func main() {
	// // 1)
	// var isTrue bool

	// isTrue = true

	// if isTrue == true { // change value true or false
	// 	fmt.Println("isTrue is", isTrue)
	// } else {
	// 	fmt.Println("isTrue is", isTrue)
	// }

	// // 2)
	// cat:= "cat"
	// if cat == "cot" {
	// 	fmt.Println("cat is cat")
	// } else {
	// 	fmt.Println("cat is not cat")
	// }

	// // 3) else if
	// nomorKu := 100
	// isTrue := false

	// if nomorKu > 88 && !isTrue {
	// 	fmt.Println(" nomorKu lebih besar daripada 88 dan isTrue is set to true")
	// } else if nomorKu < 100 && isTrue {
	// 	fmt.Println("1)")
	// } else if nomorKu == 101 {
	// 	fmt.Println("2)")
	// } else if nomorKu > 1000 && isTrue == false {
	// 	fmt.Println("3")
	// }

	// 4) switch statement
	varKu := "HTML"

	switch varKu {
	case "Go":
		fmt.Println("Go is set to programming language")
	case "python":
		fmt.Println("phyton is set to programming language")
	case "java":
		fmt.Println("java is set to programming language")
	case "node.js":
		fmt.Println("node.js is set to programming language")
	case "HTML":
		fmt.Println("HTML is not a programming language")
	default:
		fmt.Println("programming is not available")
	}
}