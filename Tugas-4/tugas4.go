package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			if i%3 == 0 {
				fmt.Println(i, " - I Love Coding")
			} else {
				fmt.Println(i, " - JCC")
			}
		} else if i%2 == 0 {
			fmt.Println(i, " - Candradimuka")
		}
	}

	//soal 2
	for i := 1; i <= 7; i++ {
		fmt.Println(strings.Repeat("#", i))
	}

	//soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	//soal 4
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}

	for i, v := range sayuran {
		nomorUrut := strconv.FormatInt(int64(i)+1, 10)
		fmt.Println(nomorUrut+".", v)
	}

	// for i := 0; i < len(sayuran); i++ {
	// 	fmt.Println(i+1, sayuran[i])
	// }

	//soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	satuan["volume balok"] = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	for key, value := range satuan {
		fmt.Println(key, "=", value)
	}
}
