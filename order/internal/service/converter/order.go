package converter

import (
	repomodel "github.com/stas-tsarev/dz1_order_service/order/internal/repository/model"
	"github.com/stas-tsarev/dz1_order_service/order/internal/service/model"
)

func RepoOrderToServiceOrder(order repomodel.Order) model.Order {
	return model.Order{
		OrderUuid:       order.OrderUuid,
		UserUuid:        order.UserUuid,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUuid: order.TransactionUuid,
		PaymentMethod:   order.PaymentMethod,
		Status:          order.Status,
	}
}
