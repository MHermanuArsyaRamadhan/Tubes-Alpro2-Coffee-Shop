package main

// Fungsi ini menyalin data menu ke array baru agar data asli tidak berubah saat proses sorting.
func salinData() ([MAX_MENU]Menu, int) {
	var salinan [MAX_MENU]Menu
	var i int

	for i = 0; i < jumlahMenu; i++ {
		salinan[i] = dataMenu[i]
	}

	return salinan, jumlahMenu
}