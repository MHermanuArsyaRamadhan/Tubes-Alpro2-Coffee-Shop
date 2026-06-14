package main

import "fmt"


// Fungsi ini digunakan untuk menghapus menu berdasarkan ID tanpa menggunakan append.
func hapusMenu() {
	var id int
	var index int
	var konfirmasi int
	var i int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	tampilkanSemuaMenu()

	fmt.Println("\nContoh ID yang ingin dihapus: 6")
	fmt.Println("Lihat ID pada tabel di atas.")
	fmt.Print("Masukkan ID menu yang ingin dihapus: ")
	fmt.Scan(&id)

	index = cariIndexByID(id)

	if index == -1 {
		fmt.Println("Menu tidak ditemukan.")
		return
	}

	fmt.Println("\nContoh konfirmasi: 1")
	fmt.Println("Keterangan: 1 = ya hapus, 0 = batal")
	fmt.Print("Yakin hapus? ")
	fmt.Scan(&konfirmasi)

	if konfirmasi != 1 {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	for i = index; i < jumlahMenu-1; i++ {
		dataMenu[i] = dataMenu[i+1]
	}

	dataMenu[jumlahMenu-1] = Menu{}
	jumlahMenu--

	fmt.Println("Menu berhasil dihapus.")
}