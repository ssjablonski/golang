package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func displayForrest(forest [][]int) {
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			fmt.Print(forest[i][j], " ")
		}
		fmt.Println()
	}
}

func createForrest(x int, y int) [][]int {
	forest := make([][]int, x)
	for i := 0; i < x; i++ {
		forest[i] = make([]int, y)
		for j := 0; j < y; j++ {
			forest[i][j] = 0
		}
	}

	return forest
}

func plantTree(forest [][]int, count int) {
	for i := 0; i < count; i++ {
		y := rand.Intn(len(forest))
		x := rand.Intn(len(forest[y]))
		random := rand.Intn(10)
		if forest[y][x] == 1 || forest[y][x] == 2 || forest[y][x] == 3 {
			count++
		} else if random < 1 {
			forest[y][x] = 1
		} else if 2 <= random && random < 4 { // 10% mlode drzewa, 20% srednie, 70% stare
			forest[y][x] = 2 // 80% na spalenie, 90%, 100%
		} else {
			forest[y][x] = 3
		}
	}
}

// func plantTree(forest [][]int, count int) {
// 	for i := 0; i < count; i++ {
// 		y := rand.Intn(len(forest))
// 		x := rand.Intn(len(forest[y]))
// 		if forest[y][x] == 1 {
// 			count++
// 		} else {
// 			forest[y][x] = 1
// 		}
// 	}
// }

func thunderBolt(forest [][]int) (int, int, bool) {
	y := rand.Intn(len(forest))
	x := rand.Intn(len(forest[y]))
	if forest[y][x] == 1 || forest[y][x] == 2 || forest[y][x] == 3 {
		// fmt.Println("Thunderbolt trafia w drzewo na pozycji", y, x)
		return y, x, true
	} else {
		// fmt.Println("Thunderbolt trafia w puste miejsce na pozycji", y, x)
		return y, x, false
	}

}

func startFire(forest [][]int, i int, j int) {
	if i < 0 || i >= len(forest) || j < 0 || j >= len(forest[i]) || (forest[i][j] != 1 && forest[i][j] != 2 && forest[i][j] != 3) {
		return
	}

	forest[i][j] = 4
	burningForest(forest, i-1, j)
	burningForest(forest, i+1, j)
	burningForest(forest, i, j-1)
	burningForest(forest, i, j+1)
	burningForest(forest, i-1, j-1)
	burningForest(forest, i-1, j+1)
	burningForest(forest, i+1, j-1)
	burningForest(forest, i+1, j+1)
}

func burningForest(forest [][]int, i int, j int) {
	if i < 0 || i >= len(forest) || j < 0 || j >= len(forest[i]) || (forest[i][j] != 1 && forest[i][j] != 2 && forest[i][j] != 3) {
		return
	}

	random := rand.Intn(100)
	if forest[i][j] == 1 {
		if random <= 79 {
			forest[i][j] = 4
			burningForest(forest, i-1, j)
			burningForest(forest, i+1, j)
			burningForest(forest, i, j-1)
			burningForest(forest, i, j+1)
			burningForest(forest, i-1, j-1)
			burningForest(forest, i-1, j+1)
			burningForest(forest, i+1, j-1)
			burningForest(forest, i+1, j+1)
		} else {
			// fmt.Println("Drzewo młode na pozycji", i, j, "ocalone!")
		}
	} else if forest[i][j] == 2 {
		if random <= 89 {
			forest[i][j] = 4
			burningForest(forest, i-1, j)
			burningForest(forest, i+1, j)
			burningForest(forest, i, j-1)
			burningForest(forest, i, j+1)
			burningForest(forest, i-1, j-1)
			burningForest(forest, i-1, j+1)
			burningForest(forest, i+1, j-1)
			burningForest(forest, i+1, j+1)
		} else {
			// fmt.Println("Drzewo srednie na pozycji", i, j, "ocalone!")
		}
	} else if forest[i][j] == 3 {
		forest[i][j] = 4
		burningForest(forest, i-1, j)
		burningForest(forest, i+1, j)
		burningForest(forest, i, j-1)
		burningForest(forest, i, j+1)
		burningForest(forest, i-1, j-1)
		burningForest(forest, i-1, j+1)
		burningForest(forest, i+1, j-1)
		burningForest(forest, i+1, j+1)
	}
}

// func burningForest(forest [][]int, i int, j int) {
// 	if i < 0 || i >= len(forest) || j < 0 || j >= len(forest[i]) || (forest[i][j] != 1) {
// 		return
// 	}

// 	if forest[i][j] == 1 {
// 		forest[i][j] = 4
// 		burningForest(forest, i-1, j)
// 		burningForest(forest, i+1, j)
// 		burningForest(forest, i, j-1)
// 		burningForest(forest, i, j+1)
// 		burningForest(forest, i-1, j-1)
// 		burningForest(forest, i-1, j+1)
// 		burningForest(forest, i+1, j-1)
// 		burningForest(forest, i+1, j+1)
// 	}

// }

func calculateBurnedTrees(forest [][]int) int {
	count := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			if forest[i][j] == 4 {
				count++
			}
		}
	}
	return count
}

func nowa() {
	var ratios = make(map[int]float64)
	var keys []int
	var minRatio float64 = math.MaxFloat64
	var optimalForestDensity int

	for i := 5; i <= 95; i += 5 {
		for j := 0; j < 1000; j++ {
			numTrees := int(float64(400) * (float64(i) / 100.0))
			las := createForrest(20, 20)
			plantTree(las, numTrees)
			var yD, xD int
			var hit bool
			for {
				yD, xD, hit = thunderBolt(las)
				if hit {
					break
				}
			}
			startFire(las, yD, xD)
			burnedTrees := calculateBurnedTrees(las)
			ratios[i] += float64(burnedTrees) / float64(numTrees)
		}
		ratios[i] = ratios[i] / 1000
		keys = append(keys, i)

		// Check if the current ratio is less than the minimum ratio found so far
		if ratios[i] < minRatio {
			minRatio = ratios[i]
			optimalForestDensity = i
		}
	}

	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println("Poziom zalesienia:", key, "%, współczynnik ratio:", ratios[key])
	}

	fmt.Println("Optymalny poziom zalesienia:", optimalForestDensity, "% z współczynnikiem ratio:", minRatio)
}

func main() {
	fmt.Println("Wybierz jedna z opcji: ")
	fmt.Println("1. Sprawdzanie najlepszego stosunku spalonych drzew do ilosci drzew w lesie z parametrem starych drzew.")
	fmt.Println("2. Symulacja pożaru w lesie z parametrem starych drzew.")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		nowa()
	} else if choice == 2 {
		var wymiarX int
		var wymiarY int
		var iloscDrzew int

		fmt.Println("Podaj wymiar lasu X: ")
		fmt.Scan(&wymiarX)
		fmt.Println("Podaj wymiar lasu Y: ")
		fmt.Scan(&wymiarY)
		fmt.Println("Podaj ilosc drzew: ")
		fmt.Scan(&iloscDrzew)
		for iloscDrzew > wymiarX*wymiarY {
			fmt.Println("Ilosc drzew nie moze byc wieksza niz ilosc miejsc w lesie!")
			fmt.Println("Podaj ilosc drzew: ")
			fmt.Scan(&iloscDrzew)
		}

		las := createForrest(wymiarX, wymiarY)
		plantTree(las, iloscDrzew)
		displayForrest(las)
		var yD, xD int
		var hit bool
		for {
			yD, xD, hit = thunderBolt(las)
			if hit {
				break
			}
		}
		fmt.Println("Pożar zaczyna się: ", yD, xD)
		startFire(las, yD, xD) // funkcja taka sama jak burningForest, ale podpala drzewo nie wazne czy mlode srednie czy stare
		displayForrest(las)
		count := calculateBurnedTrees(las)
		fmt.Println("Spalonych drzew: ", count)
		fmt.Println("Procent spalonych drzew: ", float64(count)/float64(iloscDrzew)*100, "%")
	}
}
