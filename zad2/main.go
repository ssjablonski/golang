package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
		random := rand.Intn(10) + 1
		if forest[y][x] == 1 || forest[y][x] == 2 || forest[y][x] == 3 {
			count++
		} else if random < 2 {
			forest[y][x] = 1
		} else if 2 <= random && random <= 8 { // 20% mlode drzewa, 70% srednie, 10% stare
			forest[y][x] = 2 // 80% na spalenie, 90%, 100%
		} else {
			forest[y][x] = 3
		}
	}
}

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

	random := rand.Intn(10)
	if forest[i][j] == 1 {
		if random <= 6 {
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
			fmt.Println("Drzewo młode na pozycji", i, j, "ocalone!")
		}
	} else if forest[i][j] == 2 {
		if random <= 7 {
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
			fmt.Println("Drzewo srednie na pozycji", i, j, "ocalone!")
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

func testBestForestRatio() (int, float64) {

	var ratios = make(map[int]float64)
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 400; j++ {
			las := createForrest(20, 20)
			plantTree(las, j)
			var yD, xD int
			var hit bool
			for {
				yD, xD, hit = thunderBolt(las)
				if hit {
					break
				}
			}
			startFire(las, yD, xD) // funkcja taka sama jak burningForest, ale podpala drzewo nie wazne czy mlode srednie czy stare
			burnedTrees := calculateBurnedTrees(las)
			ratios[j] += float64(burnedTrees) / float64(j) // calculate the ratio of burned trees
		}
	}

	for x := 1; x <= 100; x++ {
		ratios[x] = ratios[x] / 10000 // calculate the average ratio
	}

	var bestKey int
	var bestRatio float64 = math.MaxFloat64
	for key, value := range ratios {
		if value < bestRatio {
			bestKey = key
			bestRatio = value
		}
	}
	fmt.Println(ratios)
	return bestKey, bestRatio
}

func main() {
	fmt.Println("Wybierz jedna z opcji: ")
	fmt.Println("1. Sprawdzanie najlepszego stosunku spalonych drzew do ilosci drzew w lesie")
	fmt.Println("2. Symulacja pożaru w lesie")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		startTime := time.Now()
		bestKey, bestRatio := testBestForestRatio()
		elapsedTime := time.Since(startTime)
		fmt.Println("Czas wykonania funkcji: ", elapsedTime)
		fmt.Println("Najlepszy stosunek jest dla: ", bestKey, " drzew, stosunek: ", bestRatio)
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
