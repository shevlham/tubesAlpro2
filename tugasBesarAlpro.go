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
	fmt.Println("===================================")
	fmt.Println("=           Menu Utama            =")
	fmt.Println("===================================")
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

func menuCatatan() {
	var input int

	fmt.Println("=================================")
	fmt.Println("=          Menu Catatan         =")
	fmt.Println("=================================")
	fmt.Println("Pilihan Menu")
	fmt.Println("1. Menambahkan catatan")
	fmt.Println("2. Menghapus catatan")
	fmt.Println("3. Mengubah catatan")
	fmt.Println("4. Daftar catatan")
	fmt.Println("5. Cari Catatan")
	fmt.Println("6. Urutkan Catatan")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih Menu : ")

	fmt.Scan(&input)

	switch input {
	case 1:
		menuTambahCatatan()
	case 2:
		menuHapusCatatan()
	case 3:
		menuUbahCatatan()
	case 4:
		menuDaftarCatatan()
	case 5:
		menuCariCatatan()
	case 6:
		menuUrutkanCatatan()
	case 0:
		menuUtama()
	}
}
func menuSoal() {
	var input int

	fmt.Println("=================================")
	fmt.Println("=          Menu Soal            =")
	fmt.Println("=================================")
	fmt.Println("Pilihan Menu")
	fmt.Println("1. Cari Soal Berdasarkan Materi")
	fmt.Println("2. Cari Soal Berdasarkan Catatan")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih Menu : ")

	fmt.Scan(&input)

	switch input {
	case 1:
		cariSoalMateri()
	case 2:
		cariSoalCatatan()
	case 0:
		menuUtama()
	}
}
func menuJadwal()
	var inputjadwal int

	fmt.Println(" _________________________________")
	fmt.Println("|                                 |")
	fmt.Println("| Aplikasi Asistensi Pembelajaran |")
	fmt.Println("|_________________________________|")
	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("=   Menu Jadwal Pembelajaran      =")
	fmt.Println("===================================")
	fmt.Println("Pilihan Menu")
	fmt.Println("1. Menambahkan Jadwal")
	fmt.Println("2. Menghapus Jadwal")
	fmt.Println("3. Mengubah Jadwal")
	fmt.Println("4. Menampilkan Jadwal")
	fmt.Println("5. Mengurutkan Jadwal")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih Menu : ")

	fmt.Scan(&inputjadwal)
}
// program submenu

func menuTambahCatatan() {
	var input int

	// input id
	fmt.Println("==========================================")
	fmt.Println("=     Menu Catatan: Menambah Catatan     =")
	fmt.Println("= Catatan : ketik '0' jika ingin kembali =")
	fmt.Println("==========================================")
	fmt.Print("ID : ")
	fmt.Scan(&input)
	if input != 0 {
		tambahCatatan(input, icat)
		menuTambahCatatan()
	} else {
		menuCatatan()
	}
}

func menuHapusCatatan() {
	var input, i, x int

	// input id yang ingin dihapus
	fmt.Println("==========================================")
	fmt.Println("=      Menu Catatan: Hapus Catatan       =")
	fmt.Println("= Catatan : ketik '0' jika ingin kembali =")
	fmt.Println("==========================================")
	tampilCatatan()
	fmt.Println("==========================================")
	fmt.Print("Pilih ID Catatan yang ingin dihapus : ")
	fmt.Scan(&input)

	x = seqsearchcat(input)
	if input == 0 {
		menuUtama()
	} else if x != -1 {
		for i = x; i < icat; i++ {
			cat[i] = cat[i+1]
		}
		icat--
		fmt.Println("Catatan telah dihapus")
		menuHapusCatatan()
	} else {
		fmt.Println("ID Tidak ditemukan")
		menuHapusCatatan()
	}
}

func menuUbahCatatan() {
	var input, x int

	fmt.Println("==========================================")
	fmt.Println("=     Menu Catatan: mengubah Catatan     =")
	fmt.Println("= Catatan : ketik '0' jika ingin kembali =")
	fmt.Println("==========================================")
	tampilCatatan()
	fmt.Println("==========================================")
	fmt.Print("Pilih ID Catatan yang ingin diubah : ")
	fmt.Scan(&input)

	x = seqsearchcat(input)
	if input == 0 {
		menuCatatan()
	} else if x != -1 {
		tambahCatatan(input, x)
		menuUbahCatatan()
	} else {
		fmt.Println("Data Tidak ditemukan")
		menuUbahCatatan()
	}

}

func menuDaftarCatatan() {
	tampilCatatan()
	menuCatatan()
}

func menuCariCatatan() {
	var input, x int
	fmt.Println("==========================================")
	fmt.Println("=     Menu Catatan: Mencari Catatan      =")
	fmt.Println("= Catatan : ketik '0' jika ingin kembali =")
	fmt.Println("==========================================")
	fmt.Print("ID Catatan : ")
	fmt.Scan(&input)

	x = seqsearchcat(input)
	if input == 0 {
		menuCatatan()
	} else if x != -1 {
		fmt.Println("Catatan ditemukan :")
		fmt.Printf("%7s  %5s / %5s / %5s  %8s %15s\n", "ID Cat", "tgl", "bln", "thn", "topik", "judul")
		fmt.Printf("%7d  %5d / %5d / %d  %8v %15v\n", cat[x].IDc, cat[x].tanggal.hari, cat[x].tanggal.bulan, cat[x].tanggal.tahun, cat[x].topik, cat[x].judul)
		menuCariCatatan()
	} else {
		fmt.Println("Data Tidak ditemukan")
		menuCariCatatan()
	}

}

func menuUrutkanCatatan() {
	var input int

	fmt.Println("===================================")
	fmt.Println("= Menu Catatan: urutkan Catatan =")
	fmt.Println("===================================")
	fmt.Println("Pilihan Urutan : ")
	fmt.Println("1. Urutkan berdasarkan tanggal")
	fmt.Println("2. Urutkan berdasarkan topik")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih Urutan : ")
	fmt.Scan(&input)

	switch input {
	case 0:
		menuCatatan()
	case 1:
		selectSortCat()
		fmt.Println("Catatan telah diurutkan berdasarkan tanggal")
		tampilCatatan()
		menuUrutkanCatatan()
	case 2:
		insertSortCat()
		fmt.Println("Catatan telah diurutkan berdasarkan Topik")
		tampilCatatan()
		menuUrutkanCatatan()
	}

}

// program lainnya --------------------------------------------

func tambahCatatan(input int, j int) {
	cat[j].IDc = input

	// input materi
	fmt.Println("==================================")
	fmt.Println("=          Menu Catatan          =")
	fmt.Println("==================================")
	fmt.Println("Pilihan Materi")
	fmt.Println("1. Matematika")
	fmt.Println("2. Bahasa Indonesia")
	fmt.Println("3. Bahasa Inggris")
	fmt.Println("4. IPA")
	fmt.Println("5. IPS")
	fmt.Println("0. kembali")
	fmt.Print("Pilih Materi :")
	fmt.Scan(&input)
	switch input {
	case 1:
		cat[j].topik = "Matematika"
	case 2:
		cat[j].topik = "Bahasa Indonesia"
	case 3:
		cat[j].topik = "Bahasa Inggris"
	case 4:
		cat[j].topik = "IPA"
	case 5:
		cat[j].topik = "IPS"
	}

	// input tanggal
	fmt.Println("==================================")
	fmt.Println("=          Menu Catatan          =")
	fmt.Println("==================================")
	fmt.Println("= Catatan : bulan ditulis        =")
	fmt.Println("= menggunakan angka (1-12)       =")
	fmt.Println("==================================")
	fmt.Print("tanggal Catatan : ")
	fmt.Scan(&cat[j].tanggal.hari)
	fmt.Print("bulan Catatan : ")
	fmt.Scan(&cat[j].tanggal.bulan)
	fmt.Print("tahun Catatan : ")
	fmt.Scan(&cat[j].tanggal.tahun)

	// input judul
	fmt.Println("==================================")
	fmt.Println("=          Menu Catatan          =")
	fmt.Println("==================================")
	fmt.Print("Judul Catatan : ")
	fmt.Scan(&cat[j].judul)

	// input isi
	fmt.Println("==================================")
	fmt.Println("=          Menu Catatan           ")
	fmt.Println("==================================")
	fmt.Print("Isi Catatan : ")
	fmt.Scan(&cat[j].isi)

	// iterasi data dan balik ke menu tambah catatan
	icat++

}

func tampilCatatan() {
	var i int
	fmt.Println("____________________________________________________________________________")
	fmt.Println()
	fmt.Println("                               Daftar Catatan")
	fmt.Println("____________________________________________________________________________")
	fmt.Println()
	fmt.Printf("%10s  %5s / %5s / %5s  %18s %22s\n", "ID Cat", "tgl", "bln", "thn", "topik", "judul")
	for i = 0; i < icat; i++ {
		fmt.Printf("%10d  %5d / %5d / %5d  %18s %22s\n", cat[i].IDc, cat[i].tanggal.hari, cat[i].tanggal.bulan, cat[i].tanggal.tahun, cat[i].topik, cat[i].judul)
	}
	fmt.Println("____________________________________________________________________________")
	fmt.Println()
	fmt.Println()

}

func seqsearchcat(x int) int {
	var found, i int
	found = -1
	for i = 0; i < icat; i++ {
		if cat[i].IDc == x {
			found = i
		}
	}
	return found
}

func selectSortCat() {

	// urut tahun
	var i, j, min int
	for i = 0; i < icat; i++ {
		min = i
		for j = i + 1; j < icat; j++ {
			if cat[min].tanggal.tahun > cat[j].tanggal.tahun {
				min = j
			}
		}
		cat[i].tanggal.tahun, cat[min].tanggal.tahun = cat[min].tanggal.tahun, cat[i].tanggal.tahun
	}

	//urut bulan
	for i = 0; i < icat; i++ {
		min = i
		for j = i + 1; j < icat; j++ {
			if cat[min].tanggal.bulan > cat[j].tanggal.tahun && cat[min].tanggal.tahun == cat[j].tanggal.tahun {
				min = j
			}
		}
		cat[i].tanggal.tahun, cat[min].tanggal.tahun = cat[min].tanggal.tahun, cat[i].tanggal.tahun
	}

	// urut tanggal
	for i = 0; i < icat; i++ {
		min = i
		for j = i + 1; j < icat; j++ {
			if cat[min].tanggal.bulan > cat[j].tanggal.tahun && cat[min].tanggal.tahun == cat[j].tanggal.tahun && cat[min].tanggal.bulan == cat[j].tanggal.bulan {
				min = j
			}
		}
		cat[i].tanggal.tahun, cat[min].tanggal.tahun = cat[min].tanggal.tahun, cat[i].tanggal.tahun
	}
}
func insertSortCat() {
	var i int
	var key tCatatan
	for i = 1; i < icat; i++ {
		key = cat[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && cat[j].topik[0] > key.topik[0] {
			cat[j+1] = cat[j]
			j--
		}

		cat[j+1] = key
	}
}

func cariSoalMateri()  {}
func cariSoalCatatan() {}

