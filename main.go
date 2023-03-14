package main

import "fmt"

func main() {
	itens := []string{"Pão", "Arroz", "Carne"}
	compraDoMes(itens, "Carrefour")
}

func compraDoMes(itens []string, mercado string) {
	fmt.Println("Oque comprar:", itens, "Mercado:", mercado)
}
