package transportador

import (
	"time"
)

type Voucher struct{

	NumeroEntrega int `json:"numeroEntrega"`

	PrevisaoParaEntrega time.time `json:"previsaoParaEntrega"`
}