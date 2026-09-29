// Package utilitas berisi fungsi-fungsi MURNI (tidak ada
// fmt.Println/baca input di sini) -- hasil refactor dari Sesi 3.
// Ini pola yang sama persis yang dipakai lagi di package "belanja"
// pada Sesi 5: logika dipisah dari tampilan.
package utilitas

import "errors"

// Riwayat menyimpan satu catatan hasil kalkulator.
type Riwayat struct {
	A        float64
	B        float64
	Operator string
	Hasil    float64
}

// TODO 1: CekGanjilGenap
//   - Kembalikan string "Genap" atau "Ganjil" berdasarkan angka % 2.
//   - Ini murni fungsi (tidak ada fmt.Println di sini) -- pemanggil
//     (main.go) yang mencetak hasilnya.
func CekGanjilGenap(angka int) string {
	// TODO: implementasikan
	return ""
}

// TODO 2: Hitung
//   - Sama seperti switch di Sesi 3, tapi sekarang MENGEMBALIKAN
//     (float64, error) alih-alih mencetak langsung ke layar
//     (Sesi 4: error handling).
//   - Operator tidak dikenal -> kembalikan 0, errors.New("operator tidak valid")
//   - Pembagian dengan nol -> kembalikan 0, errors.New("tidak bisa dibagi nol")
func Hitung(a, b float64, operator string) (float64, error) {
	// TODO: implementasikan
	return 0, errors.New("belum diimplementasikan")
}

// TODO 3: DeretAngka
//   - Kembalikan []int berisi 1 sampai n (BUKAN mencetak langsung).
//   - Bila n <= 0, kembalikan nil dan
//     errors.New("n harus lebih besar dari 0").
func DeretAngka(n int) ([]int, error) {
	// TODO: implementasikan
	return nil, errors.New("belum diimplementasikan")
}
