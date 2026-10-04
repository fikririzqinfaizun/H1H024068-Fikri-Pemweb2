<?php

use App\Http\Controllers\Api\AuthController;
use App\Http\Controllers\Api\MahasiswaController;
use App\Http\Controllers\Api\MatakuliahController;
use App\Http\Controllers\Api\ProgramStudiController;
use Illuminate\Support\Facades\Route;

Route::get('/status', function () {
    return response()->json([
        'sukses' => true,
        'pesan' => 'API Pemweb II aktif',
        'waktu' => now()->toIso8601String(),
    ]);
});

Route::get('/tes', function () {
    return response()->json([
        'tes' => 'BERHASIL',
    ]);
});

/*
|--------------------------------------------------------------------------
| Autentikasi
|--------------------------------------------------------------------------
*/

Route::post('/auth/register', [AuthController::class, 'register']);

Route::post('/auth/login', [AuthController::class, 'login'])
    ->middleware('throttle:5,1');

Route::middleware('auth:sanctum')->group(function () {

    Route::get('/auth/profil', [AuthController::class, 'profil']);
    Route::post('/auth/logout', [AuthController::class, 'logout']);
    Route::post('/auth/logout-semua', [AuthController::class, 'logoutSemua']);

    // Ubah password
    Route::put('/auth/password', [AuthController::class, 'ubahPassword']);

    /*
    |--------------------------------------------------------------------------
    | Mahasiswa
    |--------------------------------------------------------------------------
    */

    Route::get('/mahasiswa', [MahasiswaController::class, 'index']);
    Route::get('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'show']);

    Route::middleware('ability:mahasiswa:tulis')->group(function () {
        Route::post('/mahasiswa', [MahasiswaController::class, 'store']);
        Route::put('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'update']);
        Route::patch('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'update']);

        // Hanya admin yang dapat menghapus mahasiswa
        Route::delete('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'destroy'])
            ->middleware('peran.admin');
    });

    /*
    |--------------------------------------------------------------------------
    | Mata Kuliah
    |--------------------------------------------------------------------------
    */

    Route::apiResource('matakuliah', MatakuliahController::class);

    /*
    |--------------------------------------------------------------------------
    | Program Studi
    |--------------------------------------------------------------------------
    */

    Route::get(
        '/program-studi/{id}/mahasiswa',
        [ProgramStudiController::class, 'mahasiswa']
    );
});