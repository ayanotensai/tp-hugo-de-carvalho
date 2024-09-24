package main

import "fmt"

func Ft_manquant(nombres []int) int {
	n := len(nombres)
	sommeAttendue := (n * (n + 1)) / 2
	sommeReelle := 0

	for _, nombre := range nombres {
		sommeReelle += nombre
	}

	return sommeAttendue - sommeReelle
}

func main() {
	fmt.Println(Ft_manquant([]int{3, 1, 2})) // résultat : 0
	fmt.Println(Ft_manquant([]int{0, 1}))    // résultat : 2
	fmt.Println(Ft_manquant([]int{0, 2, 3})) // résultat : 1
	fmt.Println(Ft_manquant([]int{1, 2, 3})) // résultat : 0
	fmt.Println(Ft_manquant([]int{0}))       // résultat : 1
}
