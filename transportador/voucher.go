package transportador

import (
	"time"

	"github.com/google/uuid"

)

type Voucher struct{

	NumeroEntrega uuid.UUID `json:"numeroEntrega"`

	PrevisaoParaEntrega time.Time `json:"previsaoParaEntrega"`
}