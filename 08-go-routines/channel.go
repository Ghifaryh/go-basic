package main

import "fmt"

func kirimData(ch chan string, data string) {
	ch <- data
}

func terimaData(ch chan string) {
	data := <-ch
	fmt.Println("Menerima data dari channel:", data)
}

func main() {
	fmt.Println("=== CHANNELS IN GO ===")

	// Membuat channel
	ch := make(chan string, 5)

	go kirimData(ch, "Hellow Worlds")
	go terimaData(ch)

	var input string
	fmt.Scanln(&input)
}
