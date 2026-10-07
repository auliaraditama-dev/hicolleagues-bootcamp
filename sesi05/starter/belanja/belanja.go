// Package belanja berisi model data dan operasi CRUD untuk aplikasi
// daftar belanja. Semua fungsi di package ini TIDAK berurusan dengan
// tampilan (itu tugas main.go) — package ini murni logika, sesuai
// pola repository/service sederhana yang dipakai di Sesi 4-5.
package belanja

import (
	"errors"
	"fmt"
)

// Item merepresentasikan satu barang belanja.
type Item struct {
	ID    int
	Nama  string
	Harga float64
	Qty   int
}

// items dan nextID menyimpan data di memori (belum pakai database;
// database baru dibahas di Fase 2).
var items []Item
var nextID = 1

// TODO 1: Tambah
//   - Validasi: nama tidak boleh kosong, harga tidak boleh negatif,
//     qty harus lebih besar dari 0. Jika salah satu gagal, kembalikan
//     Item{} dan errors.New("pesan yang jelas").
//   - Jika valid: buat Item baru dengan ID = nextID, tambahkan ke
//     items (pakai append), naikkan nextID, lalu kembalikan item
//     baru dan nil.
//
// Contoh pemanggilan yang diharapkan:
//
//	it, err := Tambah("Beras", 15000, 2)
//	// it == Item{ID: 1, Nama: "Beras", Harga: 15000, Qty: 2}, err == nil
func Tambah(nama string, harga float64, qty int) (Item, error) {
	// TODO: implementasikan
	return Item{}, errors.New("belum diimplementasikan")
}

// TODO 2: Semua
//   - Kembalikan slice items apa adanya.
func Semua() []Item {
	// TODO: implementasikan
	return nil
}

// TODO 3: Hapus
//   - Cari item dengan ID yang cocok memakai for range.
//   - Jika ketemu, hapus dari slice dengan pola (ingat Sesi 3):
//     items = append(items[:i], items[i+1:]...)
//     lalu return nil.
//   - Jika sampai akhir loop tidak ketemu, kembalikan error dengan
//     fmt.Errorf("item %d tidak ditemukan", id).
func Hapus(id int) error {
	// TODO: implementasikan
	return fmt.Errorf("belum diimplementasikan")
}

// TODO 4: Ubah
//   - Cari item dengan ID yang cocok.
//   - PENTING: ubah lewat items[i].Qty = qty (bukan lewat variabel
//     hasil range, karena itu cuma salinan — lihat catatan Sesi 3
//     tentang "jebakan range").
//   - Jika ketemu, set Qty, return nil.
//   - Jika tidak ketemu, kembalikan error seperti di Hapus.
func Ubah(id, qty int) error {
	// TODO: implementasikan
	return fmt.Errorf("belum diimplementasikan")
}

// TODO 5: Total
//   - Jumlahkan Harga * Qty (dikonversi ke float64) dari semua item.
//   - Pola akumulator seperti Sesi 1: total dimulai dari 0.
func Total() float64 {
	// TODO: implementasikan
	return 0
}
