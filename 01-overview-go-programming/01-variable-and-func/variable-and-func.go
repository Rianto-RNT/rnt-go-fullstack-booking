package main

import "fmt"

// // Assign string with 1 variable
// func saySomething(s string) string {
// 	return s
// }

// Assign string with 2 variable
func saySomething(s string) (string, string) {
	return s,"Hardcode saja"
}

func main() {
	var ketik string
	var ketikAngka int
	var ketikTulisan string

	ketik, _ = saySomething("Empat kali empat")
	ketikAngka = 11
	ketikTulisan, _ = saySomething("Tulisannya ini? atau air minum dalam kemasan 1550ml")
	
	fmt.Println(ketik)
	fmt.Println(ketikAngka)
	fmt.Println(ketikTulisan)
}
