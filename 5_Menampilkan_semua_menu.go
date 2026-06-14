package main

import "fmt"

// Fungsi ini menampilkan semua data menu yang tersimpan.
func tampilkanSemuaMenu() {
	fmt.Println("\n=== Daftar Semua Menu ===")
	tampilkanDaftar(dataMenu, jumlahMenu)
}