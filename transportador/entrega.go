package transportador

import (
	"context"
	"time"

	"github.com/google/uuid"

)

type Entrega struct {
	PedidoID int `json:"pedidoId,omitempty"`

	NumeroEntrega uuid.UUID `json:"numeroEntrega"`

	DataParaBusca time.Time `json:"dataParaBusca"`

	PrevisaoParaEntrega time.Time `json:"previsaoParaEntrega"`

	EnderecoOrigem string `json:"enderecoOrigem"`

	EnderecoDestino string `json:"enderecoDestino"`
}

type Repository interface {
	CriarEntrega(ctx context.Context, entrega Entrega) error
}
