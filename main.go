package main

import (
	"fmt"
	"time" // Paket yang diperlukan untuk pengukuran waktu
	"math/rand"
)

const NMAX = 10000

type arr[NMAX] int 

func createRandomArray(a *arr, n int) {
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n + 1)
	}
}

func findMaxIteratif(a arr, n int)int{
	var i, max int
	max = a[0]
	for i = 0; i < n;i++{
		if max < a[i]{
			max = a[i]
		}
	}
	return max
}

func findMaxRecursive(a arr, n int) int {
	if n == 1 {
		return a[0]
	}
	var maxRest int
	maxRest = findMaxRecursive(a, n-1)
	if a[n-1] > maxRest { 
		return a[n-1]
	}
	return maxRest
}
func print(a arr, n int){
	var i int
	for i =0; i <n;i ++{
		fmt.Print(a[i], " ")
	}
}
func main(){
	var a arr
	var n, max int
	var pilih int
	var cek bool
	
	cek = true
    
    fmt.Print("Masukan nilai Array, namun <= 10000: ")
	fmt.Scan(&n)
    
    fmt.Print("Nilai akan di genarate random sebanyak: ", n)
	createRandomArray(&a, n)
    
	for cek == true{ 	
		fmt.Print("\nPilih algortima Iteratif atau Rekursif\n")
		fmt.Println("Pilih 1 untuk Iteratif, Pilih 2 untuk rekursif, Pilih 3 untuk program berhenti")
		fmt.Scan(&pilih)
        
        var start time.Time
        var duration time.Duration
        
		if pilih == 1{
			fmt.Println("Mencari nilai max dengan algortima iteratif")
            
            start = time.Now()
			max = findMaxIteratif(a, n)
            duration = time.Since(start)
            
			cek = true
		} else if pilih == 2{ 
			fmt.Println("Mencari nilai max dengan algortima rekursif")
            
            start = time.Now()
			max = findMaxRecursive(a, n)
            duration = time.Since(start)
            
			cek = true
		}else if pilih == 3{
			cek = false
		}else {
			fmt.Print("Pilihan hanya boleh 1, 2 dan 3\n")
            cek = true
		}
        
        fmt.Printf("Nilai max: %d\n", max)
        fmt.Printf("Waktu eksekusi: %d ns\n", duration.Nanoseconds())
	}
    
}