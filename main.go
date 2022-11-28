package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
);

type Customer struct {
	name string
	subtotal []int
}

type Products struct {
	id int
	name string
	price int
}

func main() {
	// list products
	products := []Products{
		{ id: 1, name: "Kacang", price: 1000 },
		{ id: 2, name: "Buku", price: 20000 },
		{ id: 3, name: "Wafer", price: 2000 },
		{ id: 4, name: "Permen", price: 500 },
	}

	// customer
	var customer Customer;

	// initiate input
	input := bufio.NewScanner(os.Stdin);
	fmt.Print("Masukkan nama anda : ")
	input.Scan();
	customer.name = input.Text();
	fmt.Println();

	// temp subtotal
	var subtotal = 0;

	fmt.Println("Welcome", customer.name, "to BukaToko :)");	

	for {
		for _, product := range products {
			fmt.Println(product.id, "-", product.name, "- Rp.", product.price);
		}
		fmt.Print("Silahkan pilih produk yang ingin anda beli : ");
		input.Scan();
		choose := input.Text();
		fmt.Println();
		id, _ := strconv.Atoi(choose);

		fmt.Print("Anda akan membeli ", products[id - 1].name, " dengan harga ", products[id - 1].price, ", masukkan jumlah : ");
		input.Scan();
		qty := input.Text();
		fmt.Println();
		qtyc, _ := strconv.Atoi(qty);
		subtotal += append(customer.subtotal, products[id - 1].price * qtyc)[0];
		fmt.Println("Pembelian sukses, subtotal :", subtotal)

		fmt.Print("Mau beli lagi ? [Y/N] : ");
		input.Scan();
		answer := input.Text();
		fmt.Println();
		if answer == "Y" {
			continue;
		} else if answer == "N" {
			fmt.Print("Masukkan uang anda : ");
			input.Scan();
			money := input.Text();
			fmt.Println();
			mon, _ := strconv.Atoi(money);

			if mon >= subtotal {
				fmt.Println("Pembayaran sukses, uang kembalian Rp.", mon - subtotal, ":)")
			} else {
				fmt.Println("Uang anda tidak mencukupi, silahkan ambil uang dulu :(");
			}
			break;
		} else {
			fmt.Println("Anda salah memasukkan perintah !!!");
			break;
		}
	}
}