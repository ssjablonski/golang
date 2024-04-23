package main

import (
	"fmt"
	"math/rand"
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
		if forest[y][x] == 1 {
			count++
		} else {
			forest[y][x] = 1
		}
	}
}

func thunderBolt(forest [][]int) (int, int, bool) {
	y := rand.Intn(len(forest))
	x := rand.Intn(len(forest[y]))
	if forest[y][x] == 1 {
		fmt.Println("Thunderbolt trafia w drzewo na pozycji", y, x)
		return y, x, true
	} else {
		fmt.Println("Thunderbolt trafia w puste miejsce na pozycji", y, x)
		return y, x, false
	}

}

func burningForest(forest [][]int, i int, j int) {
	if i < 0 || i >= len(forest) || j < 0 || j >= len(forest[i]) || forest[i][j] != 1 {
		return
	}

	forest[i][j] = 3

	burningForest(forest, i-1, j)
	burningForest(forest, i+1, j)
	burningForest(forest, i, j-1)
	burningForest(forest, i, j+1)
	burningForest(forest, i-1, j-1)
	burningForest(forest, i-1, j+1)
	burningForest(forest, i+1, j-1)
	burningForest(forest, i+1, j+1)
}

func calculateBurnedTrees(forest [][]int) int {
	count := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			if forest[i][j] == 3 {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println("Hello, World!")
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
	displayForrest(las)
	fmt.Println("Pożar!")
	burningForest(las, yD, xD)
	displayForrest(las)
	count := calculateBurnedTrees(las)
	fmt.Println("Spalonych drzew: ", count)
	fmt.Println("Procent spalonych drzew: ", float64(count)/float64(iloscDrzew)*100, "%")
}
