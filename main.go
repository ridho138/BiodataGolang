package main

import (
	"fmt"
	"os"
	"strconv"
)

type Datas struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {

	absen, err := strconv.ParseInt(os.Args[1], 10, 0)

	if err != nil {
		fmt.Println("First input parameter must be integer")
		os.Exit(1)
	}

	findClassMate(absen - 1)
}

func findClassMate(absen int64) {

	var classMate = []Datas{
		{name: "ABC", address: "JKT001", job: "JOB001", reason: "REASON001"},
		{name: "DEF", address: "JKT002", job: "JOB002", reason: "REASON002"},
		{name: "GHI", address: "JKT003", job: "JOB003", reason: "REASON003"},
	}

	if len(classMate) > int(absen) {
		fmt.Println("Nama : ", classMate[absen].name)
		fmt.Println("Alamat : ", classMate[absen].address)
		fmt.Println("Pekerjaan : ", classMate[absen].job)
		fmt.Println("Alasan : ", classMate[absen].reason)
	} else {
		fmt.Println("Absen Not Found")
	}
}
