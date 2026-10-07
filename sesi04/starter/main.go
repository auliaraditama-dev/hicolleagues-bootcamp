// Utilitas Harian -- versi Sesi 4: logika dipindah ke package
// utilitas/, main.go jadi menu tipis yang memanggilnya. Bandingkan
// dengan Sesi 3: pola ini PERSIS yang dipakai lagi di Sesi 5
// (package belanja + main.go).
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"utilitas-harian/utilitas"
)

func baca(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	teks, _ := in.ReadString('\n')
	return strings.TrimSpace(teks)
}

func test() {
	fmt.Println("coba import")
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("=== Bagian 1: Ganjil atau Genap ===")
	angkaStr := baca(in, "Masukkan angka: ")
	angka, _ := strconv.Atoi(angkaStr)
	// Contoh yang SUDAH LENGKAP -- pelajari polanya sebelum
	// mengerjakan TODO 2-4 di bawah, yang bentuknya serupa:
	// panggil fungsi dari package utilitas, lalu cetak hasilnya.
	fmt.Println(utilitas.CekGanjilGenap(angka))

	fmt.Println()
	fmt.Println("=== Bagian 2: Kalkulator ===")
	aStr := baca(in, "Angka pertama: ")
	bStr := baca(in, "Angka kedua: ")
	opStr := baca(in, "Operator (+ - * /): ")
	a, _ := strconv.ParseFloat(aStr, 64)
	b, _ := strconv.ParseFloat(bStr, 64)
	fmt.Println("Input diterima:", a, opStr, b)
	var riwayat []utilitas.Riwayat
	// TODO 2: panggil utilitas.Hitung(a, b, opStr).
	//   - Bila error, cetak "Gagal:", err.
	//   - Bila sukses, cetak hasilnya (fmt.Printf dengan %.2f), dan
	//     tambahkan ke riwayat:
	//       riwayat = append(riwayat, utilitas.Riwayat{A: a, B: b, Operator: opStr, Hasil: hasil})
	_ = riwayat // TODO: hapus baris ini setelah riwayat dipakai di TODO 2

	fmt.Println()
	fmt.Println("=== Bagian 3: Deret Angka ===")
	nStr := baca(in, "Cetak angka 1 sampai berapa? ")
	n, _ := strconv.Atoi(nStr)
	// TODO 3: panggil utilitas.DeretAngka(n).
	//   - Bila error, cetak pesannya.
	//   - Bila sukses, cetak semua angkanya dengan for range.
	fmt.Println("Diminta mencetak sampai", n)

	fmt.Println()
	fmt.Println("=== Riwayat Kalkulator ===")
	// TODO 4: cetak semua isi riwayat (dari TODO 2) dengan for
	// range. Bila kosong, cetak "Belum ada riwayat."
}
