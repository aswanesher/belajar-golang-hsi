package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

// Tipe error untuk validasi input
var (
	ErrInvalidEmail = errors.New("email tidak valid")
	ErrInvalidAge   = errors.New("umur kurang dari 18 tahun")
	ErrEmptyEmail   = errors.New("email tidak boleh kosong")
	ErrEmptyAge     = errors.New("umur tidak boleh kosong")
)

func main() {
	// setup logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	http.HandleFunc("/validate", validateHandler)

	logrus.Info("Bersiap menerima request dari port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.WithError(err).Fatal("Gagal menjalankan server")
		return
	}
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"query":  r.URL.RawQuery,
	}).Info("Request diterima")

	// Deklarasi variable email dan age
	email := strings.TrimSpace(r.URL.Query().Get("email"))
	ageStr := strings.TrimSpace(r.URL.Query().Get("age"))

	// Validasi email
	if email == "" {
		logResponseAndError(w, fmt.Errorf("Input email tidak boleh kosong : %w", ErrEmptyEmail))
		return
	}

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		logResponseAndError(w, fmt.Errorf("Input email tidak valid : %w", ErrInvalidEmail))
		return
	}

	if ageStr == "" {
		logResponseAndError(w, fmt.Errorf("Input umur tidak boleh kosong : %w", ErrEmptyAge))
		return
	}

	// Konversi umur ke integer
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		logrus.WithError(err).Error("Gagal mengkonversi umur ke integer")
		logResponseAndError(w, fmt.Errorf("Input umur tidak valid : %w", err))
		return
	}

	// Validas umur
	if age < 18 {
		logResponseAndError(w, fmt.Errorf("Input umur kurang dari 18 tahun : %w", ErrInvalidAge))
		return
	}

	// Jika semua validasi berhasil
	response := map[string]string{"status": "ok"}
	responseJson(w, http.StatusOK, response)
}

func logResponseAndError(w http.ResponseWriter, err error) {
	logrus.WithError(err).Error("Terjadi kesalahan dalam validasi input")

	var statusCode = http.StatusBadRequest
	var message string

	switch {
	case errors.Is(err, ErrInvalidEmail):
		message = "Email tidak valid"
	case errors.Is(err, ErrInvalidAge):
		message = "Umur kurang dari 18 tahun"
	case errors.Is(err, ErrEmptyEmail):
		message = "Email tidak boleh kosong"
	case errors.Is(err, ErrEmptyAge):
		message = "Umur tidak boleh kosong"
	default:
		message = "Email kosong atau umur kurang dari 18 tahun"
	}

	responseJson(w, statusCode, map[string]string{"error": message})
}

func responseJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
