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

type arrMhs [NMAX]Mahasiswa
type arrKRS [NMAX]KRS

func SeqSearch(T arrMhs, n int, X string) int {
	var i, idx int

	idx = -1
	i = 0

	for i < n && idx == -1 {

		if T[i].NIM == X {
			idx = i
		}

		i++
	}

	return idx
}
func BinarySearch(A arrMhs, n int, x string) int {
	var left, right, mid int

	left = 0
	right = n - 1

	for left <= right {

		mid = (left + right) / 2

		if A[mid].NIM == x {

			return mid

		} else if x < A[mid].NIM {

			right = mid - 1

		} else {

			left = mid + 1
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
func tambahMahasiswa(A *arrMhs, B *arrKRS, n *int) {

	if *n >= NMAX {

		fmt.Println("Data mahasiswa penuh (｡•́︿•̀｡)")

	} else {

		fmt.Print("Masukkan NIM : ")
		fmt.Scan(&A[*n].NIM)

	    fmt.Print("Masukkan Nama (tanpa spasi): ")
        fmt.Scan(&A[*n].Nama)

		B[*n].NIM = A[*n].NIM
		B[*n].Jml = 0

		*n = *n + 1

		fmt.Println("Data mahasiswa berhasil ditambahkan（๑ > ᴗ <๑）")
	}
}

func editMahasiswa(A *arrMhs, n int) {
	var nim string
	var idx int

	fmt.Print("Masukkan NIM yang akan diedit : ")
	fmt.Scan(&nim)

	idx = BinarySearch(*A, n, nim)

	if idx == -1 {

		fmt.Println("Mahasiswa tidak ditemukan (｡•́︿•̀｡)")

	} else {

	    fmt.Print("Nama baru : ")
        fmt.Scan(&A[idx].Nama)

		fmt.Println("Data mahasiswa berhasil diubah（๑ > ᴗ <๑）")
	}
}

func hapusMahasiswa(A *arrMhs, B *arrKRS, n *int) {
	var nim string
	var idx int
	var i int

	fmt.Print("Masukkan NIM yang akan dihapus : ")
	fmt.Scan(&nim)

	idx = BinarySearch(*A, *n, nim)

	if idx == -1 {

		fmt.Println("Mahasiswa tidak ditemukan (｡•́︿•̀｡)")

	} else {

		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
			B[i] = B[i+1]
		}

		*n--

		fmt.Println("Data mahasiswa berhasil dihapus（๑ > ᴗ <๑）")
	}
}

func tampilSemuaMahasiswa(A arrMhs, B arrKRS, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data mahasiswa (｡•́︿•̀｡)")
	} else {
		fmt.Println("============================================================================")
		fmt.Printf("%-15s %-25s %-15s %-15s\n",
			"NIM", "Nama", "Total SKS", "Rata-rata")
		fmt.Println("============================================================================")

		for i = 0; i < n; i++ {
			fmt.Printf("%-15s %-25s %-15d %-15.2f\n",
				A[i].NIM,
				A[i].Nama,
				totalSKS(B[i]),
				rataNilai(B[i]))
		}

		fmt.Println("============================================================================")
	}
}

func tambahMK(A arrMhs, B *arrKRS, n int) {
	var nim string
	var idx int
	var j int

	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)

	idx = SeqSearch(A, n, nim)

	if idx == -1 {

		fmt.Println("Mahasiswa tidak ditemukan (｡•́︿•̀｡)")

	} else {

		j = B[idx].Jml

		if j >= MAXMK {

			fmt.Println("Mata kuliah sudah penuh (｡•́︿•̀｡)")

		} else {

			fmt.Print("Kode MK (tanpa spasi): ")
			fmt.Scan(&B[idx].MK[j].Kode)

			fmt.Print("Nama MK (tanpa spasi): ")
			fmt.Scan(&B[idx].MK[j].Nama)

			fmt.Print("SKS : ")
			fmt.Scan(&B[idx].MK[j].SKS)

			fmt.Print("Nilai Quiz : ")
			fmt.Scan(&B[idx].MK[j].Quiz)

			fmt.Print("Nilai UTS : ")
			fmt.Scan(&B[idx].MK[j].UTS)

			fmt.Print("Nilai UAS : ")
			fmt.Scan(&B[idx].MK[j].UAS)

			B[idx].MK[j].Total =
				totalNilai(
					B[idx].MK[j].Quiz,
					B[idx].MK[j].UTS,
					B[idx].MK[j].UAS)

			B[idx].MK[j].Grade =
				grade(B[idx].MK[j].Total)

			B[idx].Jml++

			fmt.Println("Mata kuliah berhasil ditambahkan（๑ > ᴗ <๑）")
		}
	}
}

func editMK(A arrMhs, B *arrKRS, n int) {
	var nim string
	var kode string
	var idx int
	var j int
	var ketemu bool

	fmt.Print("Masukkan NIM : ")
	fmt.Scan(&nim)

	idx = SeqSearch(A, n, nim)

	if idx == -1 {

		fmt.Println("Mahasiswa tidak ditemukan (｡•́︿•̀｡)")

	} else {

		fmt.Print("Masukkan Kode MK : ")
		fmt.Scan(&kode)

		ketemu = false
		j = 0

		for j < B[idx].Jml && !ketemu {

			if B[idx].MK[j].Kode == kode {

				fmt.Print("Nilai Quiz Baru : ")
				fmt.Scan(&B[idx].MK[j].Quiz)

				fmt.Print("Nilai UTS Baru : ")
				fmt.Scan(&B[idx].MK[j].UTS)

				fmt.Print("Nilai UAS Baru : ")
				fmt.Scan(&B[idx].MK[j].UAS)

				B[idx].MK[j].Total =
					totalNilai(
						B[idx].MK[j].Quiz,
						B[idx].MK[j].UTS,
						B[idx].MK[j].UAS)

				B[idx].MK[j].Grade =
					grade(B[idx].MK[j].Total)

				fmt.Println("Data mata kuliah berhasil diubah（๑ > ᴗ <๑）")

				ketemu = true
			}

			j++
		}

		if !ketemu {
			fmt.Println("Mata kuliah tidak ditemukan (｡•́︿•̀｡)")
		}
	}
}
