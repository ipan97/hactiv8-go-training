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
		input, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Maaf terjadi kesalahan : ", err)
		}
		var selectedData Biodata
		if input > 0 && input <= len(data) {
			selectedData = findSelected(data, input)
			printData(selectedData)
		} else {
			fmt.Println("Maaf data yang anda pilih tidak ditemukan.")
		}
	} else {
		fmt.Println("Maaf, Harap masukan pilihan anda")
	}
}

func findSelected(data []Biodata, input int) Biodata {
	if len(data) > 0 {
		return data[input-1]
	}
	return Biodata{}
}

func printData(data Biodata) {
	fmt.Println("Nama : ", data.Nama)
	fmt.Println("Alamat : ", data.Alamat)
	fmt.Println("Pekerjaan : ", data.Pekerjaaan)
	fmt.Println("Alasan Memilih Kelas Golang : ", data.AlasanMemilihKelasGolang)
}
