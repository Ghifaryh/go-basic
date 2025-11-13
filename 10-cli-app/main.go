package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mahasiswa struct {
	NIM     string
	Nama    string
	Jurusan string
	IPK     float64
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// sentinel error
var (
	ErrNotFound     = fmt.Errorf("Mahasiswa tidak ditemukan")
	ErrAlreadyExist = fmt.Errorf("Mahasiswa sudah terdaftar")
	ErrInvalidInput = fmt.Errorf("Inputan tidak valid")
)

var daftarMahasiswa []Mahasiswa

func tampilkanMenu() {
	fmt.Println("===Menu Mahasiswa===")
	fmt.Println("1. Tambah Mahasiswa")
	fmt.Println("2. Lihat Mahasiswa")
	fmt.Println("3. Cari Mahasiswa")
	fmt.Println("4. Update IPK")
	fmt.Println("5. Hapus Mahasiswa")
	fmt.Println("6. Tampilkan Statistik")
	fmt.Println("7. Keluar")
	fmt.Print("Silahkan beropsi: ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		tampilkanMenu()
		s := scanner.Scan()
		if !s {
			fmt.Println("Terjadi kesalahan Input")
			continue
		}
		option := strings.TrimSpace(scanner.Text())
		switch option {
		case "1":
			tambahMahasiswa(scanner)
		case "2":
			lihatSemuaMahasiswa()
		case "3":
			cariByNIM(scanner.Text())
		case "4":
			// UpdateIPK(scanner)
			fmt.Println("4. Update IPK")
		case "5":
			// HapusMahasiswa(scanner)
			fmt.Println("5. Hapus Mahasiswa")
		case "6":
			// Statistik(scanner)
			fmt.Println("6. Tampilkan Statistik")
		case "7":
			// Keluar(scanner)
			fmt.Println("7. Keluar")
			fmt.Println("Silahkan beropsi")
		default:
			fmt.Println("Pilihan tidak valdi. Silahkan coba lagi...")
		}

	}
}

func tambahMahasiswa(scanner *bufio.Scanner) {
	var m Mahasiswa
	fmt.Print("Masukan NIM: ")
	scanner.Scan()
	m.NIM = strings.TrimSpace(scanner.Text())

	fmt.Print("Masukkan Nama: ")
	scanner.Scan()
	m.Nama = strings.TrimSpace(scanner.Text())

	fmt.Print("Masukkan Jurusan: ")
	scanner.Scan()
	m.Jurusan = strings.TrimSpace(scanner.Text())

	fmt.Print("Masukkan IPK: ")
	scanner.Scan()
	ipkString := strings.TrimSpace(scanner.Text())

	var err error
	if m.IPK, err = validasiIPK(ipkString); err != nil {
		fmt.Println(err)
		return
	}

	if err := validasiMahasiswa(m.NIM, m.Nama, m.Jurusan); err != nil {
		fmt.Println(err)
		return
	}

	daftarMahasiswa = append(daftarMahasiswa, m)
	fmt.Println("[v] Mahasiswa Berhasil ditambahkan.")

}

func validasiIPK(ipkString string) (float64, error) {
	// convert string to float
	ipk, err := strconv.ParseFloat(ipkString, 64)
	if err != nil || ipk < 0 || ipk > 4 {
		return 0, ErrInvalidInput
	}
	return ipk, nil
}

func validasiMahasiswa(nim, nama, jurusan string) error {
	if nim == "" || nama == "" || jurusan == "" {
		return &ValidationError{Field: "Input", Message: "Semua Field Harus Diisi!"}
	}
	return nil
}

func lihatSemuaMahasiswa() {
	fmt.Printf("%-10s %-20s %-20s %-10s\n", "NIM", "NAMA", "JURUSAN", "IPK") // jarak spasi
	fmt.Println(strings.Repeat("-", 60))
	for _, m := range daftarMahasiswa {
		fmt.Printf("%-10s %-20s %-20s %-10f\n", m.NIM, m.Nama, m.Jurusan, m.IPK)
	}
}

func cariByNIM(nim string) (*Mahasiswa, error) { // still errors
	for _, m := range daftarMahasiswa {
		if m.NIM == nim {
			return &m, nil
		}
	}
	return nil, ErrNotFound
}
