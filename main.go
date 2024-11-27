package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("Digite a operação com dois números (exemplo: 3+5):")
	fmt.Scanln(&input)

	operatorIndex := strings.IndexAny(input, "+-*/")
	if operatorIndex == -1 {
		fmt.Println("Operação inválida! Por favor, use um dos seguintes operadores: +, -, *, /")
		return
	}

	n1Str := input[:operatorIndex]
	op := input[operatorIndex : operatorIndex+1]
	n2Str := input[operatorIndex+1:]

	n1, errN1 := strconv.ParseFloat(n1Str, 64)
	n2, errN2 := strconv.ParseFloat(n2Str, 64)
	if errN1 != nil || errN2 != nil {
		fmt.Println("Erro ao interpretar os números, por favor, use apenas números válidos.")
		return
	}

	var result float64
	var err error

	switch op {
	case "+":
		result = sum(n1, n2)
	case "-":
		result = sub(n1, n2)
	case "*":
		result = mul(n1, n2)
	case "/":
		result, err = div(n1, n2)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
	default:
		fmt.Println("Operação inválida.")
	}

	if result == float64(int(result)) {
		fmt.Println("=", int(result))
	} else {
		fmt.Println("=", result)
	}
}

func sum(n1, n2 float64) float64 {
	return n1 + n2
}

func sub(n1, n2 float64) float64 {
	return n1 - n2
}

func mul(n1, n2 float64) float64 {
	return n1 * n2
}

func div(n1, n2 float64) (float64, error) {
	if n2 == 0.0 {
		return 0, errors.New("divisão por zero não é permitida")
	}

	return n1 / n2, nil
}
