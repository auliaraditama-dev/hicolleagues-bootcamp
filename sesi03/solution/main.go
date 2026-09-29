// Utilitas Harian -- versi solusi Sesi 3.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	if sisa == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}

	fmt.Println()
	fmt.Println("=== Bagian 2: Kalkulator (4 operator) ===")
	aStr := baca(in, "Angka pertama: ")
	bStr := baca(in, "Angka kedua: ")
	opStr := baca(in, "Operator (+ - * /): ")
	a, _ := strconv.ParseFloat(aStr, 64)
	b, _ := strconv.ParseFloat(bStr, 64)

	var riwayat []Riwayat
	var hasil float64
	valid := true

	switch opStr {
	case "+":
		hasil = a + b
	case "-":
		hasil = a - b
	case "*":
		hasil = a * b
	case "/":
		if b == 0 {
			fmt.Println("Tidak bisa dibagi nol")
			valid = false
		} else {
			hasil = a / b
		}
	default:
		fmt.Println("Operator tidak valid")
		valid = false
	}

	if valid {
		fmt.Printf("Hasil: %.2f\n", hasil)
		riwayat = append(riwayat, Riwayat{A: a, B: b, Operator: opStr, Hasil: hasil})
	}

	fmt.Println()
	fmt.Println("=== Bagian 3: Deret Angka ===")
	nStr := baca(in, "Cetak angka 1 sampai berapa? ")
	n, _ := strconv.Atoi(nStr)
	for i := 1; i <= n; i++ {
		fmt.Println(i)
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
