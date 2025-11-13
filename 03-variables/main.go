package main

import "fmt"

const (
	APP_NAME    = "Belajar Go"
	APP_VERSION = "1.0.0"
	Max_USERS   = 1000
)

func main() {
	fmt.Println("=================================")
	fmt.Println("====Belajar Varibel Dasar GO=====")
	fmt.Println("=================================")
	fmt.Println()

	demoVariable()
	demoConstant()
}

func demoVariable() {
	fmt.Println("1. Deklarasi Variabel")
	fmt.Println("---------------------")

	// Cara 1: Explisit -- Type Declaration
	var nama string = "Eko"
	var umur int = 30
	var tinggi float64 = 1.75

	fmt.Printf("Nama   : %s\n", nama)
	fmt.Printf("Umur   : %d\n", umur)
	fmt.Printf("Tinggi : %.2f\n", tinggi)

	// Cara 2: Implicit --Type Inference (Go guessing the data type)
	var alamat = "Jakarta"
	var berat = 70.5

	fmt.Printf("Alamat : %s\n", alamat)
	fmt.Printf("Berat  : %.2f\n", berat)

	// Cara 3: Short Declaration -- Only can be used inside functions
	hobi := "Membaca"
	jumlahSaudara := 3

	fmt.Printf("Hobi          : %s\n", hobi)
	fmt.Printf("Jumlah Saudara: %d\n", jumlahSaudara)

	fmt.Println()
	fmt.Println("2. Deklarasi Multiple Variabel")
	fmt.Println("-------------------------------")

	var ( 	// Deklarasi multiple variabel dengan tipe data yang sama
		kota    string = "Bandung"
		provinsi string = "Jawa Barat"
		negara   string = "Indonesia"
	)

	fmt.Printf("Kota    : %s\n", kota)
	fmt.Printf("Provinsi: %s\n", provinsi)
	fmt.Printf("Negara  : %s\n", negara)

	// Deklarasi multiple variabel dengan short declaration
	hobi1, hobi2, hobi3 := "Bermain Musik", "Olahraga", "Traveling"

	fmt.Printf("Hobi 1: %s\n", hobi1)
	fmt.Printf("Hobi 2: %s\n", hobi2)
	fmt.Printf("Hobi 3: %s\n", hobi3)
	fmt.Println()

}

func demoConstant() {
	fmt.Println("3. Menggunakan Konstanta")
	fmt.Println("------------------------")

	const phi = 3.14159
	const dayOfWeek = 7

	fmt.Println("Nilai Konstanta:")
	fmt.Printf("Nama Aplikasi : %s\n", APP_NAME)
	fmt.Printf("Versi Aplikasi: %s\n", APP_VERSION)
	fmt.Printf("Maksimal User : %d\n", Max_USERS)
	fmt.Printf("Phi           : %.5f\n", phi)
	fmt.Printf("Hari dalam Seminggu: %d\n", dayOfWeek)
	fmt.Println()

	// Menghitung luas lingkaran menggunakan konstanta
	var radius float64 = 7.0
	luasLingkaran := phi * radius * radius
	fmt.Printf("Luas Lingkaran dengan jari-jari %.2f adalah %.2f\n", radius, luasLingkaran)
}