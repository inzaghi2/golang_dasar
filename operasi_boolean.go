package main

import "fmt"

func main() {
	var (
		nilaiakhir            = 90
		absensi               = 80
		lulusNiilaiAkhir bool = nilaiakhir >= 80
		lulusAbsensi     bool = absensi >= 80

		lulus bool = lulusNiilaiAkhir && lulusAbsensi
	)

	fmt.Println(lulus)
}
