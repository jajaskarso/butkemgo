package main

import (
	"fmt"
	"strings"
)

// fungsi buat soal 1

func luasPersegiPanjang(sisi1 int, sisi2 int) int {
	return sisi1 * sisi2
}

func kelilingPersegiPanjang(sisi1 int, sisi2 int) int {
	return (sisi1 + sisi2) * 2
}

func volumeBalok(sisi1 int, sisi2 int, sisi3 int) int {
	return sisi1 * sisi2 * sisi3
}

// fungsi buat soal 2

func introduce(nama, jenisKelamin, profesi, umur string) string {
	if jenisKelamin == "laki-laki" {
		kalimat := "Pak " + nama + " adalah seorang " + profesi + " berumur " + umur + " tahun"
		return kalimat
	} else if jenisKelamin == "perempuan" {
		kalimat := "Bu " + nama + " adalah seorang " + profesi + " berumur " + umur + " tahun"
		return kalimat
	} else {
		return "Mohon maaf. Jenis kelamin yang Anda masukkan tidak tersedia"
	}
}

// fungsi buat soal 3

func buahFavorit(nama string, buahKesukaan ...string) string {
	kalimat := "Halo, nama saya " + nama + ". " + "Buah kesukaan saya adalah " + strings.Join(buahKesukaan, ", ")
	return kalimat
}

func main() {

	//soal 1

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("Luas persegi panjangnya adalah:", luas)
	fmt.Println("Keliling persegi panjangnya adalah:", keliling)
	fmt.Println("Volume persegi panjangnya adalah:", volume)

	//soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	anas := introduce("Anas", "laki-laki", "pelajar", "20")
	fmt.Println(anas)

	//soal 3

	buah := []string{"semangka", "jeruk", "melon", "pepaya"}

	buahFavoritJohn := buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn) // halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//soal 4

	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	tambahDataFilm := func(title, durasi, genre, tahun string) map[string]string {
		satuDataFilm := map[string]string{
			"title":  title,
			"durasi": durasi,
			"genre":  genre,
			"tahun":  tahun,
		}
		dataFilm = append(dataFilm, satuDataFilm)
		return satuDataFilm
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
