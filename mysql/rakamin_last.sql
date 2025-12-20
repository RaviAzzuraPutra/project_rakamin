-- 1. Buat Database
CREATE DATABASE project_rakamin;
USE project_rakamin;

-- 2. Tabel User
CREATE TABLE User (
    id CHAR(36) PRIMARY KEY,
    nama VARCHAR(255),
    kata_sandi VARCHAR(255),
    notelp VARCHAR(255) UNIQUE,
    tanggal_lahir DATE,
    jenis_kelamin VARCHAR(255),
    tentang TEXT,
    pekerjaan VARCHAR(255),
    email VARCHAR(255),
    id_provinsi VARCHAR(255),
    id_kota VARCHAR(255),
    isAdmin BOOLEAN,
    updated_at DATE,
    created_at DATE
);

-- 3. Tabel Alamat
CREATE TABLE alamat (
    id CHAR(36) PRIMARY KEY,
    id_user CHAR(36),
    judul_alamat VARCHAR(255),
    nama_penerima VARCHAR(255),
    no_telp VARCHAR(255),
    detail_alamat VARCHAR(255),
    updated_at DATE,
    created_at DATE,
    FOREIGN KEY (id_user) REFERENCES User(id)
);

-- 4. Tabel Toko
CREATE TABLE toko (
    id CHAR(36) PRIMARY KEY,
    id_user CHAR(36),
    nama_toko VARCHAR(255),
    url_foto VARCHAR(255),
    updated_at DATE,
    created_at DATE,
    FOREIGN KEY (id_user) REFERENCES User(id)
);

-- 5. Tabel Category
CREATE TABLE category (
    id CHAR(36) PRIMARY KEY,
    nama_category VARCHAR(255),
    created_at DATE,
    updated_at DATE
);

-- 6. Tabel Produk
CREATE TABLE produk (
    id CHAR(36) PRIMARY KEY,
    nama_produk VARCHAR(255),
    slug VARCHAR(255),
    harga_reseller VARCHAR(255),
    harga_konsumen VARCHAR(255),
    stok INT,
    deskripsi TEXT,
    created_at DATE,
    updated_at DATE,
    id_toko CHAR(36),
    id_category CHAR(36),
    FOREIGN KEY (id_toko) REFERENCES toko(id),
    FOREIGN KEY (id_category) REFERENCES category(id)
);

-- 7. Tabel Foto Produk
CREATE TABLE foto_produk (
    id CHAR(36) PRIMARY KEY,
    id_produk CHAR(36),
    url VARCHAR(255),
    updated_at DATE,
    created_at DATE,
    FOREIGN KEY (id_produk) REFERENCES produk(id)
);

-- 8. Tabel Log Produk
CREATE TABLE log_produk (
    id CHAR(36) PRIMARY KEY,
    id_produk CHAR(36),
    nama_produk VARCHAR(255),
    slug VARCHAR(255),
    harga_reseller VARCHAR(255),
    harga_konsumen VARCHAR(255),
    deskripsi TEXT,
    created_at DATE,
    updated_at DATE,
    id_toko CHAR(36),
    id_category CHAR(36),
    FOREIGN KEY (id_toko) REFERENCES toko(id),
    FOREIGN KEY (id_category) REFERENCES category(id)
);

-- 9. Tabel Trx (Transaksi)
CREATE TABLE trx (
    id CHAR(36) PRIMARY KEY,
    id_user CHAR(36),
    alamat_pengiriman CHAR(36),
    harga_total INT,
    kode_invoice VARCHAR(255),
    method_bayar VARCHAR(255),
    updated_at DATE,
    created_at DATE,
    FOREIGN KEY (id_user) REFERENCES User(id),
    FOREIGN KEY (alamat_pengiriman) REFERENCES alamat(id)
);

-- 10. Tabel Detail Trx
CREATE TABLE detail_trx (
    id CHAR(36) PRIMARY KEY,
    id_trx CHAR(36),
    id_log_produk CHAR(36),
    id_toko CHAR(36),
    kuantitas INT,
    harga_total INT,
    updated_at DATE,
    created_at DATE,
    FOREIGN KEY (id_trx) REFERENCES trx(id),
    FOREIGN KEY (id_log_produk) REFERENCES log_produk(id),
    FOREIGN KEY (id_toko) REFERENCES toko(id)
);