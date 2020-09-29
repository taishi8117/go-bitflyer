package positions

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/taishi8117/go-bitflyer/pkg/api/v1/sendchildorder"
)

type Request struct {
	ProductCode markets.ProductCode `json:"product_code" url:"product_code"`
}

type Response []Position

type Position struct {
	ProductCode         markets.ProductCode `json:"product_code"`
	Side                sendchildorder.Side `json:"side"`
	Price               float64             `json:"price"`
	Size                float64             `json:"size"`
	Commission          float64             `json:"commission"`
	SwapPointAccumulate float64             `json:"swap_point_accumulate"`
	RequireCollateral   float64             `json:"require_collateral"`
	OpenDate            time.Time           `json:"open_date"`
	Leverage            float64             `json:"leverage"`
	Pnl                 float64             `json:"pnl"`
	Sfd                 float64             `json:"sfd"`
}

const (
	APIPath = "/v1/me/getpositions"
)

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
