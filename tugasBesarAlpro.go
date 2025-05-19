package main

import (
	"fmt"
)

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

type tJadwal struct {
	IDj     int
	topik   string
	tanggal tTanggal
}
type tSoal struct {
	IDs                                 int
	topik, pertanyaan, jawaban, tingkat string
	pilihan                             [4]string
}

// BANK SOAL
var Soal = [NMAX]tSoal{
	// Matematika
	{1, "Matematika", "Berapakah 12 + 8?", "20", "Mudah", [4]string{"18", "20", "22", "24"}},
	{2, "Matematika", "Berapakah 15 x 6?", "90", "Sedang", [4]string{"80", "85", "90", "95"}},
	{3, "Matematika", "Hasil dari 8^2 - 4^2 adalah?", "48", "Sedang", [4]string{"32", "48", "64", "52"}},
	{4, "Matematika", "Jika x + 5 = 12, maka x adalah?", "7", "Mudah", [4]string{"5", "6", "7", "8"}},
	{5, "Matematika", "Integral dari 2x dx adalah?", "x^2 + C", "Sulit", [4]string{"2x + C", "x^2 + C", "x^2/2 + C", "2x^2 + C"}},

	// Bahasa Indonesia
	{6, "Bahasa Indonesia", "Apa lawan kata dari 'besar'?", "Kecil", "Mudah", [4]string{"Kecil", "Luas", "Panjang", "Banyak"}},
	{7, "Bahasa Indonesia", "Makna konotatif dari 'ular' adalah?", "Pengkhianat", "Sedang", [4]string{"Hewan", "Licik", "Pengkhianat", "Panjang"}},
	{8, "Bahasa Indonesia", "Kalimat 'Dia pergi ke pasar.' adalah?", "Kalimat berita", "Mudah", [4]string{"Kalimat tanya", "Kalimat perintah", "Kalimat berita", "Kalimat aktif"}},
	{9, "Bahasa Indonesia", "Apa fungsi majas metafora?", "Memberi kesan imajinatif", "Sedang", [4]string{"Menjelaskan definisi", "Memberi kesan imajinatif", "Menarik perhatian", "Menghindari pengulangan"}},
	{10, "Bahasa Indonesia", "Struktur teks eksposisi terdiri dari?", "Tesis, argumentasi, penegasan ulang", "Sulit", [4]string{"Orientasi, komplikasi, resolusi", "Tesis, argumentasi, penegasan ulang", "Pembukaan, isi, penutup", "Gagasan utama, isi, simpulan"}},

	// Bahasa Inggris
	{11, "Bahasa Inggris", "What is the opposite of 'hot'?", "Cold", "Mudah", [4]string{"Cold", "Warm", "Cool", "Boiling"}},
	{12, "Bahasa Inggris", "Past tense of 'run'?", "Ran", "Sedang", [4]string{"Run", "Running", "Ran", "Runned"}},
	{13, "Bahasa Inggris", "Translate: 'Mereka sedang belajar.'", "They are studying", "Mudah", [4]string{"They studying", "They are study", "They are studying", "They studies"}},
	{14, "Bahasa Inggris", "Synonym of 'quick'?", "Fast", "Mudah", [4]string{"Slow", "Fast", "Late", "Quiet"}},
	{15, "Bahasa Inggris", "Structure of Present Perfect Continuous?", "Has/have been + V-ing", "Sulit", [4]string{"Was/were + V-ing", "Has/have + V-ed", "Has/have been + V-ing", "Had been + V-ing"}},

	// IPA
	{16, "IPA", "Perubahan dari cair ke gas disebut?", "Menguap", "Mudah", [4]string{"Membeku", "Mengembun", "Menguap", "Mencair"}},
	{17, "IPA", "Fungsi klorofil pada tumbuhan?", "Fotosintesis", "Mudah", [4]string{"Transpirasi", "Respirasi", "Fotosintesis", "Reproduksi"}},
	{18, "IPA", "Contoh gaya gesek adalah?", "Gaya saat mengerem sepeda", "Sedang", [4]string{"Tarikan tali", "Dorongan angin", "Gaya saat mengerem sepeda", "Gaya magnet"}},
	{19, "IPA", "Hukum Newton ke-2 berbunyi?", "F = m x a", "Sedang", [4]string{"F = m + a", "F = m / a", "F = m x a", "F = m - a"}},
	{20, "IPA", "Fungsi mitokondria dalam sel?", "Tempat respirasi sel", "Sulit", [4]string{"Mengatur air", "Pengatur suhu", "Tempat respirasi sel", "Sintesis protein"}},

	// IPS
	{21, "IPS", "Ibu kota Indonesia adalah?", "Jakarta", "Mudah", [4]string{"Bandung", "Jakarta", "Surabaya", "Medan"}},
	{22, "IPS", "Singkatan ASEAN adalah?", "Association of Southeast Asian Nations", "Sedang", [4]string{"Asia Economic Association", "Association of Southeast Asian Nations", "Asia-Europe Society", "Association of South Africa Nations"}},
	{23, "IPS", "Sistem pemerintahan Indonesia adalah?", "Presidensial", "Sedang", [4]string{"Parlementer", "Federasi", "Presidensial", "Monarki"}},
	{24, "IPS", "Pengaruh letak geografis Indonesia?", "Iklim tropis", "Sulit", [4]string{"Iklim dingin", "Iklim subtropis", "Iklim tropis", "Iklim kutub"}},
	{25, "IPS", "Benua tempat Indonesia berada?", "Asia", "Mudah", [4]string{"Eropa", "Afrika", "Asia", "Australia"}},
}

// variabel global array

var cat [NMAX]tCatatan
var jad [NMAX]tJadwal

// iterasi array

var icat int = 0
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
		fmt.Println("Terima Kasih Telah Mengunakan Aplikasi Asistensi Pembelajaran :)")
	default:
		menuUtama()
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
	default:
		menuCatatan()
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
		menuCariSoalMateri()
	case 2:
		MenuCariSoalCatatan()
	case 0:
		menuUtama()
	default:
		menuSoal()
	}
}
func menuJadwal() {}

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
		fmt.Println("_______________________________________________________________________________")
		fmt.Println()
		fmt.Printf("%10s  %5s / %5s / %5s  %18s %22s\n", "ID Cat", "tgl", "bln", "thn", "topik", "judul")
		fmt.Println("_______________________________________________________________________________")
		fmt.Println()
		fmt.Printf("%10d  %5d / %5d / %5d  %18s %22s\n", cat[x].IDc, cat[x].tanggal.hari, cat[x].tanggal.bulan, cat[x].tanggal.tahun, cat[x].topik, cat[x].judul)
		fmt.Println("_______________________________________________________________________________")
		fmt.Println()
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

func menuCariSoalMateri() {
	var input int
	fmt.Println("===============================================")
	fmt.Println("=  Menu Soal : Cari Soal Berdasarkan Materi  =")
	fmt.Println("==============================================")
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
		genSoal("Matematika")
		menuCariSoalMateri()
	case 2:
		genSoal("Bahasa Indonesia")
		menuCariSoalMateri()
	case 3:
		genSoal("Bahasa Inggris")
		menuCariSoalMateri()
	case 4:
		genSoal("IPA")
		menuCariSoalMateri()
	case 5:
		genSoal("IPS")
		menuCariSoalMateri()
	case 0:
		menuSoal()
	default:
		menuCariSoalMateri()
	}

}

func MenuCariSoalCatatan() {
	var input int

	fmt.Println("=============================================")
	fmt.Println("= Menu Soal : Cari Soal Berdasarkan Catatan =")
	fmt.Println("=   Catatan : ketik '0' jika ingin kembali  =")
	fmt.Println("=============================================")
	fmt.Print("Pilih ID Catatan : ")
	fmt.Scan(&input)

	if input == 0 {
		menuSoal()
	} else if seqsearchcat(input) == -1 {
		fmt.Println("ID tidak ditemukan")
		menuCariSoalMateri()
	} else {
		genSoal(cat[seqsearchcat(input)].topik)
		menuCariCatatan()
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
	case 0:
		menuCatatan()
	default:
		tambahCatatan(input, j)
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

// menampilkan seluruh catatan
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
	var i, j, min int

	// selection sort berdasarkan tahun, bulan, dan tanggal
	for i = 0; i < icat-1; i++ {
		min = i
		for j = i + 1; j < icat; j++ {
			// bandingkan tahun dulu
			if cat[j].tanggal.tahun < cat[min].tanggal.tahun ||
				(cat[j].tanggal.tahun == cat[min].tanggal.tahun && cat[j].tanggal.bulan < cat[min].tanggal.bulan) ||
				(cat[j].tanggal.tahun == cat[min].tanggal.tahun && cat[j].tanggal.bulan == cat[min].tanggal.bulan && cat[j].tanggal.hari < cat[min].tanggal.hari) {
				min = j
			}
		}
		// swap seluruh elemen
		cat[i], cat[min] = cat[min], cat[i]
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

func tingkatKesulitan() string {
	var input int
	fmt.Println("========================")
	fmt.Println(" 1. Mudah")
	fmt.Println(" 2. Sedang")
	fmt.Println(" 3. Sulit")
	fmt.Println("========================")
	fmt.Print("Pilih Tingkat kesulitan : ")
	fmt.Scan(&input)
	if input == 1 {
		return "Mudah"
	} else if input == 2 {
		return "Sedang"
	} else if input == 3 {
		return "Sulit"
	} else {
		return tingkatKesulitan()
	}
}

// Soal

func genSoal(topik string) {
	var tingkat string
	var i, n int
	tingkat = tingkatKesulitan()
	switch topik {
	case "Matematika":
		i = 0
		n = 4
	case "Bahasa Indonesia":
		i = 5
		n = 9
	case "Bahasa Inggris":
		i = 10
		n = 14
	case "IPA":
		i = 15
		n = 19
	case "IPS":
		i = 20
		n = 24
	}

	for i <= n && tingkat != Soal[i].tingkat {
		i++
	}

	fmt.Println("============================")
	fmt.Println("=   Pembuat Soal Otomatis  =")
	fmt.Println("============================")
	tampilanSoal(i)
	cekJawaban(i)
	menuSoal()

}

func tampilanSoal(i int) {
	var j int
	fmt.Println(Soal[i].pertanyaan)
	for j = 0; j < 4; j++ {
		fmt.Printf("Pilihan %v : %v\n", j+1, Soal[i].pilihan[j])
	}
}
func cekJawaban(i int) {
	var pilihan int
	fmt.Print("Pilihan Jawaban Anda : ")
	fmt.Scan(&pilihan)
	if Soal[i].pilihan[pilihan-1] == Soal[i].jawaban {
		fmt.Println("Jawaban Anda Benar")
	} else {
		fmt.Println("Jawaban Anda Salah")
		cekJawaban(i)
	}
}
