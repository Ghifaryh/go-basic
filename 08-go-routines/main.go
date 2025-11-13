package main

import (
	"fmt"
	"time"
)

func cetakAngka(name string, angka int) {
	for i := 1; i <= angka; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	fmt.Println("=== WITHOUT GOROUTINES ===")
	cetakAngka("First", 3)
	cetakAngka("Second", 3)

	fmt.Println("\n=== WITH GOROUTINES ===")
	go cetakAngka("Goroutine-1", 3) // Runs in background
	go cetakAngka("Goroutine-2", 3) // Runs in background
	cetakAngka("Main", 3)           // Runs in main thread

	time.Sleep(time.Second * 2) // Wait for all goroutines
}
