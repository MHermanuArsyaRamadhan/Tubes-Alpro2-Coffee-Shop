package main

import "fmt"

const MAX_MENU = 100

type Menu struct {
	ID        int
	Nama      string
	Kategori  string
	Harga     int
	Komposisi string
	Stok      int
}

type ItemOrder struct {
	NamaMenu    string
	Jumlah      int
	HargaSatuan int
	Subtotal    int
}

// Hanya 3 ini yang benar-benar global karena diakses dan dimodifikasi oleh banyak fungsi
var dataMenu [MAX_MENU]Menu
var jumlahMenu int
var nextID int = 1

// Fungsi utama program untuk menampilkan menu utama dan menjalankan fitur sesuai pilihan user.
func main() {
	isiDataAwal()

	var pilihan int

	for {
		fmt.Println("\n========== CAFE-MENU ==========")
		fmt.Println("1. Tambah menu")
		fmt.Println("2. Tampilkan semua menu")
		fmt.Println("3. Ubah menu")
		fmt.Println("4. Hapus menu")
		fmt.Println("5. Cari kategori dengan Sequential Search")
		fmt.Println("6. Cari kategori dengan Binary Search")
		fmt.Println("7. Urutkan harga dengan Selection Sort")
		fmt.Println("8. Urutkan harga dengan Insertion Sort")
		fmt.Println("9. Statistik menu")
		fmt.Println("10. Order menu")
		fmt.Println("11. Ubah stok menu")
		fmt.Println("0. Keluar")
		fmt.Println("Contoh input pilihan: 1")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahMenu()
		case 2:
			tampilkanSemuaMenu()
		case 3:
			ubahMenu()
		case 4:
			hapusMenu()
		case 5:
			sequentialSearchKategori()
		case 6:
			binarySearchKategori()
		case 7:
			tampilkanSelectionSortHarga()
		case 8:
			tampilkanInsertionSortHarga()
		case 9:
			tampilkanStatistik()
		case 10:
			orderMenu()
		case 11:
			ubahStokMenu()
		case 0:
			fmt.Println("Terima kasih sudah menggunakan Cafe-Menu.")
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
