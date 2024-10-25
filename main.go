package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Você vende maçãs. Defina o preço da unidade:")

	var precoUnidadeMacaStr string
	fmt.Scan(&precoUnidadeMacaStr)

	precoUnidadeMaca, err := strconv.ParseFloat(precoUnidadeMacaStr, 64)
	if err != nil {
		fmt.Println("erro ao converter o preço unitário da maçã de string para float64")
		return
	}

	fmt.Println("Uma cliente chegou. Quantos reais de maçã ela quer?")

	var pedidoCliente string
	fmt.Scan(&pedidoCliente) // R$10.50

	valorPagamentoStr := strings.TrimPrefix(pedidoCliente, "R$")

	valorPagamento, err := strconv.ParseFloat(valorPagamentoStr, 64)
	if err != nil {
		fmt.Println("erro ao converter valor do pagamento de string para float64")
		return
	}

	// quantas maçãs a cliente conseguirá comprar?
	qtdMacasCompradas := valorPagamento / precoUnidadeMaca // qual o tipo dessa variável resultado da divisão?

	fmt.Printf("A cliente pode comprar %f maçãs com o valor de %f", qtdMacasCompradas, valorPagamento)

	qtdMacasCompradasInt := int(qtdMacasCompradas)

	fmt.Printf("A cliente pode comprar %d maçãs com o valor de %f", qtdMacasCompradasInt, valorPagamento)
}
