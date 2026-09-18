@extends('layouts.app')

@section('judul', 'Detail Mahasiswa')

@section('konten')
<div class="container">

    <div class="d-flex justify-content-between align-items-center mb-4">
        <h1 class="h3">Detail Mahasiswa</h1>

        <a href="{{ route('mahasiswa.data') }}" class="btn btn-secondary">
            Kembali
        </a>
    </div>

    <div class="card mb-4">
        <div class="card-header">
            <strong>Informasi Mahasiswa</strong>
        </div>

        <div class="card-body">
            <table class="table table-borderless mb-0">
                <tr>
                    <th width="200">NIM</th>
                    <td>{{ $mahasiswa->nim }}</td>
                </tr>

                <tr>
                    <th>Nama</th>
                    <td>{{ $mahasiswa->nama }}</td>
                </tr>

                <tr>
                    <th>Program Studi</th>
                    <td>{{ $mahasiswa->programStudi->nama }}</td>
                </tr>

                <tr>
                    <th>Angkatan</th>
                    <td>{{ $mahasiswa->angkatan }}</td>
                </tr>

                <tr>
                    <th>IPK</th>
                    <td>{{ $mahasiswa->ipk }}</td>
                </tr>

                <tr>
                    <th>Email</th>
                    <td>{{ $mahasiswa->email }}</td>
                </tr>

                <tr>
                    <th>Status</th>
                    <td>
                        @if ($mahasiswa->aktif)
                            <span class="badge bg-success">Aktif</span>
                        @else
                            <span class="badge bg-secondary">Tidak Aktif</span>
                        @endif
                    </td>
                </tr>
            </table>
        </div>
    </div>

    <div class="card">
        <div class="card-header">
            <strong>Daftar Mata Kuliah</strong>
        </div>

        <div class="card-body">
            <div class="table-responsive">
                <table class="table table-striped table-bordered mb-0">
                    <thead>
                        <tr>
                            <th>No.</th>
                            <th>Kode</th>
                            <th>Mata Kuliah</th>
                            <th>SKS</th>
                            <th>Semester</th>
                            <th>Nilai</th>
                        </tr>
                    </thead>

                    <tbody>
                        @forelse ($mahasiswa->matakuliah as $index => $matakuliah)
                            <tr>
                                <td>{{ $index + 1 }}</td>
                                <td>{{ $matakuliah->kode }}</td>
                                <td>{{ $matakuliah->nama }}</td>
                                <td>{{ $matakuliah->sks }}</td>
                                <td>{{ $matakuliah->semester }}</td>
                                <td>{{ $matakuliah->pivot->nilai }}</td>
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