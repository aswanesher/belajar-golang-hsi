package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// input nama
	fmt.Print("Masukkan nama: ")
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)

	// input umur
	fmt.Print("Masukkan umur: ")
	ageInput, _ := reader.ReadString('\n')
	ageStr := strings.TrimSpace(ageInput)

	// konversi umur ke integer
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		logrus.WithError(err).Error("Gagal mengkonversi umur ke integer")
		fmt.Println("Input umur tidak valid. Pastikan Anda memasukkan angka.")
		return
	}

	// validasi umur
	if age < 18 {
		err := errors.New("umur kurang dari 18 tahun")
		logrus.WithError(err).Error("Validasi umur gagal")
		fmt.Println("Maaf, Anda belum cukup umur untuk mengikuti pelatihan ini.")
		return
	}

	// usia valid
	fmt.Printf("Selamat datang, %s! Anda berusia %d tahun.\n", name, age)
}
