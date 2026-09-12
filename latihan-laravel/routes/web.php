<?php

use App\Http\Controllers\MatakuliahController;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MahasiswaController;

Route::get('/salam', function () {
    return 'Selamat datang di Pemrograman Web II';
});

Route::get('/mahasiswa/{nim}', function (string $nim) {
    return 'NAMA SAYA FIKRI DENGAN NIM ' . $nim;
});

Route::get('/semester/{angka}', function (int $angka) {
    return '6 ' . $angka;
})->whereNumber('6');

Route::get('/data-mahasiswa', [MahasiswaController::class, 'index'])
    ->name('mahasiswa.index');

Route::get('/data-mahasiswa/{nim}', [MahasiswaController::class, 'show'])
    ->name('mahasiswa.show');

Route::get('/cari-mahasiswa', [MahasiswaController::class, 'cari']);

Route::get('/data-matakuliah', [MatakuliahController::class, 'index'])
    ->name('matakuliah.index');

Route::get('/data-matakuliah/{kode}', [MatakuliahController::class, 'show'])
    ->name('matakuliah.show');