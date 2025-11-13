package main

import "fmt"

func main() {
	fmt.Println("=============================")
	fmt.Println("    BELAJAR STRUKTUR DATA    ")
	fmt.Println("=============================")
	fmt.Println()

	//array
	demonstrasiArray()
	demonstrasiSlice()
	demonstrasiMap()
}

func demonstrasiArray() {
	fmt.Println("1. Array")
	fmt.Println("--------")

	// array dengan ukuran tetap
	var angka [5]int = [5]int{1, 2, 3, 4, 5}
	buah := [3]string{"apel", "jeruk", "mangga"}

	// ukuran otomatis
	warna := [...]string{"Merah", "Kuning", "Hijau"}

	fmt.Printf("Array angka: %v\n", angka)
	fmt.Printf("Array buah: %v\n", buah)
	fmt.Printf("Array warna: %v\n", warna)
	fmt.Printf("Banyak array warna: %d\n", len(warna))
	fmt.Printf("Buah Pertama: %v\n", buah[0])
	fmt.Printf("Warna Terakhir: %v\n", warna[len(warna)-1])
	fmt.Println()
}

func demonstrasiSlice() {
	fmt.Println("2. Slice")
	fmt.Println("--------")

	var bahasa []string
	fmt.Printf("Sebelum append: %v (length: %d, capacity: %d)\n", bahasa, len(bahasa), cap(bahasa))

	bahasa = append(bahasa, "GO")
	bahasa = append(bahasa, "Javascript", "CPP")
	fmt.Printf("Setelah append: %v (length: %d, capacity: %d)\n", bahasa, len(bahasa), cap(bahasa))
	fmt.Println()
}

func demonstrasiMap() {
	fmt.Println("3. Map")
	fmt.Println("------")

	populasi := make(map[string]int)
	populasi["Jakarta"] = 1000000
	populasi["Bandung"] = 100000
	populasi["Depok"] = 500000000

	fmt.Println("Data Populasi Kota:")
	for kota, jumlah := range populasi {
		fmt.Printf("%s: %d Jiwa\n", kota, jumlah)
	}
	fmt.Println()

	mapInterface := map[string]interface{}{
		"nama":     "Gip",
		"age":      25,
		"isActive": true,
	}
	fmt.Printf("Map: interface: %v\n", mapInterface) // %v is for array

	mapIntegerString := map[string]int{
		"Aditya": 17,
		"Gip":    30,
	}
	fmt.Printf("Map: %v\n", mapIntegerString)
}
