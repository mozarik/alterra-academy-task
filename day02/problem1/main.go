package main

import (
	"fmt"
	"math"
)

func main() {
	var T float64
	var r float64

	T = 20.0
	r = 4.0

	luasPermukaanTabung := HitungLuasTabung(T, r)

	fmt.Print(luasPermukaanTabung)
}

func HitungLuasTabung(t, r float64) (result float64) {
	result = 2 * math.Pi * r * (r + t)
	return
}
