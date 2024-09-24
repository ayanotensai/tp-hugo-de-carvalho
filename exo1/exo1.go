package main

import (
	"fmt"
	"math"
)

func Ft_coin(coins []int, montant int) int {
	piece := make([]int, montant+1)

	for i := range piece {
		piece[i] = montant + 1
	}

	piece[0] = 0

	for i := 1; i <= montant; i++ {
		for _, coin := range coins {
			if i >= coin {
				piece[i] = int(math.Min(float64(piece[i]), float64(piece[i-coin]+1)))
			}
		}
	}

	if piece[montant] > montant {
		return -1
	}

	return piece[montant]
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11))
	fmt.Println(Ft_coin([]int{2}, 3))
	fmt.Println(Ft_coin([]int{1}, 0))
}
