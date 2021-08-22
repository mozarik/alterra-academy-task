package main

import (
	"fmt"
	"math"
)

func main() {
	var T float64
	var r float64

	fmt.Scanf("%f", &T)
	fmt.Scanf("%f", &r)

	luasPermukaanTabung := HitungLuasTabung(T, r)

	fmt.Println(luasPermukaanTabung)
}

func HitungLuasTabung(t, r float64) (result float64) {
	result = 2 * math.Pi * r * (r + t)
	return
}
