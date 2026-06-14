package main


// Fungsi ini mengurutkan harga menu dari termurah ke termahal menggunakan Selection Sort.
func selectionSortHarga(daftar *[MAX_MENU]Menu, jumlah int) {
	var i int
	var j int
	var indexMin int
	var temp Menu

	for i = 0; i < jumlah-1; i++ {
		indexMin = i

		for j = i + 1; j < jumlah; j++ {
			if daftar[j].Harga < daftar[indexMin].Harga {
				indexMin = j
			}
		}

		temp = daftar[i]
		daftar[i] = daftar[indexMin]
		daftar[indexMin] = temp
	}
}

