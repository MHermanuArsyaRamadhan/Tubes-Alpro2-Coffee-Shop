package main

// Fungsi ini mengurutkan harga menu dari termurah ke termahal menggunakan Insertion Sort.
func insertionSortHarga(daftar *[MAX_MENU]Menu, jumlah int) {
	var i int
	var j int
	var key Menu

	for i = 1; i < jumlah; i++ {
		key = daftar[i]
		j = i - 1

		for j >= 0 && daftar[j].Harga > key.Harga {
			daftar[j+1] = daftar[j]
			j--
		}

		daftar[j+1] = key
	}
}
