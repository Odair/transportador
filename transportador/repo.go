package transportador

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CriarEntrega(ctx context.Context, entrega Entrega) error {
	sql := `
		INSERT INTO entregas ( PedidoId, DataParaBusca, PrevisaoParaEntrega, EnderecoOrigem, EnderecoDestino )
		VALUES ($1, $2, $3, $4, $5)`

	_, err := repo.db.ExecContext(ctx, sql, entrega.PedidoID, entrega.DataParaBusca, entrega.PrevisaoParaEntrega, entrega.EnderecoOrigem, entrega.EnderecoDestino)
	if err != nil {
		return err
	}
	return nil
}
