package main

import "fmt"

// Fungsi ini mencari menu berdasarkan kategori menggunakan Binary Search.
func binarySearchKategori() {
	var kategoriDicari string
	var hasil [MAX_MENU]Menu
	var jumlahHasil int
	var dataTerurut [MAX_MENU]Menu
	var jumlahData int
	var kiri int
	var kanan int
	var tengah int
	var posisi int
	var i int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	fmt.Println("\n=== Binary Search Kategori ===")
	fmt.Println("Contoh kategori: non-coffee")
	fmt.Println("Kategori yang tersedia misalnya: coffee, non-coffee, food, snack")
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&kategoriDicari)

	dataTerurut, jumlahData = salinData()
	insertionSortKategori(&dataTerurut, jumlahData)

	kiri = 0
	kanan = jumlahData - 1
	posisi = -1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if dataTerurut[tengah].Kategori == kategoriDicari {
			posisi = tengah
			break
		} else if kategoriDicari < dataTerurut[tengah].Kategori {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if posisi == -1 {
		fmt.Println("Menu dengan kategori tersebut tidak ditemukan.")
		return
	}

	i = posisi

	for i >= 0 && dataTerurut[i].Kategori == kategoriDicari {
		i--
	}

	i++

	for i < jumlahData && dataTerurut[i].Kategori == kategoriDicari {
		hasil[jumlahHasil] = dataTerurut[i]
		jumlahHasil++
		i++
	}

	fmt.Println("\n=== Hasil Binary Search ===")
	tampilkanDaftar(hasil, jumlahHasil)
}
