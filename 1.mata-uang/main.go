package main

import (
	"log"
	"strconv"
)

const (
	RATUS_RIBUAN = 6
	PULUH_RIBUAN = 5
	RIBUAN       = 4
	RATUSAN      = 3
	PULUHAN      = 2
)

func changeMoney(money int) {
	// disini saya memberi batasan input
	// maksimal input adalah 999999
	lengthMoney := checkLength(money)
	if lengthMoney > RATUS_RIBUAN {
		log.Println("maksimal nominal adalah 999999")
		return
	}

	// cek dulu apakah money termasuk
	// ratus ribuan, puluh ribuan, ribuan, ratusan
	// atau tidak semuanya
	if lengthMoney == RATUS_RIBUAN {
		// ambil angka paling depan untuk mengetahui
		// nilai ratus ribuan
		firstNumber := getFisrtNumber(money)
		nilaiRatusribuan := firstNumber * 100000

		// kurangi money dengan nilai ratus ribuan
		afterElimination := money - nilaiRatusribuan

		log.Printf("Rp. %v : %v", 10000, firstNumber)

		// rekursif
		changeMoney(afterElimination)
	} else if lengthMoney == PULUH_RIBUAN {
		// ambil angka pertama
		firstNumber := getFisrtNumber(money)
		nilaiPuluhribuan := firstNumber * 10000

		// kurangi money dengan nilai puluh ribuan
		afterElimination := money - nilaiPuluhribuan

		// karena ada 3 pilihan untuk pecahan puluh ribuan
		// yaitu 10000, 20000, 50000
		// maka cari selisih yang paling kecil dari nilaiPuluhribuan
		minimumDistance := getMinimumDistance(money, 10000, 20000, 50000)

		// total pecahan yang harus dikeluarkan
		total := nilaiPuluhribuan / minimumDistance
		log.Printf("Rp. %v : %v", minimumDistance, total)

		changeMoney(afterElimination)
	} else if lengthMoney == RIBUAN {
		// ambil angka pertama
		firstNumber := getFisrtNumber(money)
		nilaiRibuan := firstNumber * 1000

		// kurangi money dengan nilai puluh ribuan
		afterElimination := money - nilaiRibuan

		// karena ada 3 pilihan untuk pecahan ribuan
		// yaitu 1000, 2000, 5000
		// maka cari selisih yang paling kecil dari nilaiPuluhribuan
		minimumDistance := getMinimumDistance(money, 1000, 2000, 5000)

		// total pecahan yang harus dikeluarkan
		total := nilaiRibuan / minimumDistance
		log.Printf("Rp. %v : %v", minimumDistance, total)

		changeMoney(afterElimination)
	} else if lengthMoney == RATUSAN {
		// ambil angka pertama
		firstNumber := getFisrtNumber(money)
		nilaiRatusan := firstNumber * 100

		// kurangi money dengan nilai puluh ribuan
		afterElimination := money - nilaiRatusan

		// karena ada 3 pilihan untuk pecahan puluh ratusan
		// yaitu 100, 200, 500
		// maka cari selisih yang paling kecil dari nilaiPuluhribuan
		minimumDistance := getMinimumDistance(money, 100, 200, 500)

		// total pecahan yang harus dikeluarkan
		total := nilaiRatusan / minimumDistance
		log.Printf("Rp. %v : %v", minimumDistance, total)

		changeMoney(afterElimination)
	} else {
		if money != 0 {
			log.Printf("Rp %v : %v", 100, 1)
		}
	}

}

func checkLength(money int) int {
	moneyString := strconv.Itoa(money)
	return len(moneyString)
}

func getFisrtNumber(money int) int {
	moneyString := strconv.Itoa(money)
	tmp := moneyString[0]
	firstNumber, _ := strconv.Atoi(string(tmp))
	return firstNumber
}

func getMinimumDistance(money int, fisrtNominal, secondNominal, thirdNominal int) int {
	mapNomional := map[int]int{
		fisrtNominal:  0,
		secondNominal: 0,
		thirdNominal:  0,
	}

	// pengurangan
	for key := range mapNomional {
		distance := money - key
		mapNomional[key] = distance
	}

	sliceDistance := []int{
		mapNomional[fisrtNominal],
		mapNomional[secondNominal],
		mapNomional[thirdNominal],
	}

	// minimum
	minimum := mapNomional[fisrtNominal]

	for _, v := range sliceDistance {
		if v == 0 {
			minimum = v
			break
		}

		if v < 0 {
			continue
		}

		if v < minimum {
			minimum = v
		}
	}

	// minimumPuluhanRibu
	result := 0
	for key, v := range mapNomional {
		if v == minimum {
			result = key
		}
	}

	return result

}

func main() {
	changeMoney(145000)
}
