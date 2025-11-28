// algo pertama
package main

import (
	"fmt"
	"time"
)

func bubbleSortIterative(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func bubbleSortRecursive(arr []int, n int) {
	if n == 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	bubbleSortRecursive(arr, n-1)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah bulan pengeluaran: ")
	fmt.Scan(&n)

	pengeluaran := make([]int, n)
	fmt.Println("\nMasukkan biaya pengeluaran per bulan:")
	for i := 0; i < n; i++ {
		fmt.Printf("Bulan %d = ", i+1)
		fmt.Scan(&pengeluaran[i])
	}

	iterArr := append([]int{}, pengeluaran...)
	recArr := append([]int{}, pengeluaran...)

	startIter := time.Now()
	bubbleSortIterative(iterArr)
	durationIter := time.Since(startIter)

	startRec := time.Now()
	bubbleSortRecursive(recArr, len(recArr))
	durationRec := time.Since(startRec)

	fmt.Println("\n==============================")
	fmt.Println("HASIL SORTING")
	fmt.Println("==============================")
	fmt.Println("Pengeluaran (Iteratif):", iterArr)
	fmt.Println("Pengeluaran (Rekursif):", recArr)

	fmt.Println("\n==============================")
	fmt.Println("PERBANDINGAN WAKTU EKSEKUSI")
	fmt.Println("==============================")
	fmt.Println("Iteratif:", durationIter)
	fmt.Println("Rekursif:", durationRec)
}
