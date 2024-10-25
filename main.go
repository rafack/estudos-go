package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Defina o preço unitário da maçã:")

	var precoUnitarioMaca string
	// precoUnitarioMaca := ""
	fmt.Scan(&precoUnitarioMaca)

	precoUnitarioMacaFloat, err := strconv.ParseFloat(precoUnitarioMaca, 64)
	if err != nil {
		fmt.Println("Erro ao converter preço da maçã de string pra decimal", err.Error())
		return
	}

	fmt.Println("Quantos reais de maçã a cliente deseja?")

	var pedidoCliente string
	fmt.Scan(&pedidoCliente) // R$10.5

	valorPagamento := strings.TrimPrefix(pedidoCliente, "R$")

	valorPgtoFloat, err := strconv.ParseFloat(valorPagamento, 64)
	if err != nil {
		fmt.Println("Erro ao converter o valor do pagamento do cliente de string pra decimal", err.Error())
		return
	}

	qtdMacasAComprar := valorPgtoFloat / precoUnitarioMacaFloat

	valorPgtoFormatado := strconv.FormatFloat(valorPgtoFloat, 'f', 2, 64)

	qtdMacasInt := int(qtdMacasAComprar)
	fmt.Printf("A cliente pode comprar %d maçãs com o valor de R$%s", qtdMacasInt, valorPgtoFormatado)
}
