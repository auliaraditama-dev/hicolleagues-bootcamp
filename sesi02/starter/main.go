// Utilitas Harian -- kelanjutan dari algoritma.md (Sesi 1).
// Sesi ini BARU bisa memakai variabel, tipe data, dan operator --
// belum ada if/switch/for (itu bahan Sesi 3), jadi sebagian
// algoritma SENGAJA belum lengkap. Lihat komentar di tiap bagian.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// TODO: tambahkan "strconv" begitu kamu mulai mengubah teks
	// jadi angka di TODO 1 dan TODO 2 di bawah.
)

func baca(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	teks, _ := in.ReadString('\n')
	return strings.TrimSpace(teks)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("=== Bagian 1: Ganjil atau Genap (sebagian) ===")
	angkaStr := baca(in, "Masukkan angka: ")
	fmt.Println("Kamu memasukkan:", angkaStr)
	// TODO 1: ubah angkaStr jadi int (strconv.Atoi), lalu hitung
	// sisa bagi 2 dengan operator %, dan cetak sisanya. Kita BELUM
	// bisa mencetak "Ganjil"/"Genap" di sini karena itu butuh if
	// (Sesi 3) -- sengaja, akan dilanjutkan nanti.

	fmt.Println()
	fmt.Println("=== Bagian 2: Kalkulator (hanya operator +) ===")
	aStr := baca(in, "Angka pertama: ")
	bStr := baca(in, "Angka kedua: ")
	fmt.Println("Input diterima:", aStr, "dan", bStr)
	// TODO 2: ubah aStr dan bStr jadi float64 (strconv.ParseFloat),
	// jumlahkan, cetak hasilnya dengan format 2 angka di belakang
	// koma (%.2f). Operator lain (-, *, /) menyusul di Sesi 3
	// setelah belajar switch.

	fmt.Println()
	fmt.Println("=== Bagian 3: Deret Angka ===")
	fmt.Println("(Menyusul di Sesi 3 -- butuh perulangan 'for')")
}
