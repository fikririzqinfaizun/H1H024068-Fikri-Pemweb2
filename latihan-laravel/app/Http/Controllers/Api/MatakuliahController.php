<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Requests\StoreMatakuliahRequest;
use App\Http\Requests\UpdateMatakuliahRequest;
use App\Http\Resources\MatakuliahResource;
use App\Models\Matakuliah;

class MatakuliahController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $matakuliah = Matakuliah::query()
            ->orderBy('semester')
            ->orderBy('kode')
            ->paginate(10);

        return MatakuliahResource::collection($matakuliah);
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(StoreMatakuliahRequest $request)
    {
        $matakuliah = Matakuliah::create($request->validated());

        return (new MatakuliahResource($matakuliah))
            ->response()
            ->setStatusCode(201);
    }

    /**
     * Display the specified resource.
     */
    public function show(string $id)
    {
        $matakuliah = Matakuliah::find($id);

        if (!$matakuliah) {
            return response()->json([
                'sukses' => false,
                'pesan' => 'Sumber daya tidak ditemukan',
            ], 404);
        }

        return new MatakuliahResource($matakuliah);
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateMatakuliahRequest $request, string $id)
    {
        $matakuliah = Matakuliah::find($id);

        if (!$matakuliah) {
            return response()->json([
                'sukses' => false,
                'pesan' => 'Sumber daya tidak ditemukan',
            ], 404);
        }

        $matakuliah->update($request->validated());

        return new MatakuliahResource($matakuliah);
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        $matakuliah = Matakuliah::find($id);

        if (!$matakuliah) {
            return response()->json([
                'sukses' => false,
                'pesan' => 'Sumber daya tidak ditemukan',
            ], 404);
        }

        $matakuliah->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil dihapus',
        ]);
    }
}
