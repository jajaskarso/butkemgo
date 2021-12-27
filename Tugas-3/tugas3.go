package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	panjangPersegiPanjangBaru, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarPersegiPanjangBaru, _ := strconv.Atoi(lebarPersegiPanjang)

	var kelilingPersegiPanjang int = panjangPersegiPanjangBaru * lebarPersegiPanjangBaru

	fmt.Println(kelilingPersegiPanjang)

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	alasSegitigaBaru, _ := strconv.Atoi(alasSegitiga)
	tinggiSegitigaBaru, _ := strconv.Atoi(tinggiSegitiga)

	var luasSegitiga int = alasSegitigaBaru * tinggiSegitigaBaru / 2

	fmt.Println(luasSegitiga)

	//soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Indeksnya A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Indeksnya B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Indeksnya C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Indeksnya D")
	} else {
		fmt.Println("Indeksnya E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Indeksnya A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Indeksnya B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Indeksnya C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Indeksnya D")
	} else {
		fmt.Println("Indeksnya E")
	}

	//soal 3
	var tanggal = 12
	var bulan = 7
	var tahun = 2002

	bulanBaru1 := strconv.FormatInt(int64(bulan), 10)

	switch bulan {
	case 1:
		bulanBaru1 = "Januari"
	case 2:
		bulanBaru1 = "Februari"
	case 3:
		bulanBaru1 = "Maret"
	case 4:
		bulanBaru1 = "April"
	case 5:
		bulanBaru1 = "Mei"
	case 6:
		bulanBaru1 = "Juni"
	case 7:
		bulanBaru1 = "Juli"
	case 8:
		bulanBaru1 = "Agustus"
	case 9:
		bulanBaru1 = "September"
	case 10:
		bulanBaru1 = "Oktober"
	case 11:
		bulanBaru1 = "November"
	case 12:
		bulanBaru1 = "Desember"
	}

	tanggalBaru := strconv.FormatInt(int64(tanggal), 10)
	tahunBaru := strconv.FormatInt(int64(tahun), 10)

	fmt.Println(tanggalBaru + " " + bulanBaru1 + " " + tahunBaru)

	//soal 4

	if tahun > 1944 && tahun <= 1964 {
		fmt.Println("Anda Baby Boomers")
	} else if tahun > 1964 && tahun <= 1979 {
		fmt.Println("Anda Gen X")
	} else if tahun > 1979 && tahun <= 1994 {
		fmt.Println("Anda Gen Y")
	} else if tahun > 1994 && tahun <= 2015 {
		fmt.Println("Anda Gen Z")
	}
}
