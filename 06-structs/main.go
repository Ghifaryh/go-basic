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

func (m Mahasiswa) GetFullInfo() string {
	return fmt.Sprintf("%s (NIM: %s) - %s", m.Nama, m.NIM, m.Jurusan)
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
}
