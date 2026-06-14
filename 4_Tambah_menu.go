package main

import "fmt"

// Fungsi ini digunakan untuk menambahkan menu baru berdasarkan input dari user.
func tambahMenu() {
	var nama string
	var kategori string
	var komposisi string
	var harga int
	var stok int

	if jumlahMenu >= MAX_MENU {
		fmt.Println("Data menu sudah penuh.")
		return
	}

	fmt.Println("\n=== Tambah Menu ===")
	fmt.Println("Catatan: jika input lebih dari satu kata, gunakan underscore.")
	fmt.Println("Contoh benar: Cafe_Latte")
	fmt.Println("Contoh salah : Cafe Latte")

	fmt.Println("\nContoh nama menu: Vanilla_Latte")
	fmt.Print("Nama menu: ")
	fmt.Scan(&nama)

	fmt.Println("\nContoh kategori: coffee")
	fmt.Println("Pilihan kategori yang disarankan: coffee, non-coffee, food, snack")
	fmt.Print("Kategori: ")
	fmt.Scan(&kategori)

	fmt.Println("\nContoh harga: 28000")
	fmt.Print("Harga: ")
	fmt.Scan(&harga)

	fmt.Println("\nContoh komposisi: espresso_susu_vanilla")
	fmt.Print("Komposisi: ")
	fmt.Scan(&komposisi)

	fmt.Println("\nContoh stok: 10")
	fmt.Println("Keterangan: 0 = habis, angka positif = jumlah stok tersedia")
	fmt.Print("Stok: ")
	fmt.Scan(&stok)

	dataMenu[jumlahMenu] = Menu{
		ID:        nextID,
		Nama:      nama,
		Kategori:  kategori,
		Harga:     harga,
		Komposisi: komposisi,
		Stok:      stok,
	}

	jumlahMenu++
	nextID++

	fmt.Println("Menu berhasil ditambahkan.")
}