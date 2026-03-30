package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Validasi method
	if r.Method != http.MethodPost {
		http.Error(w, "Gunakan method POST", http.StatusMethodNotAllowed)
		return
	}

	// Ambil file dari request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Gagal mengambil file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Buat folder uploads jika belum ada
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		http.Error(w, "Gagal membuat folder", http.StatusInternalServerError)
		return
	}

	// Hindari path traversal (keamanan)
	filename := filepath.Base(handler.Filename)

	// Buat file tujuan
	dst, err := os.Create(filepath.Join("uploads", filename))
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy isi file
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File berhasil diupload: %s\n", filename)
}
