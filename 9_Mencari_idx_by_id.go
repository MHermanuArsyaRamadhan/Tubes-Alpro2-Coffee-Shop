package main

// Fungsi ini mencari posisi index menu berdasarkan ID.
func cariIndexByID(idCari int) int {
	var i int

	for i = 0; i < jumlahMenu; i++ {
		if dataMenu[i].ID == idCari {
			return i
		}
	}

	return -1
}