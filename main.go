package main

import (
	"estudosgolang/model"
	"fmt"
	"time"
)

func main() {
	lista := model.CompraDoMes{
		Itens:        []string{"Pão", "Arroz", "Carne"},
		Mercado:      "Carrefour",
		DataDaCompra: time.Now(),
	}
	fmt.Println("Oque Comprou:", lista.Itens, "/", "Mercado:", lista.Mercado, "/", "Data da Compra:", lista.DataDaCompra)
}
