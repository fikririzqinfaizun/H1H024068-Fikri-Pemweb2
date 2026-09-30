<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\MahasiswaResource;
use App\Models\ProgramStudi;

class ProgramStudiController extends Controller
{
    /**
     * Menampilkan daftar mahasiswa berdasarkan program studi.
     */
    public function mahasiswa(string $id)
    {
        $programStudi = ProgramStudi::find($id);

        if (!$programStudi) {
            return response()->json([
                'sukses' => false,
                'pesan' => 'Program studi tidak ditemukan',
            ], 404);
        }

        $mahasiswa = $programStudi->mahasiswa()
            ->paginate(10);

        return MahasiswaResource::collection($mahasiswa);
    }
}
