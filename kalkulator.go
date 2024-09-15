package main

import "fmt"

func main () {
	var angka_1 float64
	var operator string
	var angka_2 float64
	
	fmt.Print("Masukkan operator (+,- ,* ,/):  ")
	fmt.Scanf("%s\n", &operator)

	fmt.Print("Masukkan angka_1: ")
	fmt.Scanf("%f\n", &angka_1)

	
	fmt.Print(("Masukkan angka_2: "))
	fmt.Scanf("%f\n", &angka_2)
	

	switch operator {
	case "+":
		fmt.Printf("Hasil: %.2f\n", angka_1+angka_2)
	case "-":
		fmt.Printf("Hasil: %.2f\n", angka_1-angka_2)
	case "*":
		fmt.Printf("Hasil: %.2f\n", angka_1*angka_2)
	case "/":
		if angka_2 != 0 {
			fmt.Printf("Hasil: %.2f\n", angka_1/angka_2)
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
		}
	default:
		fmt.Println("Kalkulator tidak mengerti input user.")
}
}