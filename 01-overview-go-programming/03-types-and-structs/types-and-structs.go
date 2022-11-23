package main

import (
	"fmt"
	"time"
)

type Pengguna struct {
	NamaDepan    string
	NamaBelakang string
	NomorTelpon  string
	Umur         int
	TanggalLahir time.Time
}

func main() {
	pengguna := Pengguna {
		NamaDepan: "Rianto",
		NamaBelakang: "RNT",
		NomorTelpon: "081xxxxxx",
		// "Umur" gak di spesifikasi
		// "Tanggal Lahir" gak di spesifikasi
	}

	fmt.Println("Keluaran:\n", pengguna.NamaDepan, pengguna.NamaBelakang, "Tanggal Lahir\n", pengguna.TanggalLahir)
}