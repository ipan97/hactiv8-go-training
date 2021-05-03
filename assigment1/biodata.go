package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama                     string
	Alamat                   string
	Pekerjaaan               string
	AlasanMemilihKelasGolang string
}

func main() {
	data := []Biodata{
		{
			Nama:                     "Ahmad",
			Alamat:                   "Caringin Bandung",
			Pekerjaaan:               "Backend Developer",
			AlasanMemilihKelasGolang: "Karena saya ingin belajar golang",
		},
		{
			Nama:                     "Thomas Andrianto",
			Alamat:                   "Sukabumi",
			Pekerjaaan:               "Backend Developer",
			AlasanMemilihKelasGolang: "Karena saya ingin belajar golang",
		},
		{
			Nama:                     "Dewo",
			Alamat:                   "Yogyakarta",
			Pekerjaaan:               "Backend Developer",
			AlasanMemilihKelasGolang: "Karna saya ingin belajar golang",
		},
		{
			Nama:                     "Ipan Taupik Rahman",
			Alamat:                   "Ciamis",
			Pekerjaaan:               "Backend Developer",
			AlasanMemilihKelasGolang: "Karna saya ingin belajar golang",
		},
	}
	if len(os.Args) > 1 {
		pilih, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Maaf terjadi kesalahan : ", err)
		}
		var selectedData Biodata
		if pilih > 0 && pilih <= len(data) {
			selectedData = data[pilih-1]
			fmt.Println("Nama : ", selectedData.Nama)
			fmt.Println("Alamat : ", selectedData.Alamat)
			fmt.Println("Pekerjaan : ", selectedData.Pekerjaaan)
			fmt.Println("Alasan Memilih Kelas Golang : ", selectedData.AlasanMemilihKelasGolang)
		} else {
			fmt.Println("Maaf data yang anda pilih tidak ditemukan.")
		}
	} else {
		fmt.Println("Maaf, Harap masukan pilihan anda")
	}
}
