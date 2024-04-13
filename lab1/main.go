package main

import (
	"fmt"
	"math/rand"
)

func contains(slice []int, element int) bool {
	for _, item := range slice {
		if item == element {
			return false
		}
	}
	return true
}

func gra(liczba_rund int, strategia bool) {
	var wygrane int = 0
	for i := 0; i <= liczba_rund; i++ {
		var kubeczek int = rand.Intn(3)
		var wybor_gracza int = rand.Intn(3)
		var wybor_hosta int
		for {
			var result = rand.Intn(3)
			if result != kubeczek && result != wybor_gracza {
				wybor_hosta = result
				break
			}
		}
		if strategia {
			for {
				var result = rand.Intn(3)
				if result != wybor_gracza && result != wybor_hosta {
					wybor_gracza = result
					break
				}
			}
		}
		if wybor_gracza == kubeczek {
			wygrane++
		}
	}
	fmt.Println("wygrane", wygrane)
}

func gra2(liczbaRund int, strategia bool, iloscKubeczkow int, odslanianeKubeczki int) {
	var wygrane int = 0

	for i := 0; i < liczbaRund; i++ {
		var ukrytyKubeczek int = rand.Intn(iloscKubeczkow)
		var wyborGracza int = rand.Intn(iloscKubeczkow)
		// fmt.Println("ukrytyKubeczek: ", ukrytyKubeczek, " wyborGracza: ", wyborGracza)

		var kubeczkiOdsloniete []int
		for len(kubeczkiOdsloniete) < odslanianeKubeczki {
			var result = rand.Intn(iloscKubeczkow)
			if result != ukrytyKubeczek && result != wyborGracza && contains(kubeczkiOdsloniete, result) {
				kubeczkiOdsloniete = append(kubeczkiOdsloniete, result)
			}
		}

		if strategia {
			for {
				var result = rand.Intn(iloscKubeczkow)
				if result != wyborGracza && contains(kubeczkiOdsloniete, result) {
					wyborGracza = result
					break
				}
			}
		}

		if wyborGracza == ukrytyKubeczek {
			wygrane++
		}
	}

	fmt.Println("wygrane", wygrane)
}

func main() {
	// var input int
	// var input2 bool

	// fmt.Print("Wpisz 1 argument: ")
	// fmt.Scanln(&input)
	// fmt.Print("Wpisz 2 argument: ")
	// fmt.Scanln(&input2)
	// gra(input, input2)

	gra2(1000, false, 100, 2)
	gra2(1000, true, 100, 2)

}
