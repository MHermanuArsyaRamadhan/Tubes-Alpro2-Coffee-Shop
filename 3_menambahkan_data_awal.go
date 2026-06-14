package main

// Fungsi ini menambahkan data awal ke array dataMenu tanpa menggunakan append.
func tambahDataAwal(namaBaru string, kategoriBaru string, hargaBaru int, komposisiBaru string, stokBaru int) {
	if jumlahMenu >= MAX_MENU {
		return
	}

	dataMenu[jumlahMenu] = Menu{
		ID:        nextID,
		Nama:      namaBaru,
		Kategori:  kategoriBaru,
		Harga:     hargaBaru,
		Komposisi: komposisiBaru,
		Stok:      stokBaru,
	}

	jumlahMenu++
	nextID++
}
