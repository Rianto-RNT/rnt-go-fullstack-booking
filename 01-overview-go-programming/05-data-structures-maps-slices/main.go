package main

import (
	"fmt"
	"log"
)

// // 3) /*MAPS II*/
// 	type Pengguna struct {
// 		NamaDepan string
// 		NamaBelakang string
// 	}



func main() {
	// 1) /*STRING MAPS*/
	// mapsKu := make(map[string]string)
	// // first map
	// mapsKu["kuda"] = "jarn"
	// // replace first map
	// mapsKu["kuda"] = "fido"
	
	// mapsKu["kuda-lainnya"] = "arn"
	
	
	// fmt.Println(mapsKu["kuda"])
	// fmt.Println(mapsKu["kuda-lainnya"])

	// mapsKu := make(map[string]string)

	//------------------------------------------

	// 2) /*INT MAPS*/
	// mapsKu := make(map[string]int)

	// mapsKu["satu"] = 22
	// mapsKu["dua"] = 89

	// fmt.Println(mapsKu["satu"], mapsKu["dua"])

	//------------------------------------------

	// 3) /*MAPS I*/
	// mapsKu := make(map[string]Pengguna)

	// user := Pengguna {
	// 	NamaDepan: "Rianto",
	// 	NamaBelakang: "RNT",
	// }

	// mapsKu["usr"] = user

	// fmt.Println(mapsKu["usr"].NamaDepan)

	// var varBaruKu float32

	// varBaruKu = 125.168

	// fmt.Println("result:",varBaruKu)

	//------------------------------------------

	// // 4) /*SLICE*/
	// // play with character
	// var stringKuHuruf []string

	// stringKuHuruf = append(stringKuHuruf, "Rianto")
	// stringKuHuruf = append(stringKuHuruf, "Rian")
	// stringKuHuruf = append(stringKuHuruf, "ian")

	// log.Println(stringKuHuruf)
	
	// // play with number
	// var stringKuAngka []int
	
	// stringKuAngka = append(stringKuAngka, 223)
	// stringKuAngka = append(stringKuAngka, 4)
	// stringKuAngka = append(stringKuAngka, 88)
	
	// sort.Ints(stringKuAngka) // asc = sort by low to high value
	// log.Println(stringKuAngka)

	// // 5) /*SLICES sorthand*/
	nomor := []int{1,2,3,4,55,6,7,83,9,10}
	log.Println(nomor)

	// range slice of int
	fmt.Println(nomor[2:4])
	fmt.Println(nomor[2:7])
	fmt.Println(nomor[4:9])

	kataKalimat := []string{"random-aja", "seratus-ribu-rupiah", "satu", "makan-soto", "delapan", "anak-anak", "ibu", "Indonesia"}
	fmt.Println(kataKalimat)

	//range slice of string
	fmt.Println(kataKalimat[3:6])
}