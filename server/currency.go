package server

import (
	"context"

	"github.com/PulkitGuptakugelblitz/currency"
	protos "github.com/PulkitGuptakugelblitz/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
	currency.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency{
	return &Currency{l,protos.UnimplementedCurrencyServer{}}
}
func (c *Currency) GetRate(ctx context.Context, rr*protos.RateRequest)(*protos.RateResponse, error){
	c.log.Info("Handle GetRate", "base", rr.GetBase(),"destination",rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5},nil
}