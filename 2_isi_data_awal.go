package main


// Fungsi ini mengisi data awal agar program sudah memiliki contoh menu saat pertama dijalankan.
func isiDataAwal() {
	tambahDataAwal("Americano", "coffee", 18000, "espresso_air_panas", 10)
	tambahDataAwal("Cappuccino", "coffee", 25000, "espresso_susu_foam", 8)
	tambahDataAwal("Cafe_Latte", "coffee", 24000, "espresso_susu", 5)
	tambahDataAwal("Matcha_Latte", "non-coffee", 26000, "matcha_susu_gula", 7)
	tambahDataAwal("Chocolate_Ice", "non-coffee", 22000, "cokelat_susu_es", 6)
	tambahDataAwal("French_Fries", "snack", 20000, "kentang_garam_saus", 15)
	tambahDataAwal("Chicken_Sandwich", "food", 32000, "roti_ayam_selada_saus", 0)
}
