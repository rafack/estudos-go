package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Insira uma lista de números inteiros, separados por vírgula:")

	var listaNumerosInput string
	fmt.Scan(&listaNumerosInput)

	numeros := strings.Split(listaNumerosInput, ",")

	fmt.Println(numeros)

	// exemplo de lista
	//[6 3 1 5 7]
	// 0 1 2 3 4 (índices)
	//numeros[0] - 6
	//numeros[1] - 3

	resultado := numeros[0] + numeros[1]
	fmt.Println("Resultado:", resultado)

	numerosInt := make([]int, 0, len(numeros)) // https://go.dev/tour/moretypes/13
	for _, n := range numeros {
		numInt, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("erro ao tentar converter %s para inteiro\n", n)
			continue
		}
		numerosInt = append(numerosInt, numInt)
	}

	fmt.Println(numerosInt)

	res := numerosInt[0] + numerosInt[1]
	fmt.Println(res)
}
