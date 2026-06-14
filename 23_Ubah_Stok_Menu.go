package main

import "fmt"


// Fungsi ini digunakan untuk mengubah stok menu secara langsung berdasarkan ID.
func ubahStokMenu() {
	var id int
	var index int
	var stokBaru int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	tampilkanSemuaMenu()

	fmt.Println("\nContoh ID yang ingin diubah stoknya: 3")
	fmt.Println("Lihat ID pada tabel di atas.")
	fmt.Print("Masukkan ID menu yang ingin diubah stoknya: ")
	fmt.Scan(&id)

	index = cariIndexByID(id)

	if index == -1 {
		fmt.Println("Menu tidak ditemukan.")
		return
	}

	fmt.Printf("\nStok saat ini untuk %s: %d\n", dataMenu[index].Nama, dataMenu[index].Stok)
	fmt.Println("Contoh stok baru: 20")
	fmt.Println("Keterangan: 0 = habis, angka positif = jumlah stok tersedia")
	fmt.Print("Masukkan stok baru: ")
	fmt.Scan(&stokBaru)

	if stokBaru < 0 {
		fmt.Println("Stok tidak boleh negatif.")
		return
	}

	dataMenu[index].Stok = stokBaru

	fmt.Printf("Stok %s berhasil diubah menjadi %d.\n", dataMenu[index].Nama, dataMenu[index].Stok)
}