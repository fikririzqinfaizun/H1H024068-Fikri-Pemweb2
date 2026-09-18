@extends('layouts.app')

@section('judul', 'Mata Kuliah Mahasiswa')

@section('konten')

<div class="container">

    <div class="d-flex justify-content-between align-items-center mb-4">
        <div>
            <h1 class="h3 mb-1">Mata Kuliah Mahasiswa</h1>
            <p class="text-muted mb-0">
                Daftar mata kuliah dan nilai mahasiswa
            </p>
        </div>

        <a href="{{ route('mahasiswa.detail', $mahasiswa->id) }}"
           class="btn btn-secondary">
            Kembali
        </a>
    </div>

    <!-- Informasi Mahasiswa -->
    <div class="card mb-4">
        <div class="card-header">
            <strong>Informasi Mahasiswa</strong>
        </div>

        <div class="card-body">
            <div class="row">
                <div class="col-md-6">
                    <p class="mb-2">
                        <strong>NIM:</strong>
                        {{ $mahasiswa->nim }}
                    </p>

                    <p class="mb-2">
                        <strong>Nama:</strong>
                        {{ $mahasiswa->nama }}
                    </p>
                </div>

                <div class="col-md-6">
                    <p class="mb-2">
                        <strong>Program Studi:</strong>
                        {{ $mahasiswa->programStudi->nama }}
                    </p>

                    <p class="mb-2">
                        <strong>IPK:</strong>
                        {{ $mahasiswa->ipk }}
                    </p>
                </div>
            </div>
        </div>
    </div>

    <!-- Daftar Mata Kuliah -->
    <div class="card">
        <div class="card-header">
            <strong>Daftar Mata Kuliah</strong>
        </div>

        <div class="card-body">

            <div class="table-responsive">
                <table class="table table-striped table-bordered mb-0">

                    <thead>
                        <tr>
                            <th>No</th>
                            <th>Kode</th>
                            <th>Mata Kuliah</th>
                            <th>SKS</th>
                            <th>Semester</th>
                            <th>Nilai</th>
                        </tr>
                    </thead>

                    <tbody>

                        @forelse ($mahasiswa->matakuliah as $index => $mk)

                            <tr>
                                <td>{{ $index + 1 }}</td>
                                <td>{{ $mk->kode }}</td>
                                <td>{{ $mk->nama }}</td>
                                <td>{{ $mk->sks }}</td>
                                <td>{{ $mk->semester }}</td>
                                <td>{{ $mk->pivot->nilai }}</td>
                            </tr>

                        @empty

                            <tr>
                                <td colspan="6" class="text-center">
                                    Belum ada mata kuliah.
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