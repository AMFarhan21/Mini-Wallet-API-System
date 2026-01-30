## Arsitektur
Aplikasi ini menggunakan Clean Architecture:

#### Handler(HTTP layer) 
- Bertanggungjawab untuk menerima request, melakukan validasi input, dan mengembalikan response
#### Service (Business layer berfungsi)
- Sebagai tempat logika bisnis, seperti validasi saldo, proses transfer, serta koordinasi antar repository
#### Repository (Data access layer)
- Berfungsi untuk berkomunikasi dengan database

## Money Handling
- Saya menggunakan bigint atau int64 untuk balance dikarenakan untuk mencegah overflow saat menangani transaksi bervolume tinggi atau mata uang dengan satuan kecil
- Mencegah kesalhaan pembulatan dari float yang sangat krusial di sistem finansial.
- Penggunaan bigint juga untuk mencegah kesalahan pembulatan desimal. Bigint menyimpan angka bilangan bulat misalkan saldo itu 1000,500 disimpan sebagai 1000500, sehingga perhitungan tetap presisi

## COncurrency Handling
Saya menggunakan GORM locking clauses agar
- Wallet hanya bisa dimodifikasi satu transaksi dalam satu waktu
- Mencegah dua kali pembayaran dalam sekali aksi

## Postman Api documentation
https://documenter.getpostman.com/view/45402659/2sBXVo97rg

## Cara Menjalankan Project
1. Clone repo ini
2. Buat .env dan sesuaikan dengan .env.example
3. Jalankan docker
````bash
docker compose up -d
````
4. Buka
````bash
docker exec -it postgres psql -U {POSTGRES_USER} -d {POSTGRES_DB}
````
5. Jalankan SQL schema dan seeder
````sql
CREATE TABLE users (
    id serial primary key,
    name varchar(255),
    email varchar(255) unique
);

CREATE TABLE wallets (
    id serial primary key,
    user_id int not null references users(id),
    balance bigint not null default 0
);

CREATE TABLE transactions (
    id serial primary key,
    wallet_id int not null references wallets(id),
    type varchar(20),
    amount bigint,
    reference_id varchar(255),
    created_at timestamp default current_timestamp
);

INSERT INTO users (name, email) VALUES
('andi', 'andi@gmail.com'),
('farhan', 'farhan@gmail.com'),
('asep', 'asep@gmail.com'),
('budi', 'budi@gmail.com');

INSERT INTO wallets (user_id, balance) VALUES
(1, 0),
(2, 0),
(3, 0),
(4, 0);
````

6. Jalankan server
````bash
go run app/gin-server/main.go
````
