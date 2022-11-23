package main

import "fmt"

func rubahKePointer(s *string) {
	valueBaru := "Putih"
	*s = valueBaru
	fmt.Printf("s di set ke: %v\n", s)
}


func main() {
	var stringHuruf string

	stringHuruf = "Hitam"

	fmt.Println("String saya di set ke", stringHuruf)

	// panngil func rubah ke pointer
	rubahKePointer(&stringHuruf)
	fmt.Println("Setelah func rubahKePointer dipanggil maka stringHuruf di set ke:", stringHuruf)	
}