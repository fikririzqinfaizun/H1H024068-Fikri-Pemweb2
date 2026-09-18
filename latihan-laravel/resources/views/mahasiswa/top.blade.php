@extends('layouts.app')

@section('judul', '10 Mahasiswa dengan IPK Tertinggi')

@section('konten')
<div class="container">

    <div class="d-flex justify-content-between align-items-center mb-4">
        <h1 class="h3">10 Mahasiswa dengan IPK Tertinggi</h1>

        <a href="{{ route('mahasiswa.data') }}" class="btn btn-secondary">
            Kembali
        </a>
    </div>

    <div class="card">
        <div class="card-header">
            <strong>Top 10 Mahasiswa Teknik Komputer</strong>
        </div>

        <div class="card-body">
            <div class="table-responsive">
                <table class="table table-striped table-bordered mb-0">
                    <thead>
                        <tr>
                            <th>Ranking</th>
                            <th>NIM</th>
                            <th>Nama</th>
                            <th>Program Studi</th>
                            <th>IPK</th>
                            <th>Aksi</th>
                        </tr>
                    </thead>

                    <tbody>
                        @forelse ($mahasiswa as $index => $mhs)
                            <tr>
                                <td>{{ $index + 1 }}</td>
                                <td>{{ $mhs->nim }}</td>
                                <td>{{ $mhs->nama }}</td>
                                <td>{{ $mhs->programStudi->nama }}</td>
                                <td>{{ $mhs->ipk }}</td>
                                <td>
                                    <a
                                        href="{{ route('mahasiswa.detail', $mhs->id) }}"
                                        class="btn btn-sm btn-primary"
                                    >
                                        Detail
                                    </a>
                                </td>
                            </tr>
                        @empty
                            <tr>
                                <td colspan="6" class="text-center">
                                    Data mahasiswa tidak ditemukan.
                                </td>
                            </tr>
                        @endforelse
                    </tbody>
                </table>
            </div>
        </div>
    </div>

</div>
@endsection