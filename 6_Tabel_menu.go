package main

import "fmt"

func tampilkanDaftar(daftar [MAX_MENU]Menu, jumlah int) {
	var i int

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