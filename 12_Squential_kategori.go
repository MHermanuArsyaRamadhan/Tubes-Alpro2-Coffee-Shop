package main

import "fmt"

// Fungsi ini mencari menu berdasarkan kategori menggunakan Sequential Search.
func sequentialSearchKategori() {
	var kategoriDicari string
	var hasil [MAX_MENU]Menu
	var jumlahHasil int
	var i int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	fmt.Println("\n=== Sequential Search Kategori ===")
	fmt.Println("Contoh kategori: coffee")
	fmt.Println("Kategori yang tersedia misalnya: coffee, non-coffee, food, snack")
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&kategoriDicari)

	for i = 0; i < jumlahMenu; i++ {
		if dataMenu[i].Kategori == kategoriDicari {
			hasil[jumlahHasil] = dataMenu[i]
			jumlahHasil++
		}
	}

	if jumlahHasil == 0 {
		fmt.Println("Menu dengan kategori tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\n=== Hasil Sequential Search ===")
	tampilkanDaftar(hasil, jumlahHasil)
}
