package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struct untuk merepresentasikan mahasiswa
type Mahasiswa struct {
	NIM int
}

// Fungsi untuk menggenerate array mahasiswa acak
func generateRandomMahasiswa(size int) []Mahasiswa {
	rand.Seed(time.Now().UnixNano())
	mahasiswas := make([]Mahasiswa, size)
	for i := 0; i < size; i++ {
		// Generate NIM acak
		mahasiswas[i] = Mahasiswa{NIM: rand.Intn(1000000)}
	}
	return mahasiswas
}

// Fungsi untuk mengurutkan array mahasiswa dengan algoritma bubble sort
func bubbleSort(mahasiswas []Mahasiswa) {
	n := len(mahasiswas)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Jika NIM lebih besar, maka tukar
			if mahasiswas[j].NIM > mahasiswas[j+1].NIM {
				mahasiswas[j], mahasiswas[j+1] = mahasiswas[j+1], mahasiswas[j]
			}
		}
	}
}

// Fungsi untuk mengurutkan array mahasiswa dengan algoritma insertion sort
func insertionSort(mahasiswas []Mahasiswa) {
	for i := 1; i < len(mahasiswas); i++ {
		// Simpan nilai yang akan diurutkan
		key := mahasiswas[i]
		j := i - 1
		// Jika nilai sebelumnya lebih besar, maka geser
		for j >= 0 && mahasiswas[j].NIM > key.NIM {
			mahasiswas[j+1] = mahasiswas[j]
			j--
		}
		// Jika nilai sebelumnya lebih kecil, maka simpan nilai yang diurutkan
		mahasiswas[j+1] = key
	}
}

// Fungsi untuk mengurutkan array mahasiswa dengan algoritma quick sort
func quickSort(mahasiswas []Mahasiswa, low, high int) {
	if low < high {
		// Tentukan pivot
		pi := partition(mahasiswas, low, high)
		// Rekursifkan untuk bagian kiri dan kanan
		quickSort(mahasiswas, low, pi-1)
		quickSort(mahasiswas, pi+1, high)
	}
}

// Fungsi untuk mengurutkan array mahasiswa dengan algoritma partition
func partition(mahasiswas []Mahasiswa, low, high int) int {
	// Tentukan pivot
	pivot := mahasiswas[high].NIM
	i := low - 1
	for j := low; j < high; j++ {
		// Jika nilai sebelumnya lebih kecil, maka geser
		if mahasiswas[j].NIM <= pivot {
			i++
			mahasiswas[i], mahasiswas[j] = mahasiswas[j], mahasiswas[i]
		}
	}
	// Simpan pivot di posisi yang tepat
	mahasiswas[i+1], mahasiswas[high] = mahasiswas[high], mahasiswas[i+1]
	return i + 1
}

// Fungsi untuk mengurutkan array mahasiswa dengan algoritma merge sort
func mergeSort(mahasiswas []Mahasiswa) []Mahasiswa {
	if len(mahasiswas) < 2 {
		return mahasiswas
	}
	mid := len(mahasiswas) / 2
	left := mergeSort(mahasiswas[:mid])
	right := mergeSort(mahasiswas[mid:])
	// Merge kiri dan kanan
	return merge(left, right)
}

// Fungsi untuk menggabungkan dua array yang sudah diurutkan
func merge(left, right []Mahasiswa) []Mahasiswa {
	result := make([]Mahasiswa, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		// Jika nilai kiri lebih kecil, maka append
		if left[i].NIM <= right[j].NIM {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// Jika masih ada nilai di kiri atau kanan, maka append
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// Fungsi untuk menambahkan data ke dalam map
func appendToMap(m map[int]time.Duration, size int, duration time.Duration) map[int]time.Duration {
	if m == nil {
		m = make(map[int]time.Duration)
	}
	m[size] = duration
	return m
}

// Fungsi untuk mencetak hasil
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
