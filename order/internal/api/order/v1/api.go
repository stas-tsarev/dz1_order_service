package v1

import (
	"github.com/stas-tsarev/dz1_order_service/order/internal/service"
	order_v1 "github.com/stas-tsarev/dz1_order_service/shared/pkg/openapi/order/v1"
)

type api struct {
	order_v1.UnimplementedHandler
	OrderService service.OrderService
}

func NewApi(OrderService service.OrderService) *api {
	return &api{
		OrderService: OrderService,
	}
}
