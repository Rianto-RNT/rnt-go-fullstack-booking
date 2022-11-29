package main

import (
	"encoding/json"
	"fmt"
)

type Hero struct {
	NamaDepan    string `json:"nama_depan"`
	NamaBelakang string `json:"nama_belakang"`
	Departement  string `json:"departement"`
	Role         string `json:"role"`
	Active       string `json:"active"`
}

func main() {
	jsonKu := `
	[
		{
			"nama_depan": "Ryan",
			"nama_belakang": "Morrison",
			"departement": "IT",
			"role": "Backend Developer",
			"active": true
		},
		{
			"nama_depan": "Rianto",
			"nama_belakang": "RNT",
			"departement": "IT",
			"role": "Data Engineer",
			"active": false
		}
	]`

	var unmarshalled []Hero

	err := json.Unmarshal([]byte(jsonKu), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalled json", err)
	}

	fmt.Printf("unmarshalled: %v\n", unmarshalled)
//==================================================

	// write json from struct
	var sliceKu []Hero

	var m1 Hero

	m1.NamaDepan = "rianto"
	m1.NamaBelakang = "RNT"
	m1.Departement = "IT"
	m1.Role = "Project Manager"
	// m1.Active = false

	sliceKu = append(sliceKu, m1)

	var m2 Hero

	m2.NamaDepan = "ryan"
	m2.NamaBelakang = "morrison"
	m2.Departement = "IT"
	m2.Role = "Backend Developer"
	// m2.Active = false

	sliceKu = append(sliceKu, m2)

	jsonBaru , err := json.MarshalIndent(sliceKu, "", "     ")
	if err != nil {
		fmt.Println("error marshaling >>>", err)
	}

	fmt.Println(string(jsonBaru))

}