package main

import "fmt"


// Fungsi ini menampilkan data menu yang sudah diurutkan berdasarkan harga menggunakan Insertion Sort.
func tampilkanInsertionSortHarga() {
	var hasil [MAX_MENU]Menu
	var jumlahData int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	hasil, jumlahData = salinData()
	insertionSortHarga(&hasil, jumlahData)

	fmt.Println("\n=== Hasil Insertion Sort Harga ===")
	fmt.Println("Data diurutkan dari harga termurah ke harga termahal.")
	tampilkanDaftar(hasil, jumlahData)
}