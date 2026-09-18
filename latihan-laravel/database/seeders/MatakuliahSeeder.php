<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\Matakuliah;

class MatakuliahSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $matakuliah = [
            [
                'kode' => 'TK101',
                'nama' => 'Pemrograman Dasar',
                'sks' => 3,
                'semester' => 1,
            ],
            [
                'kode' => 'TK102',
                'nama' => 'Sistem Digital',
                'sks' => 3,
                'semester' => 1,
            ],
            [
                'kode' => 'TK201',
                'nama' => 'Struktur Data',
                'sks' => 3,
                'semester' => 2,
            ],
            [
                'kode' => 'TK202',
                'nama' => 'Pemrograman Web',
                'sks' => 3,
                'semester' => 2,
            ],
            [
                'kode' => 'TK301',
                'nama' => 'Jaringan Komputer',
                'sks' => 3,
                'semester' => 3,
            ],
            [
                'kode' => 'TK302',
                'nama' => 'Basis Data',
                'sks' => 3,
                'semester' => 3,
            ],
            [
                'kode' => 'TK401',
                'nama' => 'Internet of Things',
                'sks' => 3,
                'semester' => 4,
            ],
            [
                'kode' => 'TK402',
                'nama' => 'Keamanan Jaringan',
                'sks' => 3,
                'semester' => 4,
            ],
            [
                'kode' => 'TK501',
                'nama' => 'Pemrograman Perangkat Keras',
                'sks' => 3,
                'semester' => 5,
            ],
            [
                'kode' => 'TK502',
                'nama' => 'Sistem Operasi',
                'sks' => 3,
                'semester' => 5,
            ],
        ];

        foreach ($matakuliah as $data) {
            Matakuliah::create($data);
        }
    }
}
