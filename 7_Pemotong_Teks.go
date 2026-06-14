package main


// Fungsi ini memotong teks yang terlalu panjang agar tabel tetap tertata.
func potongTeks(teks string, batas int) string {
	if len(teks) <= batas {
		return teks
	}

	if batas <= 3 {
		return teks[:batas]
	}

	return teks[:batas-3] + "..."
}
