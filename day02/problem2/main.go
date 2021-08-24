package main

import "fmt"

func main() {
	var studentScore int = 80

	grade := cekNilai(studentScore)

	fmt.Println("Nilai ", grade)
}

func cekNilai(nilai int) string {
	if nilai >= 80 && nilai <= 100 {
		return "A"
	} else if nilai >= 65 && nilai <= 79 {
		return "B"
	} else if nilai >= 50 && nilai <= 64 {
		return "C"
	} else if nilai >= 35 && nilai <= 49 {
		return "D"
	} else if nilai >= 0 && nilai <= 34 {
		return "E"
	} else {
		return "Nilai tidak valid"
	}
}
