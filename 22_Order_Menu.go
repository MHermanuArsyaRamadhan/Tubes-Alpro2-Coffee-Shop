package main

import "fmt"



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