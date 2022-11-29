package main

import (
	"fmt"
	"rnt-go-booking/01-overview-go-programming/10-channels/helpers"
)

// // 1)
// func PrintText(s string) {
// 	fmt.Println(s)
// }

// 2
const bankNomor = 100
// 2)
func KalkulasiNilai(channelKu chan int) {
	nomorAcak := helpers.NomorAcak(bankNomor)
	channelKu <- nomorAcak
}

func main() {
	// // 1)
	// PrintText("RNT")

	// 2)
	channelKu := make(chan int)
	defer close(channelKu)

	go KalkulasiNilai(channelKu)

	nmr := <- channelKu
	fmt.Println(nmr)
}