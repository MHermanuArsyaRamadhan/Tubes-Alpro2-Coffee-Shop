package main

// Fungsi ini mengecek apakah sebuah kategori sudah ada di dalam daftar kategori unik.
func kategoriSudahAda(daftar [MAX_MENU]string, jumlah int, kat string) bool {
	var i int

	for i = 0; i < jumlah; i++ {
		if daftar[i] == kat {
			return true
		}
	}

	return false
}