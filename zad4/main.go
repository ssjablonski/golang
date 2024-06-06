package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Inicjalizacja kolektora
	c := colly.NewCollector()

	var allRows [][]string
	tableCount := 0

	// Funkcja wykonywana przy każdym znalezionym elemencie HTML <table> <tbody>
	c.OnHTML("table.wikitable > tbody", func(h *colly.HTMLElement) {
		tableCount++
		if tableCount != 2 {
			return
		}
		fmt.Println("Found the second table")

		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			var row []string
			el.ForEach("th, td", func(_ int, col *colly.HTMLElement) {
				row = append(row, strings.TrimSpace(col.Text))
			})
			if len(row) > 0 {
				allRows = append(allRows, row)
			}
		})
	})

	// Odwiedzenie strony
	c.Visit("https://pl.wikipedia.org/wiki/W%C5%82adcy_Polski")

	// Sprawdzenie czy dane zostały pobrane
	if len(allRows) == 0 {
		fmt.Println("No data found")
		return
	}

	// Zapis danych do pliku CSV
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Zapis danych
	for _, row := range allRows {
		writer.Write(row)
	}

	fmt.Println("Data has been written to output.csv")
}
