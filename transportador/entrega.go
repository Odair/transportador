package transportador

import (
	"context"
	"time"
)

type Entrega struct {
	PedidoID int `json:"pedidoId,omitempty"`

	DataParaBusca time.Time `json:"dataParaBusca"`

	PrevisaoParaEntrega time.Time `json:"previsaoParaEntrega"`

	EnderecoOrigem string `json:"enderecoOrigem"`

	EnderecoDestino string `json:"enderecoDestino"`
}

type Repository interface {
	CriarEntrega(ctx context.Context, entrega Entrega) error
}
