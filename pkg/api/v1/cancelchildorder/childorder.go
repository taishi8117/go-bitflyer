package cancelchildorder

import (
	"encoding/json"
	"net/http"

	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/pkg/errors"
)

type Request struct {
	ProductCode            markets.ProductCode `json:"product_code"`
	ChildOrderId           string              `json:"child_order_id,omitempty"`
	ChildOrderAcceptanceId string              `json:"child_order_acceptance_id,omitempty"`
}

type Response struct{}

const (
	APIPath = "/v1/me/cancelchildorder"
)

func (req *Request) Method() string {
	return http.MethodPost
}

func (req *Request) Query() string {
	return ""
}

func (req *Request) Payload() []byte {
	body, err := json.Marshal(*req)
	if err != nil {
		panic(errors.Wrap(err, "json serialization error"))
	}
	return body
}
