@extends('layouts.app')

@section('konten')
<div class="container">
    <h1>Data Mahasiswa</h1>

    @if (session('success'))
        <div class="alert alert-success">
            {{ session('success') }}
        </div>
    @endif

    <table class="table table-striped">
        <thead>
            <tr>
                <th>NIM</th>
                <th>Nama</th>
                <th>Program Studi</th>
                <th>Angkatan</th>
                <th>IPK</th>
            </tr>
        </thead>

        <tbody>
            @foreach ($daftarMahasiswa as $mahasiswa)
                <tr>
                    <td>{{ $mahasiswa->nim }}</td>
                    <td>{{ $mahasiswa->nama }}</td>
                    <td>{{ $mahasiswa->programStudi->nama }}</td>
                    <td>{{ $mahasiswa->angkatan }}</td>
                    <td>{{ $mahasiswa->ipk }}</td>
                </tr>
            @endforeach
        </tbody>
    </table>

    {{ $daftarMahasiswa->links() }}
</div>
@endsection