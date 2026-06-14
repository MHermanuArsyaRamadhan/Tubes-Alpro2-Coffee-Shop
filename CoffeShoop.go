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

var dataMenu [MAX_MENU]Menu
var jumlahMenu int
var nextID int = 1


//BAGIAN PERTAMA 
// MAIN PROGRAM

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





// BAGIAN KE 2
// DATA AWAL

func isiDataAwal() {
	tambahDataAwal("Americano", "coffee", 18000, "espresso_air_panas", 10)
	tambahDataAwal("Cappuccino", "coffee", 25000, "espresso_susu_foam", 8)
	tambahDataAwal("Cafe_Latte", "coffee", 24000, "espresso_susu", 5)
	tambahDataAwal("Matcha_Latte", "non-coffee", 26000, "matcha_susu_gula", 7)
	tambahDataAwal("Chocolate_Ice", "non-coffee", 22000, "cokelat_susu_es", 6)
	tambahDataAwal("French_Fries", "snack", 20000, "kentang_garam_saus", 15)
	tambahDataAwal("Chicken_Sandwich", "food", 32000, "roti_ayam_selada_saus", 0)
}



// Fungsi untuk menambahkan data awal ke array dataMenu
func tambahDataAwal(namaBaru string, kategoriBaru string, hargaBaru int, komposisiBaru string, stokBaru int) {
	var nama, kategori, komposisi string
	var harga, stok int


// Untuk menyalin nilai parameter ke variable lokal
	nama = namaBaru
	kategori = kategoriBaru
	harga = hargaBaru
	komposisi = komposisiBaru
	stok = stokBaru
// Untuk MemValidasi kapasitas data menu
	if jumlahMenu >= MAX_MENU {
		return
	}
	
	
// Untuk Menyimpan data menu baru ke array global
	var m Menu
	m.ID = nextID
	m.Nama = nama
	m.Kategori = kategori
	m.Harga = harga
	m.Komposisi = komposisi
	m.Stok = stok

	dataMenu[jumlahMenu] = m

	jumlahMenu++
	nextID++
}




// BAGIAN 3
// CRUD MENU (TAMBAH, UBAH, HAPUS)


// Fungsi ini digunakan untuk menambahkan menu baru berdasarkan input dari user.
func tambahMenu() {
	var nama string
	var kategori string
	var komposisi string
	var harga int
	var stok int

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

	fmt.Println("\nContoh stok: 10")
	fmt.Println("Keterangan: 0 = habis, angka positif = jumlah stok tersedia")
	fmt.Print("Stok: ")
	fmt.Scan(&stok)

	dataMenu[jumlahMenu] = Menu{
		ID:        nextID,
		Nama:      nama,
		Kategori:  kategori,
		Harga:     harga,
		Komposisi: komposisi,
		Stok:      stok,
	}

	jumlahMenu++
	nextID++

	fmt.Println("Menu berhasil ditambahkan.")
}


// BAGIAN 4
// FUNGSI TAMPILAN (DISPLAY HELPER)


// Fungsi ini menampilkan semua data menu yang tersimpan.
func tampilkanSemuaMenu() {
	fmt.Println("\n=== Daftar Semua Menu ===")
	tampilkanDaftar(dataMenu, jumlahMenu)
}

// Fungsi ini untuk menampilkan data menu dalam bentuk tabel yang lebih rapi.
func tampilkanDaftar(daftarMenu [MAX_MENU]Menu, jumlahData int) {
	// Deklarasi variable lokal untuk menampung nilai dari parameter
	var daftar [MAX_MENU]Menu
	var jumlah int
	var i int

// Menyalin nilai parameter ke variable lokal
	daftar = daftarMenu
	jumlah = jumlahData

	if jumlah == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	fmt.Println("+----+--------------------+--------------+----------+-------+-----------+------------------------------+")
	fmt.Println("| ID | Nama Menu          | Kategori     | Harga    | Stok  | Status    | Komposisi                    |")
	fmt.Println("+----+--------------------+--------------+----------+-------+-----------+------------------------------+")

	for i = 0; i < jumlah; i++ {
		fmt.Printf("| %-2d | %-18s | %-12s | Rp%-6d | %-5d | %-9s | %-28s |\n",
			daftar[i].ID,
			potongTeks(daftar[i].Nama, 18),
			potongTeks(daftar[i].Kategori, 12),
			daftar[i].Harga,
			daftar[i].Stok,
			statusTeks(daftar[i].Stok),
			potongTeks(daftar[i].Komposisi, 28),
		)
	}

	fmt.Println("+----+--------------------+--------------+----------+-------+-----------+------------------------------+")
}

// Fungsi ini memotong teks yang terlalu panjang agar tabel tetap tertata.
func potongTeks(teksAsli string, batasAsli int) string {
// Deklarasi variable lokal untuk menampung nilai dari parameter
	var teks string
	var batas int

// Menyalin nilai parameter ke variable lokal
	teks = teksAsli
	batas = batasAsli

	if len(teks) <= batas {
		return teks
	}

	if batas <= 3 {
		return teks[:batas]
	}

	return teks[:batas-3] + "..."
}

// Fungsi ini mengubah nilai stok menjadi teks status agar mudah dibaca.
func statusTeks(stokAsli int) string {
// Deklarasi variable lokal untuk menampung nilai dari parameter
	var stok int

// Menyalin nilai parameter ke variable lokal
	stok = stokAsli

	if stok <= 0 {
		return "Habis"
	}

	return "Tersedia"
}


// BAGIAN 5
// FUNGSI BANTUAN (HELPER UMUM)


// Fungsi ini mencari posisi index menu berdasarkan ID.
func cariIndexByID(idCari int) int {
// Deklarasi variable lokal untuk menampung nilai dari parameter
	var id int
	var i int

// Menyalin nilai parameter ke variable lokal
	id = idCari

	for i = 0; i < jumlahMenu; i++ {
		if dataMenu[i].ID == id {
			return i
		}
	}

	return -1
}


// BAGIAN 6:
// CRUD MENU - LANJUTAN (UBAH & HAPUS)


// Fungsi ini digunakan untuk mengubah data menu berdasarkan ID yang dipilih user.
func ubahMenu() {
	var id int
	var index int
	var nama string
	var kategori string
	var komposisi string
	var harga int
	var stok int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	tampilkanSemuaMenu()

	fmt.Println("\nContoh ID yang ingin diubah: 2")
	fmt.Println("Lihat ID pada tabel di atas.")
	fmt.Print("Masukkan ID menu yang ingin diubah: ")
	fmt.Scan(&id)

	index = cariIndexByID(id)

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

	fmt.Println("\nContoh stok baru: 10")
	fmt.Println("Keterangan: 0 = habis, angka positif = jumlah stok tersedia")
	fmt.Print("Stok: ")
	fmt.Scan(&stok)

	dataMenu[index].Nama = nama
	dataMenu[index].Kategori = kategori
	dataMenu[index].Harga = harga
	dataMenu[index].Komposisi = komposisi
	dataMenu[index].Stok = stok

	fmt.Println("Data menu berhasil diubah.")
}

// Fungsi ini digunakan untuk menghapus menu berdasarkan ID 
func hapusMenu() {
	var id int
	var index int
	var konfirmasi int
	var i int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	tampilkanSemuaMenu()

	fmt.Println("\nContoh ID yang ingin dihapus: 6")
	fmt.Println("Lihat ID pada tabel di atas.")
	fmt.Print("Masukkan ID menu yang ingin dihapus: ")
	fmt.Scan(&id)

	index = cariIndexByID(id)

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

	for i = index; i < jumlahMenu-1; i++ {
		dataMenu[i] = dataMenu[i+1]
	}

	dataMenu[jumlahMenu-1] = Menu{}
	jumlahMenu--

	fmt.Println("Menu berhasil dihapus.")
}


// BAGIAN 7
// PENCARIAN (SEARCHING)


// Fungsi ini mencari menu berdasarkan kategori menggunakan Sequential Search.
func sequentialSearchKategori() {
	var kategoriDicari string
	var hasil [MAX_MENU]Menu
	var jumlahHasil int
	var i int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	fmt.Println("\n=== Sequential Search Kategori ===")
	fmt.Println("Contoh kategori: coffee")
	fmt.Println("Kategori yang tersedia misalnya: coffee, non-coffee, food, snack")
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&kategoriDicari)

	for i = 0; i < jumlahMenu; i++ {
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
	var dataTerurut [MAX_MENU]Menu
	var jumlahData int
	var kiri int
	var kanan int
	var tengah int
	var posisi int
	var i int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	fmt.Println("\n=== Binary Search Kategori ===")
	fmt.Println("Contoh kategori: non-coffee")
	fmt.Println("Kategori yang tersedia misalnya: coffee, non-coffee, food, snack")
	fmt.Print("Masukkan kategori yang dicari: ")
	fmt.Scan(&kategoriDicari)

	dataTerurut, jumlahData = salinData()
	insertionSortKategori(&dataTerurut, jumlahData)

	kiri = 0
	kanan = jumlahData - 1
	posisi = -1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

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

	i = posisi

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


// BAGIAN 8
// PENGURUTAN (SORTING)


// Fungsi ini menyalin data menu ke array baru agar data asli tidak berubah saat proses sorting.
func salinData() ([MAX_MENU]Menu, int) {
	var salinan [MAX_MENU]Menu
	var i int

	for i = 0; i < jumlahMenu; i++ {
		salinan[i] = dataMenu[i]
	}

	return salinan, jumlahMenu
}

// Fungsi ini mengurutkan data berdasarkan kategori menggunakan Insertion Sort untuk kebutuhan Binary Search.
func insertionSortKategori(daftarAsli *[MAX_MENU]Menu, jumlahAsli int) {
// Deklarasi variable lokal untuk menampung nilai dari parameter
	var daftar *[MAX_MENU]Menu
	var jumlah int
	var i, j int
	var key Menu

// Menyalin nilai parameter ke variable lokal
	daftar = daftarAsli
	jumlah = jumlahAsli

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

// Fungsi ini menampilkan data menu yang sudah diurutkan berdasarkan harga menggunakan Selection Sort.
func tampilkanSelectionSortHarga() {
	var hasil [MAX_MENU]Menu
	var jumlahData int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	hasil, jumlahData = salinData()
	selectionSortHarga(&hasil, jumlahData)

	fmt.Println("\n=== Hasil Selection Sort Harga ===")
	fmt.Println("Data diurutkan dari harga termurah ke harga termahal.")
	tampilkanDaftar(hasil, jumlahData)
}

// Fungsi ini mengurutkan harga menu dari termurah ke termahal menggunakan Selection Sort.
func selectionSortHarga(daftarAsli *[MAX_MENU]Menu, jumlahAsli int) {
// Deklarasi variable lokal untuk menampung nilai dari parameter
	var daftar *[MAX_MENU]Menu
	var jumlah int
	var i, j int
	var indexMin int
	var temp Menu

// Menyalin nilai parameter ke variable lokal
	daftar = daftarAsli
	jumlah = jumlahAsli

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

// Fungsi ini menampilkan data menu yang sudah diurutkan berdasarkan harga menggunakan Insertion Sort.
func tampilkanInsertionSortHarga() {
	var hasil [MAX_MENU]Menu
	var jumlahData int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	hasil, jumlahData = salinData()
	insertionSortHarga(&hasil, jumlahData)

	fmt.Println("\n=== Hasil Insertion Sort Harga ===")
	fmt.Println("Data diurutkan dari harga termurah ke harga termahal.")
	tampilkanDaftar(hasil, jumlahData)
}

// Fungsi ini mengurutkan harga menu dari termurah ke termahal menggunakan Insertion Sort.
func insertionSortHarga(daftarAsli *[MAX_MENU]Menu, jumlahAsli int) {
	// Deklarasi variable lokal untuk menampung nilai dari parameter
	var daftar *[MAX_MENU]Menu
	var jumlah int
	var i, j int
	var key Menu

	// Menyalin nilai parameter ke variable lokal
	daftar = daftarAsli
	jumlah = jumlahAsli

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


// BAGIAN 9: STATISTIK MENU

// Fungsi ini menampilkan statistik jumlah menu, jumlah menu per kategori, dan rata-rata harga.
func tampilkanStatistik() {
	var kategoriUnik [MAX_MENU]string
	var jumlahKategori int
	var totalHarga int
	var i int
	var k int
	var j int
	var jumlahPerKategori int
	var totalPerKategori int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	for i = 0; i < jumlahMenu; i++ {
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

	for k = 0; k < jumlahKategori; k++ {
		jumlahPerKategori = 0
		totalPerKategori = 0

		for j = 0; j < jumlahMenu; j++ {
			if dataMenu[j].Kategori == kategoriUnik[k] {
				jumlahPerKategori++
				totalPerKategori += dataMenu[j].Harga
			}
		}

		fmt.Println("Kategori:", kategoriUnik[k])
		fmt.Println("Jumlah menu:", jumlahPerKategori)
		fmt.Println("Rata-rata harga: Rp", totalPerKategori/jumlahPerKategori)
		fmt.Println("--------------------------")
	}
}

// Fungsi ini mengecek apakah sebuah kategori sudah ada di dalam daftar kategori unik.
func kategoriSudahAda(daftarAsli [MAX_MENU]string, jumlahAsli int, katAsli string) bool {
// Deklarasi variable lokal untuk menampung nilai dari parameter
	var daftar [MAX_MENU]string
	var jumlah int
	var kat string
	var i int

// Menyalin nilai parameter ke variable lokal
	daftar = daftarAsli
	jumlah = jumlahAsli
	kat = katAsli

	for i = 0; i < jumlah; i++ {
		if daftar[i] == kat {
			return true
		}
	}

	return false
}


// BAGIAN 10
// TRANSAKSI (ORDER & UBAH STOK)


// Fungsi ini digunakan untuk melakukan order satu atau lebih menu sekaligus dan mengurangi stok secara otomatis.
func orderMenu() {
	var keranjang [MAX_MENU]ItemOrder
	var jumlahItem int
	var totalBayar int
	var lanjut int
	var id int
	var index int
	var jumlahOrder int
	var i int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	jumlahItem = 0
	totalBayar = 0

	for {
		tampilkanSemuaMenu()

		fmt.Println("\nContoh ID menu yang ingin dipesan: 1")
		fmt.Println("Lihat ID pada tabel di atas.")
		fmt.Print("Masukkan ID menu yang ingin dipesan: ")
		fmt.Scan(&id)

		index = cariIndexByID(id)

		if index == -1 {
			fmt.Println("Menu tidak ditemukan.")
		} else if dataMenu[index].Stok <= 0 {
			fmt.Println("Maaf, menu ini sedang habis.")
		} else {
			fmt.Println("\nContoh jumlah order: 2")
			fmt.Printf("Stok saat ini: %d\n", dataMenu[index].Stok)
			fmt.Print("Masukkan jumlah yang ingin dipesan: ")
			fmt.Scan(&jumlahOrder)

			if jumlahOrder <= 0 {
				fmt.Println("Jumlah order tidak valid.")
			} else if jumlahOrder > dataMenu[index].Stok {
				fmt.Printf("Stok tidak cukup. Stok tersedia: %d\n", dataMenu[index].Stok)
			} else {
				keranjang[jumlahItem].NamaMenu = dataMenu[index].Nama
				keranjang[jumlahItem].Jumlah = jumlahOrder
				keranjang[jumlahItem].HargaSatuan = dataMenu[index].Harga
				keranjang[jumlahItem].Subtotal = jumlahOrder * dataMenu[index].Harga

				dataMenu[index].Stok = dataMenu[index].Stok - jumlahOrder
				totalBayar = totalBayar + keranjang[jumlahItem].Subtotal
				jumlahItem++

				fmt.Printf("Menu %s x%d berhasil ditambahkan ke keranjang.\n", dataMenu[index].Nama, jumlahOrder)
			}
		}

		fmt.Println("\nContoh pilihan: 1")
		fmt.Println("Keterangan: 1 = tambah menu lagi, 0 = selesai order")
		fmt.Print("Tambah menu lagi? ")
		fmt.Scan(&lanjut)

		if lanjut != 1 {
			break
		}
	}

	if jumlahItem == 0 {
		fmt.Println("Tidak ada item yang dipesan.")
		return
	}

	fmt.Println("\n=== Ringkasan Order ===")
	fmt.Println("+----+--------------------+--------+---------------+--------------+")
	fmt.Println("| No | Menu               | Jumlah | Harga Satuan  | Subtotal     |")
	fmt.Println("+----+--------------------+--------+---------------+--------------+")

	for i = 0; i < jumlahItem; i++ {
		fmt.Printf("| %-2d | %-18s | %-6d | Rp%-11d | Rp%-10d |\n",
			i+1,
			potongTeks(keranjang[i].NamaMenu, 18),
			keranjang[i].Jumlah,
			keranjang[i].HargaSatuan,
			keranjang[i].Subtotal,
		)
	}

	fmt.Println("+----+--------------------+--------+---------------+--------------+")
	fmt.Printf("| %-43s | Rp%-10d |\n", "TOTAL", totalBayar)
	fmt.Println("+----+--------------------+--------+---------------+--------------+")
	fmt.Println("Order berhasil dicatat.")
}

// Fungsi ini digunakan untuk mengubah stok menu secara langsung berdasarkan ID.
func ubahStokMenu() {
	var id int
	var index int
	var stokBaru int

	if jumlahMenu == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}

	tampilkanSemuaMenu()

	fmt.Println("\nContoh ID yang ingin diubah stoknya: 3")
	fmt.Println("Lihat ID pada tabel di atas.")
	fmt.Print("Masukkan ID menu yang ingin diubah stoknya: ")
	fmt.Scan(&id)

	index = cariIndexByID(id)

	if index == -1 {
		fmt.Println("Menu tidak ditemukan.")
		return
	}

	fmt.Printf("\nStok saat ini untuk %s: %d\n", dataMenu[index].Nama, dataMenu[index].Stok)
	fmt.Println("Contoh stok baru: 20")
	fmt.Println("Keterangan: 0 = habis, angka positif = jumlah stok tersedia")
	fmt.Print("Masukkan stok baru: ")
	fmt.Scan(&stokBaru)

	if stokBaru < 0 {
		fmt.Println("Stok tidak boleh negatif.")
		return
	}

	dataMenu[index].Stok = stokBaru

	fmt.Printf("Stok %s berhasil diubah menjadi %d.\n", dataMenu[index].Nama, dataMenu[index].Stok)
}
