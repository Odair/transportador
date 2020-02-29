package transportador

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/google/uuid"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s service) CriarEntrega(ctx context.Context, entregaParam Entrega) (string, error) {
	logger := log.With(s.logger, "method", "CriarEntrega")

	entrega := Entrega{
		PedidoID:            entregaParam.PedidoID,
		DataParaBusca:       entregaParam.DataParaBusca,
		PrevisaoParaEntrega: entregaParam.PrevisaoParaEntrega,
		EnderecoOrigem:      entregaParam.EnderecoOrigem,
		EnderecoDestino:     entregaParam.EnderecoDestino,
	}

	entrega.NumeroEntrega = guuid.New()

	if err := s.repostory.CriarEntrega(ctx, entrega); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("criou entrega")

	voucher := Voucher{NumeroEntrega: entrega.NumeroEntrega}

	return "Success", nil
}
