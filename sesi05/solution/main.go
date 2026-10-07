// Aplikasi CLI Daftar Belanja — versi solusi lengkap (Sesi 5).
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"daftar-belanja/belanja"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("=== Daftar Belanja ===")
		fmt.Println("1. Tambah item")
		fmt.Println("2. Lihat semua item")
		fmt.Println("3. Ubah jumlah (qty)")
		fmt.Println("4. Hapus item")
		fmt.Println("5. Total harga")
		fmt.Println("6. Cari item (bonus)")
		fmt.Println("0. Keluar")
		pilihan := baca(in, "Pilih menu: ")

		switch pilihan {
		case "1":
			menuTambah(in)
		case "2":
			menuLihat()
		case "3":
			menuUbah(in)
		case "4":
			menuHapus(in)
		case "5":
			menuTotal()
		case "6":
			menuCari(in)
		case "0":
			fmt.Println("Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

func baca(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	teks, _ := in.ReadString('\n')
	return strings.TrimSpace(teks)
}

func menuTambah(in *bufio.Reader) {
	nama := baca(in, "Nama item: ")

	hargaStr := baca(in, "Harga: ")
	harga, err := strconv.ParseFloat(hargaStr, 64)
	if err != nil {
		fmt.Println("Harga harus berupa angka.")
		return
	}

	qtyStr := baca(in, "Qty: ")
	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		fmt.Println("Qty harus berupa angka bulat.")
		return
	}

	it, err := belanja.Tambah(nama, harga, qty)
	if err != nil {
		fmt.Println("Gagal:", err)
		return
	}
	fmt.Printf("Ditambahkan: [%d] %s x%d @Rp%.0f\n", it.ID, it.Nama, it.Qty, it.Harga)
}

func menuLihat() {
	semua := belanja.Semua()
	if len(semua) == 0 {
		fmt.Println("Belum ada item.")
		return
	}

	fmt.Printf("%-4s %-15s %10s %5s %12s\n", "ID", "Nama", "Harga", "Qty", "Subtotal")
	for _, it := range semua {
		subtotal := it.Harga * float64(it.Qty)
		fmt.Printf("%-4d %-15s %10.0f %5d %12.0f\n", it.ID, it.Nama, it.Harga, it.Qty, subtotal)
	}
}

func menuUbah(in *bufio.Reader) {
	idStr := baca(in, "ID item: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus berupa angka.")
		return
	}

	qtyStr := baca(in, "Qty baru: ")
	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		fmt.Println("Qty harus berupa angka bulat.")
		return
	}

	if err := belanja.Ubah(id, qty); err != nil {
		fmt.Println("Gagal:", err)
		return
	}
	fmt.Println("Qty berhasil diubah.")
}

func menuHapus(in *bufio.Reader) {
	idStr := baca(in, "ID item yang dihapus: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID harus berupa angka.")
		return
	}

	if err := belanja.Hapus(id); err != nil {
		fmt.Println("Gagal:", err)
		return
	}
	fmt.Println("Item berhasil dihapus.")
}

func menuTotal() {
	total := belanja.Total()
	fmt.Printf("Total belanja: Rp%.0f\n", total)
}

func menuCari(in *bufio.Reader) {
	kata := baca(in, "Cari nama mengandung: ")
	hasil := belanja.CariByNama(kata)
	if len(hasil) == 0 {
		fmt.Println("Tidak ada item yang cocok.")
		return
	}
	for _, it := range hasil {
		fmt.Printf("[%d] %s x%d @Rp%.0f\n", it.ID, it.Nama, it.Qty, it.Harga)
	}
}
