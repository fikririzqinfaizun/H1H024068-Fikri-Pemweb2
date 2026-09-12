<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $daftarMatakuliah = [
            [
                'kode' => 'TK101',
                'nama' => 'Pemrograman Web II',
                'sks' => 3
            ],
            [
                'kode' => 'TK102',
                'nama' => 'Internet of Things',
                'sks' => 3
            ],
            [
                'kode' => 'TK103',
                'nama' => 'Sistem Operasi',
                'sks' => 3
            ],
            [
                'kode' => 'TK104',
                'nama' => 'Jaringan Komputer',
                'sks' => 2
            ],
            [
                'kode' => 'TK105',
                'nama' => 'Manajemen Proyek',
                'sks' => 2
            ],
        ];

        // Mengambil kata kunci dari query string ?q=
        $kataKunci = $request->query('q', '');

        // Melakukan pencarian jika kata kunci tidak kosong
        if ($kataKunci !== '') {
            $daftarMatakuliah = array_filter(
                $daftarMatakuliah,
                function ($matakuliah) use ($kataKunci) {
                    return str_contains(
                        strtolower($matakuliah['nama']),
                        strtolower($kataKunci)
                    );
                }
            );
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $daftarMatakuliah,
            'kataKunci' => $kataKunci
        ]);
    }

    public function show(string $kode)
    {
        return view('matakuliah.show', [
            'kode' => $kode
        ]);
    }
}