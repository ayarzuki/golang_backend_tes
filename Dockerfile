# Gunakan base image dari Golang
FROM golang:1.18

# Atur direktori kerja
WORKDIR /app

# Salin file go.mod dan go.sum ke direktori kerja
COPY go.mod go.sum ./

# Download semua dependensi
RUN go mod download

# Salin sumber kode ke direktori kerja
COPY . .

# Bangun aplikasi
RUN go build -o main .

# Ekspose port yang digunakan oleh aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]