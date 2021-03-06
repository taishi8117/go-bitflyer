package cancelallchildorders

import (
	"encoding/json"
	"net/http"

	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/pkg/errors"
)

type Request struct {
	ProductCode markets.ProductCode `json:"product_code"`
}

type Response struct{}

const (
	APIPath = "/v1/me/cancelallchildorders"
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
