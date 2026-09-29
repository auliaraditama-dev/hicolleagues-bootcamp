# Algoritma — Sesi 1

Rancangan algoritma untuk 3 studi kasus dari Sesi 1, ditulis dalam
PSEUDOCODE dulu — belum ada bahasa pemrograman. Kode Go sungguhan
baru mulai ditulis di Sesi 2, dan akan melanjutkan PERSIS algoritma
di bawah ini. Ini adalah awal dari proyek "Utilitas Harian" yang
akan tumbuh sampai Sesi 4.

## 1. Cek Ganjil atau Genap

```
input angka
sisa = angka MOD 2
if sisa == 0 then
    print "Genap"
else
    print "Ganjil"
end if
```

## 2. Kalkulator Sederhana

```
input a
input b
input operator

if operator == "+" then
    hasil = a + b
else if operator == "-" then
    hasil = a - b
else if operator == "*" then
    hasil = a * b
else if operator == "/" then
    if b == 0 then
        print "Tidak bisa dibagi nol"
    else
        hasil = a / b
    end if
end if

print hasil
```

## 3. Cetak Deret Angka 1 sampai N

```
input n
for i = 1 to n
    print i
end for
```

## Roadmap implementasi Go (kenapa tidak langsung lengkap di Sesi 2)

| Sesi | Yang bisa dibangun | Kenapa baru sebagian |
|---|---|---|
| **2** | Hitung `sisa` (ganjil/genap belum diputuskan), kalkulator KHUSUS operator "+" | Baru belajar variabel/tipe/operator — belum ada `if`/`switch` |
| **3** | Lengkapi ganjil/genap (`if-else`), kalkulator 4 operator (`switch`), deret angka (`for`), + riwayat hasil (`struct`+`slice`) | Sudah belajar kontrol aliran, perulangan, dan struktur data |
| **4** | Pindahkan semua logika ke package `utilitas`, jadi fungsi dengan `(hasil, error)`, `main.go` jadi menu tipis | Sudah belajar fungsi, multiple return, dan error handling |
| **5** | *(proyek terpisah, domain baru: daftar belanja)* — memakai pola persis dari Sesi 4: package + fungsi + error | Hands-on Project 1 |

Perhatikan pola di baris Sesi 4: `package` berisi fungsi murni,
`main.go` cuma memanggil dan menampilkan. Ini bukan kebetulan — ini
pola yang akan kamu pakai lagi di Sesi 5 (package `belanja`), dan
seterusnya di hampir semua backend Go pada bootcamp ini.
