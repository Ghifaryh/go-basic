package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
	IPK     float64
}

func main() {
	demonstrasiStructs()
}

func (m Mahasiswa) PrintInfoMahasiswa() {
	fmt.Println("Data Mahasiswa/i")
	fmt.Println("Nama:", m.Nama)
	fmt.Println("NIM:", m.NIM)
	fmt.Println("Jurusan:", m.Jurusan)
	fmt.Println("IPK:", m.IPK)
}

// Receiver
func (m Mahasiswa) GetFullInfo() string {
	return fmt.Sprintf("%s (NIM: %s) - %s, IPK: %.2f", m.Nama, m.NIM, m.Jurusan, m.IPK)
}

// Pointer
func (m *Mahasiswa) UpdateIPK(ipk float64) {
	m.IPK = ipk
}

func demonstrasiStructs() {
	fmt.Println("Struct")
	fmt.Println("------")

	mahasiswa := Mahasiswa{
		Nama:    "Gip",
		NIM:     "👍",
		Jurusan: "Informatika",
		IPK:     3.70,
	}

	mahasiswa.PrintInfoMahasiswa()
	println()
	fmt.Println("Info lengkap: ", mahasiswa.GetFullInfo())
	println()

	// Update IPK using pointer
	mahasiswa.UpdateIPK(3.85)
	fmt.Println("Info lengkap baru: ", mahasiswa.GetFullInfo())

	// & --> ampersand -> Ngambil alamat memmory
	x := 7
	y := &x
	println("X: ", x) // X:  7
	println("Y: ", y) // Y:  0xc0000a2e18

	num := 25
	ubahNilaiX(&num)
	fmt.Println(num)

	// fmt.Println("Cek lulus: ", mahasiswa.IsLulus())

	minIPK, status := mahasiswa.IsLulus()
	fmt.Printf("Status Lulus: %v dengan minimal IPK %.2f\n", status, minIPK)
	println()

	// Constructor
	mahasiswa2 := NewMahasiswa("Budi", "002", "Sistem Informasi", 3.20)
	mahasiswa2.PrintInfoMahasiswa()
}

func ubahNilaiX(x *int) {
	*x = 100
}

// cek lulus
//
//	func (m Mahasiswa) IsLulus() bool {
//		return m.IPK >= 2.5
//	}
func (m Mahasiswa) IsLulus() (float64, string) {
	if m.IPK >= 2.5 {
		return 2.5, "Lulus"
	} else {
		return 2.5, "Tidak Lulus"
	}
}

// constructor
func NewMahasiswa(nama, nim, jurusan string, ipk float64) *Mahasiswa {
	return &Mahasiswa{
		Nama:    nama,
		NIM:     nim,
		Jurusan: jurusan,
		IPK:     ipk,
	}
}
