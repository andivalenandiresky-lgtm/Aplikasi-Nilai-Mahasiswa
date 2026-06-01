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

