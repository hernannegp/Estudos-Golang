package model

import (
	"errors"
	"time"
)

type Compras struct {
	Itens        []string
	Mercado      string
	DataDaCompra time.Time
}

func ComprasDoMes(produtos []string, mercado string, data time.Time) (*Compras, error) {
	if mercado == "" {
		return nil, errors.New("Mercado é obrigatorio")
	}
	if len(produtos) == 0 {
		return nil, errors.New("Lista vazia")
	}
	for _, value := range produtos {
		if value == "" {
			return nil, errors.New("Lista esta com alguns itens vazios")
		}
	}
	return &Compras{
		Itens:        produtos,
		Mercado:      mercado,
		DataDaCompra: data,
	}, nil
}
