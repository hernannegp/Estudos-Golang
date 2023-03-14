package model

import (
	"time"
)

type CompraDoMes struct {
	Itens        []string
	Mercado      string
	DataDaCompra time.Time
}
