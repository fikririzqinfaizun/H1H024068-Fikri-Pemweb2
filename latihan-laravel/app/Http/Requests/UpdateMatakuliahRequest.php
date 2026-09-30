<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Validation\Rule;

class UpdateMatakuliahRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     */
    public function rules(): array
    {
        return [
            'kode' => [
                'required',
                'string',
                'max:10',
                Rule::unique('matakuliahs', 'kode')->ignore($this->route('matakuliah')),
            ],
            'nama' => ['required', 'string', 'max:100'],
            'sks' => ['required', 'integer', 'min:1', 'max:6'],
            'semester' => ['required', 'integer', 'min:1', 'max:8'],
        ];
    }
}
