// Utilitas Harian -- versi solusi Sesi 4.
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

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("=== Bagian 1: Ganjil atau Genap ===")
	angkaStr := baca(in, "Masukkan angka: ")
	angka, _ := strconv.Atoi(angkaStr)
	fmt.Println(utilitas.CekGanjilGenap(angka))

	fmt.Println()
	fmt.Println("=== Bagian 2: Kalkulator ===")
	aStr := baca(in, "Angka pertama: ")
	bStr := baca(in, "Angka kedua: ")
	opStr := baca(in, "Operator (+ - * /): ")
	a, _ := strconv.ParseFloat(aStr, 64)
	b, _ := strconv.ParseFloat(bStr, 64)

	var riwayat []utilitas.Riwayat
	hasil, err := utilitas.Hitung(a, b, opStr)
	if err != nil {
		fmt.Println("Gagal:", err)
	} else {
		fmt.Printf("Hasil: %.2f\n", hasil)
		riwayat = append(riwayat, utilitas.Riwayat{A: a, B: b, Operator: opStr, Hasil: hasil})
	}

	fmt.Println()
	fmt.Println("=== Bagian 3: Deret Angka ===")
	nStr := baca(in, "Cetak angka 1 sampai berapa? ")
	n, _ := strconv.Atoi(nStr)
	deret, err := utilitas.DeretAngka(n)
	if err != nil {
		fmt.Println("Gagal:", err)
	} else {
		for _, x := range deret {
			fmt.Println(x)
		}
	}

	fmt.Println()
	fmt.Println("=== Riwayat Kalkulator ===")
	if len(riwayat) == 0 {
		fmt.Println("Belum ada riwayat.")
	}
	for _, r := range riwayat {
		fmt.Printf("%.0f %s %.0f = %.2f\n", r.A, r.Operator, r.B, r.Hasil)
	}
}
