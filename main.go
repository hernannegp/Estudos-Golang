package main

import (
	"estudosgolang/model"
	"fmt"
	"time"
)

func main() {
	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "arroz")
	nomeDosItens = append(nomeDosItens, "feijao")
	nomeDosItens = append(nomeDosItens, "mandioca")
	nomeDosItens = append(nomeDosItens, "batata")
	nomeDosItens = append(nomeDosItens, "frango")

	compra, err := model.ComprasDoMes(nomeDosItens, "Carrefour", time.Now())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compra)
	}
}
