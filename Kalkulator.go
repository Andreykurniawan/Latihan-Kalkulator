package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Kalkulator Sederhana")
	fmt.Println("Pilih operasi:")
	fmt.Println("1. Penjumlahan (+)")
	fmt.Println("2. Pengurangan (-)")
	fmt.Println("3. Perkalian (*)")
	fmt.Println("4. Pembagian (/)")
	fmt.Println("5. Akar kuadrat (√)")
	fmt.Println("6. Pangkat (^)")
	fmt.Print("Masukkan nomor operasi: ")
	fmt.Scanln(&input)

	switch input {
	case "1":
		fmt.Println("Masukkan angka pertama:")
		num1 := getInput()
		fmt.Println("Masukkan angka kedua:")
		num2 := getInput()
		result := num1 + num2
		fmt.Printf("Hasil: %.2f\n", result)
	case "2":
		fmt.Println("Masukkan angka pertama:")
		num1 := getInput()
		fmt.Println("Masukkan angka kedua:")
		num2 := getInput()
		result := num1 - num2
		fmt.Printf("Hasil: %.2f\n", result)
	case "3":
		fmt.Println("Masukkan angka pertama:")
		num1 := getInput()
		fmt.Println("Masukkan angka kedua:")
		num2 := getInput()
		result := num1 * num2
		fmt.Printf("Hasil: %.2f\n", result)
	case "4":
		fmt.Println("Masukkan angka pertama:")
		num1 := getInput()
		fmt.Println("Masukkan angka kedua:")
		num2 := getInput()
		if num2 == 0 {
			fmt.Println("Pembagian dengan nol tidak diizinkan.")
		} else {
			result := num1 / num2
			fmt.Printf("Hasil: %.2f\n", result)
		}
	case "5":
		fmt.Println("Masukkan angka:")
		num := getInput()
		if num < 0 {
			fmt.Println("Akar kuadrat dari angka negatif tidak diizinkan.")
		} else {
			result := math.Sqrt(num)
			fmt.Printf("Hasil: %.2f\n", result)
		}
	case "6":
		fmt.Println("Masukkan angka:")
		num1 := getInput()
		fmt.Println("Masukkan pangkat:")
		num2 := getInput()
		result := math.Pow(num1, num2)
		fmt.Printf("Hasil: %.2f\n", result)
	default:
		fmt.Println("Operasi tidak valid.")
	}
}

func getInput() float64 {
	var input string
	fmt.Scanln(&input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input tidak valid. Masukkan angka.")
		return getInput()
	}
	return num
}
