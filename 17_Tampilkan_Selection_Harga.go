package main

import "fmt"

// Fungsi ini menampilkan data menu yang sudah diurutkan berdasarkan harga menggunakan Selection Sort.
func tampilkanSelectionSortHarga() {
	var hasil [MAX_MENU]Menu
	var jumlahData int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	hasil, jumlahData = salinData()
	selectionSortHarga(&hasil, jumlahData)

	fmt.Println("\n=== Hasil Selection Sort Harga ===")
	fmt.Println("Data diurutkan dari harga termurah ke harga termahal.")
	tampilkanDaftar(hasil, jumlahData)
}