package main

import "fmt"

const NMAX int = 100
const MAXMK int = 20

type Mahasiswa struct {
	NIM  string
	Nama string
}

type NilaiMK struct {
	Kode  string
	Nama  string
	Grade string

	SKS int

	Quiz float64
	UTS  float64
	UAS  float64

	Total float64
}

type KRS struct {
	NIM string
	Jml int
	MK  [MAXMK]NilaiMK
}

var mhs [NMAX]Mahasiswa
var krs [NMAX]KRS
var n int

func seqSearch(nim string) int {
	var i int

	for i = 0; i < n; i++ {
		if mhs[i].NIM == nim {
			return i
		}
	}

	return -1
}

func totalNilai(quiz, uts, uas float64) float64 {
	var total float64

	total = (0.2 * quiz) +
		(0.3 * uts) +
		(0.5 * uas)

	return total
}

func grade(total float64) string {

	if total >= 80 {
		return "A"
	} else if total >= 75 {
		return "AB"
	} else if total >= 70 {
		return "B"
	} else if total >= 65 {
		return "BC"
	} else if total >= 60 {
		return "C"
	} else if total >= 50 {
		return "D"
	}

	return "E"
}

func totalSKS(data KRS) int {
	var i int
	var total int

	total = 0

	for i = 0; i < data.Jml; i++ {
		total = total + data.MK[i].SKS
	}

	return total
}

func rataNilai(data KRS) float64 {
	var i int
	var total float64

	if data.Jml == 0 {
		return 0
	}

	total = 0

	for i = 0; i < data.Jml; i++ {
		total = total + data.MK[i].Total
	}

	return total / float64(data.Jml)
}

func swap(i, j int) {
	var tempM Mahasiswa
	var tempK KRS

	tempM = mhs[i]
	mhs[i] = mhs[j]
	mhs[j] = tempM

	tempK = krs[i]
	krs[i] = krs[j]
	krs[j] = tempK
}
func tambahMahasiswa() {

	if n >= NMAX {
		fmt.Println("Data mahasiswa penuh")
		return
	}

	fmt.Print("Masukkan NIM  : ")
	fmt.Scan(&mhs[n].NIM)

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&mhs[n].Nama)

	krs[n].NIM = mhs[n].NIM
	krs[n].Jml = 0

	n++

	fmt.Println("Data mahasiswa berhasil ditambahkan")
}

func editMahasiswa() {
	var nim string
	var idx int

	fmt.Print("Masukkan NIM yang akan diedit : ")
	fmt.Scan(&nim)

	idx = seqSearch(nim)

	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan")
		return
	}

	fmt.Print("Nama baru : ")
	fmt.Scan(&mhs[idx].Nama)

	fmt.Println("Data mahasiswa berhasil diubah")
}

func hapusMahasiswa() {
	var nim string
	var idx int
	var i int

	fmt.Print("Masukkan NIM yang akan dihapus : ")
	fmt.Scan(&nim)

	idx = seqSearch(nim)

	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan")
		return
	}

	for i = idx; i < n-1; i++ {
		mhs[i] = mhs[i+1]
		krs[i] = krs[i+1]
	}

	n--

	fmt.Println("Data mahasiswa berhasil dihapus")
}

func tampilSemuaMahasiswa() {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data mahasiswa")
		return
	}

	fmt.Println("========================================================")
	fmt.Println("NIM\tNama\t\tTotal SKS\tRata-rata")
	fmt.Println("========================================================")

	for i = 0; i < n; i++ {

		fmt.Println(
			mhs[i].NIM,
			"\t",
			mhs[i].Nama,
			"\t\t",
			totalSKS(krs[i]),
			"\t\t",
			rataNilai(krs[i]))
	}

	fmt.Println("========================================================")
}
func tambahMK() {
	var nim string
	var idx int
	var j int

	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)

	idx = seqSearch(nim)

	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan")
		return
	}

	j = krs[idx].Jml

	if j >= MAXMK {
		fmt.Println("Mata kuliah sudah penuh")
		return
	}

	fmt.Print("Kode MK : ")
	fmt.Scan(&krs[idx].MK[j].Kode)

	fmt.Print("Nama MK : ")
	fmt.Scan(&krs[idx].MK[j].Nama)

	fmt.Print("SKS : ")
	fmt.Scan(&krs[idx].MK[j].SKS)

	fmt.Print("Nilai Quiz : ")
	fmt.Scan(&krs[idx].MK[j].Quiz)

	fmt.Print("Nilai UTS : ")
	fmt.Scan(&krs[idx].MK[j].UTS)

	fmt.Print("Nilai UAS : ")
	fmt.Scan(&krs[idx].MK[j].UAS)

	krs[idx].MK[j].Total =
		totalNilai(
			krs[idx].MK[j].Quiz,
			krs[idx].MK[j].UTS,
			krs[idx].MK[j].UAS)

	krs[idx].MK[j].Grade =
		grade(krs[idx].MK[j].Total)

	krs[idx].Jml++

	fmt.Println("Mata kuliah berhasil ditambahkan")
}

func editMK() {
	var nim string
	var kode string
	var idx int
	var j int

	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)

	idx = seqSearch(nim)

	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan")
		return
	}

	fmt.Print("Masukkan Kode MK : ")
	fmt.Scan(&kode)

	for j = 0; j < krs[idx].Jml; j++ {

		if krs[idx].MK[j].Kode == kode {

			fmt.Print("Nilai Quiz Baru : ")
			fmt.Scan(&krs[idx].MK[j].Quiz)

			fmt.Print("Nilai UTS Baru : ")
			fmt.Scan(&krs[idx].MK[j].UTS)

			fmt.Print("Nilai UAS Baru : ")
			fmt.Scan(&krs[idx].MK[j].UAS)

			krs[idx].MK[j].Total =
				totalNilai(
					krs[idx].MK[j].Quiz,
					krs[idx].MK[j].UTS,
					krs[idx].MK[j].UAS)

			krs[idx].MK[j].Grade =
				grade(krs[idx].MK[j].Total)

			fmt.Println("Data mata kuliah berhasil diubah")
			return
		}
	}

	fmt.Println("Mata kuliah tidak ditemukan")
}

func hapusMK() {
	var nim string
	var kode string
	var idx int
	var j int
	var k int

	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)

	idx = seqSearch(nim)

	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan")
		return
	}
