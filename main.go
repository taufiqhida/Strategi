package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Mahasiswa struct {
	NIM int
}

func generateRandomMahasiswa(size int) []Mahasiswa {
	rand.Seed(time.Now().UnixNano())
	mahasiswas := make([]Mahasiswa, size)
	for i := 0; i < size; i++ {
		mahasiswas[i] = Mahasiswa{NIM: rand.Intn(1000000)}
	}
	return mahasiswas
}

func bubbleSort(mahasiswas []Mahasiswa) {
	n := len(mahasiswas)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if mahasiswas[j].NIM > mahasiswas[j+1].NIM {
				mahasiswas[j], mahasiswas[j+1] = mahasiswas[j+1], mahasiswas[j]
			}
		}
	}
}

func insertionSort(mahasiswas []Mahasiswa) {
	for i := 1; i < len(mahasiswas); i++ {
		key := mahasiswas[i]
		j := i - 1
		for j >= 0 && mahasiswas[j].NIM > key.NIM {
			mahasiswas[j+1] = mahasiswas[j]
			j--
		}
		mahasiswas[j+1] = key
	}
}

func quickSort(mahasiswas []Mahasiswa, low, high int) {
	if low < high {
		pi := partition(mahasiswas, low, high)
		quickSort(mahasiswas, low, pi-1)
		quickSort(mahasiswas, pi+1, high)
	}
}

func partition(mahasiswas []Mahasiswa, low, high int) int {
	pivot := mahasiswas[high].NIM
	i := low - 1
	for j := low; j < high; j++ {
		if mahasiswas[j].NIM <= pivot {
			i++
			mahasiswas[i], mahasiswas[j] = mahasiswas[j], mahasiswas[i]
		}
	}
	mahasiswas[i+1], mahasiswas[high] = mahasiswas[high], mahasiswas[i+1]
	return i + 1
}

func mergeSort(mahasiswas []Mahasiswa) []Mahasiswa {
	if len(mahasiswas) < 2 {
		return mahasiswas
	}
	mid := len(mahasiswas) / 2
	left := mergeSort(mahasiswas[:mid])
	right := mergeSort(mahasiswas[mid:])
	return merge(left, right)
}

func merge(left, right []Mahasiswa) []Mahasiswa {
	result := make([]Mahasiswa, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].NIM <= right[j].NIM {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func appendToMap(m map[int]time.Duration, size int, duration time.Duration) map[int]time.Duration {
	if m == nil {
		m = make(map[int]time.Duration)
	}
	m[size] = duration
	return m
}

func printResults(results map[string]map[int]time.Duration) {
	for algo, data := range results {
		fmt.Println(algo)
		for size, duration := range data {
			fmt.Printf("  %d items: %v ms\n", size, float64(duration)/float64(time.Millisecond))
		}
		fmt.Println("---")
	}
}

func main() {
	sizes := []int{10, 100, 1000, 10000, 100000}
	results := make(map[string]map[int]time.Duration)

	for _, size := range sizes {
		// Generate data acak
		data := generateRandomMahasiswa(size)

		// Bubble Sort
		startTime := time.Now()
		bubbleSort(append([]Mahasiswa(nil), data...))
		results["Bubble Sort"] = appendToMap(results["Bubble Sort"], size, time.Since(startTime))

		// Insertion Sort
		startTime = time.Now()
		insertionSort(append([]Mahasiswa(nil), data...))
		results["Insertion Sort"] = appendToMap(results["Insertion Sort"], size, time.Since(startTime))

		// Quick Sort
		startTime = time.Now()
		quickSort(append([]Mahasiswa(nil), data...), 0, len(data)-1)
		results["Quick Sort"] = appendToMap(results["Quick Sort"], size, time.Since(startTime))

		// Merge Sort
		startTime = time.Now()
		mergeSort(append([]Mahasiswa(nil), data...))
		results["Merge Sort"] = appendToMap(results["Merge Sort"], size, time.Since(startTime))
	}

	// Cetak hasil
	printResults(results)
}
