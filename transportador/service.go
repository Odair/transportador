package transportador

import "context"

type Service interface {
	CriarEntrega(ctx context.Context, entrega Entrega) (Voucher, error)
}
