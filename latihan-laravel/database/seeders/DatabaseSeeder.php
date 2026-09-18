<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\Mahasiswa;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        $this->call(ProgramStudiSeeder::class);
        $this->call(MatakuliahSeeder::class);

        Mahasiswa::factory()->count(30)->create();
    }
}