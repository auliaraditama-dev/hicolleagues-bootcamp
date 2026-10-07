// Aplikasi CLI Daftar Belanja — Hands-on Project 1 (Sesi 5).
//
// Cara menjalankan:
//   go run .
//
// File ini hanya mengurus TAMPILAN (menu, input, output). Semua
// logika (validasi, penyimpanan) ada di package belanja.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// TODO: begitu kamu mulai memanggil fungsi-fungsi di package
	// belanja (mulai dari TODO 1 di menuTambah), hapus tanda "//"
	// di baris bawah ini agar package-nya bisa dipakai.
	// "daftar-belanja/belanja"
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
		case "0":
			fmt.Println("Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

// baca menampilkan prompt, membaca satu baris input, lalu
// menghapus spasi dan enter di ujungnya. Dipakai oleh semua menu
// di bawah supaya tidak mengulang kode yang sama.
func baca(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	teks, _ := in.ReadString('\n')
	return strings.TrimSpace(teks)
}

// TODO 1: menuTambah
//   - Minta nama, harga, dan qty dari pengguna (pakai baca()).
//   - Ubah harga dan qty dari string ke angka. Tambahkan
//     "strconv" ke daftar import di atas, lalu pakai
//     strconv.ParseFloat(teks, 64) dan strconv.Atoi(teks).
//     Kalau gagal (err != nil), cetak pesan error dan return
//     (jangan lanjut ke belanja.Tambah).
//   - Panggil belanja.Tambah(nama, harga, qty).
//   - Jika error, cetak "Gagal: <pesan error>".
//   - Jika berhasil, cetak konfirmasi item yang ditambahkan
//     (boleh pakai %+v untuk lihat semua field dulu).
func menuTambah(in *bufio.Reader) {
	// TODO: implementasikan
	fmt.Println("(belum diimplementasikan)")
}

// TODO 2: menuLihat
//   - Ambil semua item dengan belanja.Semua().
//   - Jika kosong (len == 0), cetak "Belum ada item."
//   - Jika tidak kosong, cetak dalam bentuk tabel rapi, misalnya:
//     ID   Nama            Harga   Qty     Subtotal
//     1    Beras           15000     2       30000
//
//   Gunakan fmt.Printf dengan format lebar kolom, contoh:
//     fmt.Printf("%-4d %-15s %10.0f %5d %12.0f\n",
//         it.ID, it.Nama, it.Harga, it.Qty, it.Harga*float64(it.Qty))
func menuLihat() {
	// TODO: implementasikan
	fmt.Println("(belum diimplementasikan)")
}

// TODO 3: menuUbah
//   - Minta ID item dan qty baru dari pengguna.
//   - Ubah ID dan qty ke int dengan strconv.Atoi, tangani error
//     seperti di menuTambah.
//   - Panggil belanja.Ubah(id, qty).
//   - Cetak hasil (berhasil, atau "Gagal: <pesan error>").
func menuUbah(in *bufio.Reader) {
	// TODO: implementasikan
	fmt.Println("(belum diimplementasikan)")
}

// TODO 4: menuHapus
//   - Minta ID item yang ingin dihapus.
//   - Ubah ke int, tangani error.
//   - Panggil belanja.Hapus(id).
//   - Cetak hasil (berhasil, atau "Gagal: item tidak ditemukan").
func menuHapus(in *bufio.Reader) {
	// TODO: implementasikan
	fmt.Println("(belum diimplementasikan)")
}

// TODO 5: menuTotal
//   - Panggil belanja.Total().
//   - Cetak total harga belanja dengan format rapi, misalnya:
//     fmt.Printf("Total belanja: Rp%.0f\n", total)
func menuTotal() {
	// TODO: implementasikan
	fmt.Println("(belum diimplementasikan)")
}
