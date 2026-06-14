package main


// Fungsi ini mengubah nilai stok menjadi teks status agar mudah dibaca.
func statusTeks(stokItem int) string {
	if stokItem <= 0 {
		return "Habis"
	}

	return "Tersedia"
}