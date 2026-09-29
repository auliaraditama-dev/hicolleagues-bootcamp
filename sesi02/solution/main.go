// Utilitas Harian -- versi solusi Sesi 2.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	angka, _ := strconv.Atoi(angkaStr)
	sisa := angka % 2
	fmt.Println("Sisa bagi 2:", sisa)

	fmt.Println()
	fmt.Println("=== Bagian 2: Kalkulator (hanya operator +) ===")
	aStr := baca(in, "Angka pertama: ")
	bStr := baca(in, "Angka kedua: ")
	a, _ := strconv.ParseFloat(aStr, 64)
	b, _ := strconv.ParseFloat(bStr, 64)
	hasil := a + b
	fmt.Printf("Hasil: %.2f\n", hasil)

	fmt.Println()
	fmt.Println("=== Bagian 3: Deret Angka ===")
	fmt.Println("(Menyusul di Sesi 3 -- butuh perulangan 'for')")
}
