package main

import "fmt"

// konstanta array

const NMAX int = 100

// Kumpulan Struct

type tTanggal struct {
	hari  int
	bulan int
	tahun int
}

type tCatatan struct {
	IDc               int
	judul, isi, topik string
	tanggal           tTanggal
}

type tSoal struct {
	IDs                                 int
	topik, pertanyaan, jawaban, tingkat string
}

type tJadwal struct {
	IDj     int
	topik   string
	tanggal tTanggal
}

// variabel global array

var cat [NMAX]tCatatan
var sol [NMAX]tSoal
var jad [NMAX]tJadwal

// iterasi array

var icat int = 0
var isol int = 0
var ijad int = 0

// fungsi utama

func main() {
	menuUtama()
}

// program seluruh menu

func menuUtama() {
	var input int

	fmt.Println(" _________________________________")
	fmt.Println("|                                 |")
	fmt.Println("| Aplikasi Asistensi Pembelajaran |")
	fmt.Println("|_________________________________|")
	fmt.Println()
	fmt.Println("=================================")
	fmt.Println("=           Menu Utama          =")
	fmt.Println("=================================")
	fmt.Println("Pilihan Menu")
	fmt.Println("1. Catatan")
	fmt.Println("2. Soal")
	fmt.Println("3. Jadwal Pembelajaran")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih Menu : ")

	fmt.Scan(&input)

	switch input {
	case 1:
		menuCatatan()
	case 2:
		menuSoal()
	case 3:
		menuJadwal()
	case 0:
	}
}

func menuCatatan() {}
func menuSoal()    {}
func menuJadwal()  {}

// program  lainnya
