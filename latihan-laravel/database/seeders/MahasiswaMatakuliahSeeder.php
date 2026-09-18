<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\Mahasiswa;
use App\Models\Matakuliah;

class MahasiswaMatakuliahSeeder extends Seeder
{
    public function run(): void
    {
        $mahasiswa = Mahasiswa::all();
        $matakuliah = Matakuliah::all();

        foreach ($mahasiswa as $mhs) {
            $pilihanMatakuliah = $matakuliah->random(5);

            foreach ($pilihanMatakuliah as $mk) {
                $mhs->matakuliah()->attach($mk->id, [
                    'nilai' => fake()->randomFloat(2, 60, 100),
                ]);
            }
        }
    }
}