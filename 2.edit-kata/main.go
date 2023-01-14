package main

import (
	"fmt"
	"math"
)

func isOnceEdited(firstWord, secondWord string) bool {
	const (
		MAX_DIFFERENCE = 1
	)

	// cek panjang masing-masing kata
	lengthFirstWord := len(firstWord)
	lengthSecondWord := len(secondWord)

	// jika selisih panjang kedua kata lebih dari 1
	// maka return false
	lengthDifference := lengthFirstWord - lengthSecondWord
	absoluteDifference := math.Abs(float64(lengthDifference))

	if absoluteDifference > MAX_DIFFERENCE {
		return false
	}

	// mapping masing-masing huruf dari inputan
	mapFirstWord := make(map[string]string)
	mapSecondWord := make(map[string]string)

	// proses mapping
	for _, v := range firstWord {
		mapFirstWord[string(v)] = string(v)
	}
	for _, v := range secondWord {
		mapSecondWord[string(v)] = string(v)
	}

	// hitung ada berapakan perbedaan huruf antara dua inputan
	// jika lebih dari 1 maka return false
	totalDifference := 0
	for _, v := range firstWord {
		_, ok := mapSecondWord[string(v)]
		if !ok {
			totalDifference++
		}
	}
	for _, v := range secondWord {
		_, ok := mapFirstWord[string(v)]
		if !ok {
			totalDifference++
		}
	}

	if totalDifference > MAX_DIFFERENCE {
		return false
	}

	return true

}

func main() {
	fmt.Println(isOnceEdited("telkom", "telecom"))
	fmt.Println(isOnceEdited("telkom", "tlkom"))
}
