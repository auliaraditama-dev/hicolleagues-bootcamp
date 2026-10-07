// Package belanja berisi model data dan operasi CRUD untuk aplikasi
// daftar belanja (versi solusi lengkap).
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

var items []Item
var nextID = 1

// Tambah memvalidasi lalu menyimpan item baru.
func Tambah(nama string, harga float64, qty int) (Item, error) {
	if nama == "" {
		return Item{}, errors.New("nama tidak boleh kosong")
	}
	if harga < 0 {
		return Item{}, errors.New("harga tidak boleh negatif")
	}
	if qty <= 0 {
		return Item{}, errors.New("qty harus lebih besar dari 0")
	}

	it := Item{
		ID:    nextID,
		Nama:  nama,
		Harga: harga,
		Qty:   qty,
	}
	nextID++
	items = append(items, it)
	return it, nil
}

// Semua mengembalikan seluruh item yang tersimpan.
func Semua() []Item {
	return items
}

// Hapus menghapus item dengan ID yang cocok.
func Hapus(id int) error {
	for i, it := range items {
		if it.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item %d tidak ditemukan", id)
}

// Ubah mengubah qty item dengan ID yang cocok.
func Ubah(id, qty int) error {
	for i := range items {
		if items[i].ID == id {
			items[i].Qty = qty
			return nil
		}
	}
	return fmt.Errorf("item %d tidak ditemukan", id)
}

// Total menjumlahkan harga * qty seluruh item.
func Total() float64 {
	total := 0.0
	for _, it := range items {
		total += it.Harga * float64(it.Qty)
	}
	return total
}

// CariByNama adalah fungsi bonus (lihat README) untuk mencari item
// berdasarkan potongan nama, tidak peka huruf besar/kecil.
func CariByNama(kata string) []Item {
	var hasil []Item
	for _, it := range items {
		if containsFold(it.Nama, kata) {
			hasil = append(hasil, it)
		}
	}
	return hasil
}

func containsFold(s, sub string) bool {
	sLower, subLower := toLower(s), toLower(sub)
	return index(sLower, subLower) >= 0
}
