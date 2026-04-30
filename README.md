# Belajar Golang Dasar

Repositori ini berisi dokumentasi dan kode latihan proses belajar saya dalam mendalami bahasa pemrograman Go (Golang). Proyek ini mencakup dasar-dasar sintaks hingga konsep yang lebih lanjut sebagai bagian dari persiapan transisi karir saya ke Backend Engineering.

## 🎯 Fokus Pembelajaran
Tujuan dari repositori ini adalah untuk memahami fondasi kuat Go, termasuk:
- Manajemen memori dasar menggunakan Pointer.
- Struktur data (Array, Slice, Map).
- Implementasi Interface dan Type Assertions.
- Penanganan Error (Custom Error & Panic).

## 🛠️ Daftar Materi & Struktur Folder
Repositori ini disusun secara modular berdasarkan topik pembelajaran:

### Dasar-Dasar & Tipe Data
- `01-helloworld` s/d `05-string`: Pengenalan dasar dan tipe data primitif.
- `06-variable` s/d `09-type-declaration`: Deklarasi dan manipulasi variabel.

### Logika & Control Flow
- `10-matematika` s/d `12-operasi-boolean`: Operator aritmatika dan logika.
- `16-if`, `17-switch`, `18-for`: Struktur kendali perulangan dan percabangan.

### Struktur Data & Collection
- `13-array`, `14-slice`, `15-map`: Mengelola kumpulan data di Go.

### Function & Modularization
- `20-function` s/d `23-closure`: Berbagai jenis fungsi, termasuk Anonymous dan Recursive function.
- `24-defer` & `25-panic`: Pengelolaan eksekusi fungsi dan *error recovery*.
- `35-import` & `36-package-initialization`: Cara kerja modularisasi dan package.

### Object Oriented (Go Style)
- `26-struct`, `27-interface`, `28-interface-kosong-any`: Implementasi cetak biru data dan kontrak.
- `31-pointer` s/d `34-pointer-method`: Pendalaman penggunaan memori dan method dengan receiver pointer.

### Error Handling
- `37-error` & `38-custom-error`: Standarisasi dan pembuatan error kustom.

## ⚙️ Cara Menggunakan
1. Pastikan Go terinstal di sistem Anda.
2. Clone repository:
   ```bash
   git clone [https://github.com/abdul452/golang-dasar.git](https://github.com/abdul452/golang-dasar.git)
3. Jalankan salah satu modul (contoh):
- cd 14-slice
- go run main.go
