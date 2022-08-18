
# Stage pertama digunakan untuk mendownload package kebutuhan aplikasi dan mem-build aplikasi ke binary file. 
# Disini saya juga menggunakan aplikasi UPX yang bertujuan untuk mengecilkan lagi/kompress file binary
FROM golang:1.18-alpine
RUN apk add --update upx
WORKDIR /src
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -a && upx -q9 main

# Stage kedua saya memindahkan file binary tersebut ke image alpine. stage kedua ini yang akan jadi wadah jalannya
# aplikasi. Disini saya menggunakan image alpine, karena sizenya yang cukup kecil.
FROM alpine:3.12
WORKDIR /app
COPY --from=0 /src/main .
EXPOSE 8080
ENTRYPOINT ["./main"]

# Dengan multistage, docker image akan relatif lebih kecil karena pada image tidak ada file2 code dan yang dijalankan
# di container hanyalah hasil build dari apps saja