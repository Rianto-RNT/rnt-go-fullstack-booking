package main

import "fmt"

// 4)
type User struct {
	NamaDepan string
	NamaBelakang string
}

func main() {
	// // 1)
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// // 2)
	// stringStuf := []string{"keyboard", "watch", "iPad", "phone", "hdd", "wallet", "cirgaret"}
		
	// for _, x := range stringStuf {
	// 	fmt.Println(x)
	// }

	// // 3)
	// mapsStuf := make(map[string]string)
	// mapsStuf ["keyboards"] = "keyboard"
	// mapsStuf ["watches"] = "watch"
	// mapsStuf ["earphones"] = "earphone"

	// for i, x := range mapsStuf {
	// 	fmt.Println(i, x)
	// }

	// 4)
	var sliceKu []User

	usrOne := User{
		NamaDepan: "Rianto",
	}

	usrTwo := User{
		NamaDepan: "Morrison",
	}

	sliceKu= append(sliceKu, usrOne)
	sliceKu= append(sliceKu, usrTwo)

	for i, x := range sliceKu {
		fmt.Println(i, x)
	}
}