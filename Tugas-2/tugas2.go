package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//soal 1
	var (
		kata1 = "Jabar "
		kata2 = "Coding "
		kata3 = "Camp "
		kata4 = "Golang "
		kata5 = "2021"
	)

	fmt.Println(kata1 + kata2 + kata3 + kata4 + kata5)

	//soal 2
	halo := "Halo Dunia"
	haloBaru := strings.Replace(halo, "Dunia", "Golang", -1)
	fmt.Println(haloBaru)

	//soal 3
	var (
		kataPertama = "saya "
		kataKedua   = "senang "
		kataKetiga  = "belajar "
		kataKeempat = "golang "
	)
	kataKetigaBaru := strings.Replace(kataKetiga, "r", "R", 1)
	fmt.Println(kataPertama + strings.Title(kataKedua) + kataKetigaBaru + strings.ToUpper(kataKeempat))

	//soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var numangkaPertama, err = strconv.Atoi(angkaPertama)
	var numangkaKedua, err2 = strconv.Atoi(angkaKedua)
	var numangkaKetiga, err3 = strconv.Atoi(angkaKetiga)
	var numangkaKeempat, err4 = strconv.Atoi(angkaKeempat)

	if err == nil && err2 == nil && err3 == nil && err4 == nil {
		fmt.Println(numangkaPertama + numangkaKedua + numangkaKetiga + numangkaKeempat)
	}

	//soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimatBaru := strings.ReplaceAll(kalimat, "halo", "Hi")
	angkaBaru := strconv.Itoa(angka)
	fmt.Println(kalimatBaru + " " + angkaBaru)
}
