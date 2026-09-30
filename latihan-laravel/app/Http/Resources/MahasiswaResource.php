<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class MahasiswaResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        $fields = $request->query('fields');

        $allowedFields = [
            'id',
            'nim',
            'nama',
            'email',
            'angkatan',
            'ipk',
            'aktif',
            'program_studi',
            'dibuat_pada',
        ];

        if ($fields) {
            $requestedFields = array_map(
                'trim',
                explode(',', $fields)
            );

            $requestedFields = array_values(
                array_intersect($requestedFields, $allowedFields)
            );
        } else {
            $requestedFields = $allowedFields;
        }

        $data = [];

        foreach ($requestedFields as $field) {
            switch ($field) {
                case 'id':
                    $data['id'] = $this->id;
                    break;

                case 'nim':
                    $data['nim'] = $this->nim;
                    break;

                case 'nama':
                    $data['nama'] = $this->nama;
                    break;

                case 'email':
                    $data['email'] = $this->email;
                    break;

                case 'angkatan':
                    $data['angkatan'] = $this->angkatan;
                    break;

                case 'ipk':
                    $data['ipk'] = (float) $this->ipk;
                    break;

                case 'aktif':
                    $data['aktif'] = $this->aktif;
                    break;

                case 'program_studi':
                    if ($this->relationLoaded('programStudi')) {
                        $data['program_studi'] = [
                            'id' => $this->programStudi->id,
                            'kode' => $this->programStudi->kode,
                            'nama' => $this->programStudi->nama,
                        ];
                    }
                    break;

                case 'dibuat_pada':
                    $data['dibuat_pada'] = $this->created_at?->toIso8601String();
                    break;
            }
        }

        return $data;
    }
}
