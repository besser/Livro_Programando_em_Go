package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numArgs := len(os.Args)

	if numArgs < 3 {
		fmt.Println("Uso: conversor <valores> <unidade> [celsius, fahrenheit, km ou mi]")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[numArgs-1]
	valoresOrigem := os.Args[1 : numArgs-1]

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "km" {
		unidadeDestino = "mi"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", unidadeOrigem)
		os.Exit(1)
	}

	fmt.Printf("\n");

	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = (valorOrigem * 1.8) + 32
		} else {
			valorDestino = valorOrigem / 1.609344
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}

	fmt.Printf("\n");
}
