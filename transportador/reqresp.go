package transportador

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	CriarEntregaRequest struct {
		Entrega Entrega
	}
	CriarEntregaResponse struct {
		Ok string `json:"ok"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeEntregaReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CriarEntregaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
