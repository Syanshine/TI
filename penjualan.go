package main

import "fmt"


func main () {
	var nama_customer, nama_barang string
	var harga float32 = 25000
	var hasil_discount, total_harga float32
	var kuantitas_barang int
	
	fmt.Print("nama_customer: ")
	fmt.Scanf("%s\n", &nama_customer)
	
	fmt.Print("nama_barang: ")
	fmt.Scanf("%s\n", &nama_barang)
	
	fmt.Print("kuantitas_barang: ")
	fmt.Scanf("%d\n", &kuantitas_barang)

	switch {
	case kuantitas_barang > 5:
		hasil_discount = 0.1
	default:
		hasil_discount = 0.02	
	}

	sub_total := float32(kuantitas_barang) * harga
	total_harga = sub_total - (hasil_discount * sub_total)

fmt.Println("hasil diskon: ", hasil_discount)
fmt.Println("quantity: ", kuantitas_barang)
fmt.Println("harga: ", harga)
fmt.Println("Total harga adalah: ", total_harga)
}
	
