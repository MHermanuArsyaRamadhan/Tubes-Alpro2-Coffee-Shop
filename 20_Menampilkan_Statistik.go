package main

import "fmt"

// Fungsi ini menampilkan statistik jumlah menu, jumlah menu per kategori, dan rata-rata harga.
func tampilkanStatistik() {
	var kategoriUnik [MAX_MENU]string
	var jumlahKategori int
	var totalHarga int
	var i int
	var k int
	var j int
	var jumlahPerKategori int
	var totalPerKategori int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	for i = 0; i < jumlahMenu; i++ {
		totalHarga += dataMenu[i].Harga

		if !kategoriSudahAda(kategoriUnik, jumlahKategori, dataMenu[i].Kategori) {
			kategoriUnik[jumlahKategori] = dataMenu[i].Kategori
			jumlahKategori++
		}
	}

	fmt.Println("\n=== Statistik Menu Cafe ===")
	fmt.Println("Total menu:", jumlahMenu)
	fmt.Println("Rata-rata harga semua menu: Rp", totalHarga/jumlahMenu)

	fmt.Println("\nJumlah menu dan rata-rata harga per kategori:")

	for k = 0; k < jumlahKategori; k++ {
		jumlahPerKategori = 0
		totalPerKategori = 0

		for j = 0; j < jumlahMenu; j++ {
			if dataMenu[j].Kategori == kategoriUnik[k] {
				jumlahPerKategori++
				totalPerKategori += dataMenu[j].Harga
			}
		}

		fmt.Println("Kategori:", kategoriUnik[k])
		fmt.Println("Jumlah menu:", jumlahPerKategori)
		fmt.Println("Rata-rata harga: Rp", totalPerKategori/jumlahPerKategori)
		fmt.Println("--------------------------")
	}
}
