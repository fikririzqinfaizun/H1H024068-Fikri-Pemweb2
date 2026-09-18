<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MahasiswaWebController;

Route::get('/mahasiswa-data', [MahasiswaWebController::class, 'index'])
    ->name('mahasiswa.data');
Route::get('/mahasiswa/{id}', [MahasiswaWebController::class, 'show'])
    ->name('mahasiswa.detail');
Route::get('/mahasiswa-top', [MahasiswaWebController::class, 'topTeknikKomputer'])
    ->name('mahasiswa.top');
Route::get('/mahasiswa/{id}/matakuliah', [MahasiswaWebController::class, 'matakuliah'])
    ->name('mahasiswa.matakuliah');