# Informasi Aplikasi

 ini adalah Aplikasi Back end sederhana: sistem manajemen karyawan yang memungkinkan pengguna untuk membuat, memperbarui, menghapus, dan melihat data karyawan, departemen, dan posisi.

## Library yang Digunakan

- go get github.com/gofiber/fiber/v2
- go get github.com/spf13/viper
- go get gorm.io/gorm
- go get gorm.io/driver/postgres

## Cara Instalasi

1. Clone repository ini:

   ```sh

   git clone https://github.com/username/repo.git
   
   cd repo
   ```

2. Install dependencies menggunakan `go mod`:

   ```sh
   go mod tidy
   ```

3. Buat file konfigurasi `config.yaml` berdasarkan template yang disediakan.

4. Jalankan aplikasi:

   ```sh
   go run cmd/myapp/main.go
   ```

## Penambahan Library

Untuk menambahkan library baru, gunakan perintah `go get` diikuti dengan nama library. Contoh:

   ```sh
   go get github.com/username/library
   ```

Setelah menambahkan library, jangan lupa untuk menjalankan `go mod tidy` untuk merapikan dependencies:

   ```sh
   go mod tidy
   ```

## Routes

jalankan aplikasi `go run cmd/myapp/main.go`

lalu buka `http://localhost:3009/employees` atau sesuai dengan port yang anda setting pada config.yaml.

gunakan postman untuk test create data, update data, delete data, get data.

Gunakan `http://localhost:3009/employees` untuk test create data,
contoh request:

```JSON
{
    "name": "Kevin Joe",
    "age": 25,
    "gender": "male",
    "position_id": 1,
    "department_id": 2
}
```

`http://localhost:3009/employees/:id` untuk test update data, delete data, get data.

## TODO

Otentifikasi
