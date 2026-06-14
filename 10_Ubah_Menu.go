package main

import "fmt"

// Fungsi ini digunakan untuk mengubah data menu berdasarkan ID yang dipilih user.
func ubahMenu() {
	var id int
	var index int
	var nama string
	var kategori string
	var komposisi string
	var harga int
	var stok int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	tampilkanSemuaMenu()

	fmt.Println("\nContoh ID yang ingin diubah: 2")
	fmt.Println("Lihat ID pada tabel di atas.")
	fmt.Print("Masukkan ID menu yang ingin diubah: ")
	fmt.Scan(&id)

	index = cariIndexByID(id)

	if index == -1 {
		fmt.Println("Menu tidak ditemukan.")
		return
	}

	fmt.Println("\nMasukkan data baru.")
	fmt.Println("Catatan: isi semua data baru. Gunakan underscore jika lebih dari satu kata.")

	fmt.Println("\nContoh nama baru: Cappuccino_Ice")
	fmt.Print("Nama menu: ")
	fmt.Scan(&nama)

	fmt.Println("\nContoh kategori baru: coffee")
	fmt.Print("Kategori: ")
	fmt.Scan(&kategori)

	fmt.Println("\nContoh harga baru: 27000")
	fmt.Print("Harga: ")
	fmt.Scan(&harga)

	fmt.Println("\nContoh komposisi baru: espresso_susu_es_foam")
	fmt.Print("Komposisi: ")
	fmt.Scan(&komposisi)

	fmt.Println("\nContoh stok baru: 10")
	fmt.Println("Keterangan: 0 = habis, angka positif = jumlah stok tersedia")
	fmt.Print("Stok: ")
	fmt.Scan(&stok)

	dataMenu[index].Nama = nama
	dataMenu[index].Kategori = kategori
	dataMenu[index].Harga = harga
	dataMenu[index].Komposisi = komposisi
	dataMenu[index].Stok = stok

	fmt.Println("Data menu berhasil diubah.")
}
