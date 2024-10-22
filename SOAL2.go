package main

import (
	"fmt"
)

func urutkanIndeksEkonomi(indeksEkonomi []float64) {
	for i := 0; i < len(indeksEkonomi); i++ {
		minIndex := i
		for j := i + 1; j < len(indeksEkonomi); j++ {
			if indeksEkonomi[j] < indeksEkonomi[minIndex]{
				minIndex = j 
			}
		}
		indeksEkonomi[i], indeksEkonomi[minIndex] = indeksEkonomi[minIndex], indeksEkonomi[i]
	}
}

func main() {
	indeks := []float64{5.2, 4.8, 6.5, 4.3, 5.9}
	
	urutkanIndeksEkonomi(indeks)

	fmt.Println(indeks)
}