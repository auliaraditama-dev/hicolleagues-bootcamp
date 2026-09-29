// Utilitas Harian -- kelanjutan dari Sesi 2. Sekarang kita sudah
// belajar if-else, switch, for, DAN struct/slice/map (Sesi 3),
// jadi semua bagian yang tadi "sebagian" bisa dilengkapi.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Riwayat menyimpan satu catatan hasil kalkulator (struct, Sesi 3).
type Riwayat struct {
	A        float64
	B        float64
	Operator string
	Hasil    float64
}

func baca(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	teks, _ := in.ReadString('\n')
	return strings.TrimSpace(teks)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("=== Bagian 1: Ganjil atau Genap ===")
	angkaStr := baca(in, "Masukkan angka: ")
	angka, _ := strconv.Atoi(angkaStr)
	sisa := angka % 2
	fmt.Println("Sisa bagi 2:", sisa)
	// TODO 1: sekarang kita SUDAH belajar if-else (Sesi 3).
	// Tambahkan di bawah baris di atas: bila sisa == 0 cetak
	// "Genap", selain itu cetak "Ganjil".

	fmt.Println()
	fmt.Println("=== Bagian 2: Kalkulator (4 operator) ===")
	aStr := baca(in, "Angka pertama: ")
	bStr := baca(in, "Angka kedua: ")
	opStr := baca(in, "Operator (+ - * /): ")
	a, _ := strconv.ParseFloat(aStr, 64)
	b, _ := strconv.ParseFloat(bStr, 64)
	fmt.Println("Input diterima:", a, opStr, b)
	// TODO 2:
	//   1. Deklarasikan "var riwayat []Riwayat" (slice kosong) di
	//      luar switch, supaya masih ada dipakai TODO 4 di bawah.
	//   2. Pakai switch pada opStr untuk mendukung "+", "-", "*", "/".
	//      Untuk "/", tangani pembagian dengan nol (cetak pesan,
	//      jangan hitung, jangan tambah ke riwayat). Operator lain
	//      yang tidak dikenal -> cetak "Operator tidak valid".
	//   3. Setiap kali dapat hasil yang valid, cetak hasilnya dan
	//      tambahkan ke riwayat:
	//        riwayat = append(riwayat, Riwayat{A: a, B: b, Operator: opStr, Hasil: hasil})

	fmt.Println()
	fmt.Println("=== Bagian 3: Deret Angka ===")
	nStr := baca(in, "Cetak angka 1 sampai berapa? ")
	n, _ := strconv.Atoi(nStr)
	fmt.Println("Mencetak 1 sampai", n)
	// TODO 3: pakai for untuk mencetak angka 1 sampai n.

	fmt.Println()
	fmt.Println("=== Riwayat Kalkulator ===")
	// TODO 4: pakai for range pada riwayat (dari TODO 2) untuk
	// mencetak semua catatan. Format bebas, contoh:
	//   fmt.Printf("%.0f %s %.0f = %.2f\n", r.A, r.Operator, r.B, r.Hasil)
	// Bila riwayat kosong (mis. TODO 2 belum selesai), cetak
	// "Belum ada riwayat."
}
