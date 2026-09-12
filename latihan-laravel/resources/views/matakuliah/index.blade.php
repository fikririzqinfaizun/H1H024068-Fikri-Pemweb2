@extends('layouts.app')

@section('judul', 'Daftar Mata Kuliah')

@section('konten')

<h1 class="h3 mb-4">Daftar Mata Kuliah</h1>

<x-kartu-info judul="Informasi">
    Berikut adalah daftar mata kuliah pada halaman ini.
</x-kartu-info>

{{-- Form pencarian --}}
<form method="GET" action="{{ route('matakuliah.index') }}" class="mb-4">
    <div class="input-group">
        <input
            type="text"
            name="q"
            value="{{ $kataKunci }}"
            class="form-control"
            placeholder="Cari mata kuliah..."
        >

        <button type="submit" class="btn btn-primary">
            Cari
        </button>
    </div>
</form>

{{-- Menampilkan kata kunci pencarian --}}
@if ($kataKunci !== '')
    <div class="alert alert-info">
        Hasil pencarian untuk:
        <strong>{{ $kataKunci }}</strong>
    </div>
@endif

{{-- Tabel mata kuliah --}}
<table class="table table-bordered bg-white">
    <thead>
        <tr>
            <th>Kode</th>
            <th>Nama Mata Kuliah</th>
            <th>SKS</th>
            <th>Aksi</th>
        </tr>
    </thead>

    <tbody>
        @forelse ($daftarMatakuliah as $matakuliah)
            <tr>
                <td>{{ $matakuliah['kode'] }}</td>

                <td>{{ $matakuliah['nama'] }}</td>

                <td>
                    <x-badge-sks :sks="$matakuliah['sks']" />
                </td>

                <td>
                    <a
                        href="{{ route('matakuliah.show', $matakuliah['kode']) }}"
                        class="btn btn-sm btn-primary"
                    >
                        Detail
                    </a>
                </td>
            </tr>
        @empty
            <tr>
                <td colspan="4" class="text-center">
                    Data mata kuliah tidak ditemukan.
                </td>
            </tr>
        @endforelse
    </tbody>
</table>

@endsection