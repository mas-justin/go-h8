package main

import "fmt"

type Profil struct {
	Nama string
	Umur int    
	Alamat Alamat 
	Pekerjaan string 
	Agama string 
}

type Alamat struct {
	Jalan string
	Kota string 
	Provinsi string
}

func main() {
	fmt.Println("Hello World!")
	
	profil := Profil{
		Nama: "Reydika",
		Umur: 31,
		Alamat: Alamat{
			Jalan: "Terminal Ragunan",
			Kota: "Jakarta Selatan",
			Provinsi: "DKI Jakarta",
		},
		Pekerjaan: "Programmer",
		Agama: "Islam",
	}

	fmt.Println(profil)
}