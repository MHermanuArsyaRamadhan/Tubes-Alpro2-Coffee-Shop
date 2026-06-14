package main

// Fungsi ini mengurutkan data berdasarkan kategori menggunakan Insertion Sort untuk kebutuhan Binary Search.
func insertionSortKategori(daftar *[MAX_MENU]Menu, jumlah int) {
	var i int
	var j int
	var key Menu

	for i = 1; i < jumlah; i++ {
		key = daftar[i]
		j = i - 1

		for j >= 0 && daftar[j].Kategori > key.Kategori {
			daftar[j+1] = daftar[j]
			j--
		}

		daftar[j+1] = key
	}
}