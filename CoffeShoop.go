package main

import "fmt"

const MAX_MENU = 100

type Menu struct {
	ID        int
	Nama      string
	Kategori  string
	Harga     int
	Komposisi string
	Tersedia  bool
}

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
		case 0:
			fmt.Println("Terima kasih sudah menggunakan Cafe-Menu.")
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}

// Fungsi ini mengisi data awal agar program sudah memiliki contoh menu saat pertama dijalankan.
func isiDataAwal() {
	tambahDataAwal("Americano", "coffee", 18000, "espresso_air_panas", true)
	tambahDataAwal("Cappuccino", "coffee", 25000, "espresso_susu_foam", true)
	tambahDataAwal("Cafe_Latte", "coffee", 24000, "espresso_susu", true)
	tambahDataAwal("Matcha_Latte", "non-coffee", 26000, "matcha_susu_gula", true)
	tambahDataAwal("Chocolate_Ice", "non-coffee", 22000, "cokelat_susu_es", true)
	tambahDataAwal("French_Fries", "snack", 20000, "kentang_garam_saus", true)
	tambahDataAwal("Chicken_Sandwich", "food", 32000, "roti_ayam_selada_saus", false)
}

// Fungsi ini menambahkan data awal ke array dataMenu tanpa menggunakan append.
func tambahDataAwal(nama string, kategori string, harga int, komposisi string, tersedia bool) {
	if jumlahMenu >= MAX_MENU {
		return
	}

	dataMenu[jumlahMenu] = Menu{
		ID:        nextID,
		Nama:      nama,
		Kategori:  kategori,
		Harga:     harga,
		Komposisi: komposisi,
		Tersedia:  tersedia,
	}

	jumlahMenu++
	nextID++
}

// Fungsi ini digunakan untuk menambahkan menu baru berdasarkan input dari user.
func tambahMenu() {
	var nama, kategori, komposisi string
	var harga, status int

	if jumlahMenu >= MAX_MENU {
		fmt.Println("Data menu sudah penuh.")
		return
	}

	fmt.Println("\n=== Tambah Menu ===")
	fmt.Println("Catatan: jika input lebih dari satu kata, gunakan underscore.")
	fmt.Println("Contoh benar: Cafe_Latte")
	fmt.Println("Contoh salah : Cafe Latte")

	fmt.Println("\nContoh nama menu: Vanilla_Latte")
	fmt.Print("Nama menu: ")
	fmt.Scan(&nama)

	fmt.Println("\nContoh kategori: coffee")
	fmt.Println("Pilihan kategori yang disarankan: coffee, non-coffee, food, snack")
	fmt.Print("Kategori: ")
	fmt.Scan(&kategori)

	fmt.Println("\nContoh harga: 28000")
	fmt.Print("Harga: ")
	fmt.Scan(&harga)

	fmt.Println("\nContoh komposisi: espresso_susu_vanilla")
	fmt.Print("Komposisi: ")
	fmt.Scan(&komposisi)

	fmt.Println("\nContoh status: 1")
	fmt.Println("Keterangan: 1 = tersedia, 0 = habis")
	fmt.Print("Tersedia? ")
	fmt.Scan(&status)

	dataMenu[jumlahMenu] = Menu{
		ID:        nextID,
		Nama:      nama,
		Kategori:  kategori,
		Harga:     harga,
		Komposisi: komposisi,
		Tersedia:  status == 1,
	}

	jumlahMenu++
	nextID++

	fmt.Println("Menu berhasil ditambahkan.")
}

// Fungsi ini menampilkan semua data menu yang tersimpan.
func tampilkanSemuaMenu() {
	fmt.Println("\n=== Daftar Semua Menu ===")
	tampilkanDaftar(dataMenu, jumlahMenu)
}

// Fungsi ini menampilkan data menu dalam bentuk tabel yang lebih rapi.
func tampilkanDaftar(daftar [MAX_MENU]Menu, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	fmt.Println("+----+--------------------+--------------+----------+-----------+------------------------------+")
	fmt.Println("| ID | Nama Menu          | Kategori     | Harga    | Status    | Komposisi                    |")
	fmt.Println("+----+--------------------+--------------+----------+-----------+------------------------------+")

	for i := 0; i < jumlah; i++ {
		fmt.Printf("| %-2d | %-18s | %-12s | Rp%-6d | %-9s | %-28s |\n",
			daftar[i].ID,
			potongTeks(daftar[i].Nama, 18),
			potongTeks(daftar[i].Kategori, 12),
			daftar[i].Harga,
			statusTeks(daftar[i].Tersedia),
			potongTeks(daftar[i].Komposisi, 28),
		)
	}

	fmt.Println("+----+--------------------+--------------+----------+-----------+------------------------------+")
}

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

// Fungsi ini mengubah nilai boolean menjadi teks agar status menu mudah dibaca.
func statusTeks(tersedia bool) string {
	if tersedia {
		return "Tersedia"
	}

	return "Habis"
}

// Fungsi ini mencari posisi index menu berdasarkan ID.
func cariIndexByID(id int) int {
	for i := 0; i < jumlahMenu; i++ {
		if dataMenu[i].ID == id {
			return i
		}
	}

	return -1
}

// Fungsi ini digunakan untuk mengubah data menu berdasarkan ID yang dipilih user.
func ubahMenu() {
	var nama, kategori, komposisi string
	var id, harga, status int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	tampilkanSemuaMenu()

	fmt.Println("\nContoh ID yang ingin diubah: 2")
	fmt.Println("Lihat ID pada tabel di atas.")
	fmt.Print("Masukkan ID menu yang ingin diubah: ")
	fmt.Scan(&id)

	index := cariIndexByID(id)

	if index == -1 {
		fmt.Println("Menu tidak ditemukan.")
		return
	}

	fmt.Println("\nMasukkan data baru.")
	fmt.Println("Catatan: isi semua data baru. Gunakan underscore jika lebih dari satu kata.")

	fmt.Println("\nContoh nama baru: Cappuccino_Ice")
	fmt.Print("Nama menu: ")
	fmt.Scan(&nama)

	fmt.Println("\nContoh kategori baru: coffee")
	fmt.Print("Kategori: ")
	fmt.Scan(&kategori)

	fmt.Println("\nContoh harga baru: 27000")
	fmt.Print("Harga: ")
	fmt.Scan(&harga)

	fmt.Println("\nContoh komposisi baru: espresso_susu_es_foam")
	fmt.Print("Komposisi: ")
	fmt.Scan(&komposisi)

	fmt.Println("\nContoh status baru: 1")
	fmt.Println("Keterangan: 1 = tersedia, 0 = habis")
	fmt.Print("Tersedia? ")
	fmt.Scan(&status)

	dataMenu[index].Nama = nama
	dataMenu[index].Kategori = kategori
	dataMenu[index].Harga = harga
	dataMenu[index].Komposisi = komposisi
	dataMenu[index].Tersedia = status == 1

	fmt.Println("Data menu berhasil diubah.")
}

// Fungsi ini digunakan untuk menghapus menu berdasarkan ID tanpa menggunakan append.
func hapusMenu() {
	var id, konfirmasi int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	tampilkanSemuaMenu()

	fmt.Println("\nContoh ID yang ingin dihapus: 6")
	fmt.Println("Lihat ID pada tabel di atas.")
	fmt.Print("Masukkan ID menu yang ingin dihapus: ")
	fmt.Scan(&id)

	index := cariIndexByID(id)

	if index == -1 {
		fmt.Println("Menu tidak ditemukan.")
		return
	}

	fmt.Println("\nContoh konfirmasi: 1")
	fmt.Println("Keterangan: 1 = ya hapus, 0 = batal")
	fmt.Print("Yakin hapus? ")
	fmt.Scan(&konfirmasi)

	if konfirmasi != 1 {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	for i := index; i < jumlahMenu-1; i++ {
		dataMenu[i] = dataMenu[i+1]
	}

	dataMenu[jumlahMenu-1] = Menu{}
	jumlahMenu--

	fmt.Println("Menu berhasil dihapus.")
}

// Fungsi ini mencari menu berdasarkan kategori menggunakan Sequential Search.
func sequentialSearchKategori() {
	var kategoriDicari string
	var hasil [MAX_MENU]Menu
	var jumlahHasil int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	fmt.Println("\n=== Sequential Search Kategori ===")
	fmt.Println("Contoh kategori: coffee")
	fmt.Println("Kategori yang tersedia misalnya: coffee, non-coffee, food, snack")
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&kategoriDicari)

	for i := 0; i < jumlahMenu; i++ {
		if dataMenu[i].Kategori == kategoriDicari {
			hasil[jumlahHasil] = dataMenu[i]
			jumlahHasil++
		}
	}

	if jumlahHasil == 0 {
		fmt.Println("Menu dengan kategori tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\n=== Hasil Sequential Search ===")
	tampilkanDaftar(hasil, jumlahHasil)
}

// Fungsi ini mencari menu berdasarkan kategori menggunakan Binary Search.
func binarySearchKategori() {
	var kategoriDicari string
	var hasil [MAX_MENU]Menu
	var jumlahHasil int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	fmt.Println("\n=== Binary Search Kategori ===")
	fmt.Println("Contoh kategori: non-coffee")
	fmt.Println("Kategori yang tersedia misalnya: coffee, non-coffee, food, snack")
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&kategoriDicari)

	dataTerurut, jumlahData := salinData()
	insertionSortKategori(&dataTerurut, jumlahData)

	kiri := 0
	kanan := jumlahData - 1
	posisi := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if dataTerurut[tengah].Kategori == kategoriDicari {
			posisi = tengah
			break
		} else if kategoriDicari < dataTerurut[tengah].Kategori {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if posisi == -1 {
		fmt.Println("Menu dengan kategori tersebut tidak ditemukan.")
		return
	}

	i := posisi

	for i >= 0 && dataTerurut[i].Kategori == kategoriDicari {
		i--
	}

	i++

	for i < jumlahData && dataTerurut[i].Kategori == kategoriDicari {
		hasil[jumlahHasil] = dataTerurut[i]
		jumlahHasil++
		i++
	}

	fmt.Println("\n=== Hasil Binary Search ===")
	tampilkanDaftar(hasil, jumlahHasil)
}

// Fungsi ini menyalin data menu ke array baru agar data asli tidak berubah saat proses sorting.
func salinData() ([MAX_MENU]Menu, int) {
	var hasil [MAX_MENU]Menu

	for i := 0; i < jumlahMenu; i++ {
		hasil[i] = dataMenu[i]
	}

	return hasil, jumlahMenu
}

// Fungsi ini mengurutkan data berdasarkan kategori menggunakan Insertion Sort untuk kebutuhan Binary Search.
func insertionSortKategori(daftar *[MAX_MENU]Menu, jumlah int) {
	for i := 1; i < jumlah; i++ {
		key := daftar[i]
		j := i - 1

		for j >= 0 && daftar[j].Kategori > key.Kategori {
			daftar[j+1] = daftar[j]
			j--
		}

		daftar[j+1] = key
	}
}

// Fungsi ini menampilkan data menu yang sudah diurutkan berdasarkan harga menggunakan Selection Sort.
func tampilkanSelectionSortHarga() {
	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	hasil, jumlahData := salinData()
	selectionSortHarga(&hasil, jumlahData)

	fmt.Println("\n=== Hasil Selection Sort Harga ===")
	fmt.Println("Data diurutkan dari harga termurah ke harga termahal.")
	tampilkanDaftar(hasil, jumlahData)
}

// Fungsi ini mengurutkan harga menu dari termurah ke termahal menggunakan Selection Sort.
func selectionSortHarga(daftar *[MAX_MENU]Menu, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		indexMin := i

		for j := i + 1; j < jumlah; j++ {
			if daftar[j].Harga < daftar[indexMin].Harga {
				indexMin = j
			}
		}

		temp := daftar[i]
		daftar[i] = daftar[indexMin]
		daftar[indexMin] = temp
	}
}

// Fungsi ini menampilkan data menu yang sudah diurutkan berdasarkan harga menggunakan Insertion Sort.
func tampilkanInsertionSortHarga() {
	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	hasil, jumlahData := salinData()
	insertionSortHarga(&hasil, jumlahData)

	fmt.Println("\n=== Hasil Insertion Sort Harga ===")
	fmt.Println("Data diurutkan dari harga termurah ke harga termahal.")
	tampilkanDaftar(hasil, jumlahData)
}

// Fungsi ini mengurutkan harga menu dari termurah ke termahal menggunakan Insertion Sort.
func insertionSortHarga(daftar *[MAX_MENU]Menu, jumlah int) {
	for i := 1; i < jumlah; i++ {
		key := daftar[i]
		j := i - 1

		for j >= 0 && daftar[j].Harga > key.Harga {
			daftar[j+1] = daftar[j]
			j--
		}

		daftar[j+1] = key
	}
}

// Fungsi ini menampilkan statistik jumlah menu, jumlah menu per kategori, dan rata-rata harga.
func tampilkanStatistik() {
	var kategoriUnik [MAX_MENU]string
	var jumlahKategori int
	var totalHarga int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	for i := 0; i < jumlahMenu; i++ {
		totalHarga += dataMenu[i].Harga

		if !kategoriSudahAda(kategoriUnik, jumlahKategori, dataMenu[i].Kategori) {
			kategoriUnik[jumlahKategori] = dataMenu[i].Kategori
			jumlahKategori++
		}
	}

	fmt.Println("\n=== Statistik Menu Cafe ===")
	fmt.Println("Total menu:", jumlahMenu)
	fmt.Println("Rata-rata harga semua menu: Rp", totalHarga/jumlahMenu)

	fmt.Println("\nJumlah menu dan rata-rata harga per kategori:")

	for i := 0; i < jumlahKategori; i++ {
		var jumlah, totalKategori int

		for j := 0; j < jumlahMenu; j++ {
			if dataMenu[j].Kategori == kategoriUnik[i] {
				jumlah++
				totalKategori += dataMenu[j].Harga
			}
		}

		fmt.Println("Kategori:", kategoriUnik[i])
		fmt.Println("Jumlah menu:", jumlah)
		fmt.Println("Rata-rata harga: Rp", totalKategori/jumlah)
		fmt.Println("--------------------------")
	}
}

// Fungsi ini mengecek apakah sebuah kategori sudah ada di dalam daftar kategori unik.
func kategoriSudahAda(daftar [MAX_MENU]string, jumlah int, kategori string) bool {
	for i := 0; i < jumlah; i++ {
		if daftar[i] == kategori {
			return true
		}
	}

	return false
}
