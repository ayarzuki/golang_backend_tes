# REST API Product dengan Golang dan MySQL

## Arsitektur
Aplikasi ini menggunakan arsitektur berbasis package yang memisahkan antara konfigurasi, model, dan handler. Alasan menggunakan arsitektur ini adalah untuk memudahkan dalam pengembangan, pemeliharaan, dan mengikuti prinsip Single Responsibility. Arsitektur yang digunakan adalah arsitektur berbasis package yang modular dan mudah dipelihara. Salah satu pendekatan yang saya gunakan adalah Clean Architecture, yang memisahkan komponen berdasarkan tanggung jawab mereka.
### Keuntungan Menggunakan Clean Architecture
Berikut adalah beberapa keuntungan menggunakan Clean Architecture dalam proyek ini:
1. **Pemisahan Tanggung Jawab**: Clean Architecture memisahkan tanggung jawab ke dalam komponen yang jelas dan terisolasi, seperti Model, Handlers, Config, Repositories, dan Services. Hal ini memudahkan pemeliharaan kode dan memungkinkan pengembangan lebih lanjut.

2. **Mudah diuji**: Dengan memisahkan logika bisnis dari infrastruktur, Anda dapat menguji komponen secara terpisah dan memastikan bahwa sistem Anda berfungsi dengan benar.

3. **Fleksibilitas**: Clean Architecture memungkinkan Anda untuk menggantikan atau mengubah komponen sistem tanpa perlu mengubah bagian lainnya secara signifikan. Misalnya, jika Anda ingin mengganti database yang digunakan, Anda hanya perlu mengubah bagian Config dan Repositories tanpa harus mengubah logika bisnis dalam Handlers atau Services.

4. **Skalabilitas**: Dalam proyek RESTful API, Clean Architecture memungkinkan Anda untuk menambahkan lebih banyak fitur atau endpoint dengan mudah, karena Anda dapat menambahkan komponen yang diperlukan tanpa perlu mengubah bagian lain dari sistem.

## Struktur Folder
```Proyek ini memiliki struktur folder berikut:
.
├── main.go
├── config
│ └── db.go
├── models
│ └── product.go
├── handlers
│ ├── product_handlers.go
│ ├── product_handlers_test.go
│ └── .env  
└── README.md
```

- `main.go` berfungsi untuk menjalankan aplikasi dan menentukan routing.
- Folder `config` berisi konfigurasi koneksi database dan automigrasi model.
- Folder `models` berisi struktur model data, dalam hal ini adalah `Product`.
- Folder `handlers` berisi logika bisnis untuk setiap endpoint, dalam hal ini adalah penambahan product dan menampilkan list product dengan sorting. Dalam proyek ini, kita memiliki `AddProduct` dan `ListProducts`.

## Cara Menjalankan Aplikasi secara lokal
1. Pastikan Golang, MySQL atau PostgreSQL, dan library yang diperlukan sudah terinstall.
2. Sesuaikan konfigurasi koneksi database pada file `config/db.go`. File .env terdapat dalam folder handlers.
3. Jalankan aplikasi dengan perintah `go run main.go`.
4. Aplikasi akan berjalan pada `localhost:8080` dan siap menerima request melalui endpoint yang tersedia.


## Pengujian API
### Tambahkan Produk
Untuk menambahkan produk, gunakan perintah curl berikut:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Product A", "price": 1000, "description": "Deskripsi Product A", "quantity": 10}' http://localhost:8080/api/products
```

Keterangan:
    - `-X POST` : Menggunakan method POST.
    - `-H` "Content-Type: application/json" : Menggunakan header dengan tipe konten JSON.
    - `-d '{"name": "Product A", "price": 1000, "description": "Deskripsi Product A", "quantity": 10}'` : Data yang akan dikirimkan dalam bentuk JSON.
    - `http://localhost:8080/api/products` : Endpoint API yang akan dipanggil.

![Alt Text](https://i.postimg.cc/P5GL546R/POST.png)

### Dapatkan Daftar Produk dengan Pengurutan Berdasarkan Harga Termurah

Untuk mendapatkan daftar produk dengan pengurutan berdasarkan harga termurah, gunakan perintah curl berikut:

```bash
curl -X GET http://localhost:8080/api/products?sort=price
```

Keterangan:
    - `-X GET` : Menggunakan method GET.
    - `http://localhost:8080/api/products?sort=price` : Endpoint API yang akan dipanggil, dengan parameter `sort` yang bernilai `price` untuk mengurutkan berdasarkan `harga termurah`.

![Alt Text](https://i.postimg.cc/qv8Nh3SM/termurah.png)

### Dapatkan Daftar Produk dengan Pengurutan Berdasarkan Harga Termahal

Untuk mendapatkan daftar produk dengan pengurutan berdasarkan harga termahal, gunakan perintah curl berikut:

```bash
curl -X GET http://localhost:8080/api/products?sort=-price
```
Keterangan:
    - `-X GET` : Menggunakan method GET.
    - `http://localhost:8080/api/products?sort=-price` : Endpoint API yang akan dipanggil, dengan parameter `sort` yang bernilai `-price` untuk mengurutkan berdasarkan `harga termahal`.

![Alt Text](https://i.postimg.cc/mrjhVGTv/termahal.png)

### Dapatkan Daftar Produk dengan Pengurutan Berdasarkan Nama (A-Z)

Untuk mendapatkan daftar produk dengan pengurutan berdasarkan nama dari A-Z, gunakan perintah curl berikut:

```bash
curl -X GET http://localhost:8080/api/products?sort=name
```
Keterangan:
    - `-X GET` : Menggunakan method GET.
    - `http://localhost:8080/api/products?sort=name` : Endpoint API yang akan dipanggil, dengan parameter `sort` yang bernilai `name` untuk mengurutkan berdasarkan nama dari `A-Z`.

![Alt Text](https://i.postimg.cc/5yTXvqfj/a-to-z.png)

### Dapatkan Daftar Produk dengan Pengurutan Berdasarkan Nama (Z-A)

Untuk mendapatkan daftar produk dengan pengurutan berdasarkan nama dari Z-A, gunakan perintah curl berikut:

```bash
curl -X GET http://localhost:8080/api/products?sort=-name
```
Keterangan:
    - `-X GET` : Menggunakan method GET.
    - `http://localhost:8080/api/products?sort=-name` : Endpoint API yang akan dipanggil, dengan parameter sort yang bernilai `-name` untuk mengurutkan berdasarkan nama dari `Z-A`.

![Alt Text](https://i.postimg.cc/xdncKdVG/z-to-a.png)