package main

import "fmt"

func main() {
	var minimarketStatus = "open"
	var telur = "soldout"
	var buah = "aman"
	if minimarketStatus == "open" {
		fmt.Println("saya akan membeli telur dan buah")
		if telur == "soldout" || buah == "soldout" {
			fmt.Println("belanjaan saya tidak lengkap")
		} else if telur == "soldout" {
			fmt.Println("telur habis")
		} else if buah == "soldout" {
			fmt.Println("buah habis")
		}
	} else {
		fmt.Println("minimarket tutup, saya pulang lagi")
	}
}
