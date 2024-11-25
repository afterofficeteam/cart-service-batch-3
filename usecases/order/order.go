package order

import (
	model "cart-service/models"
	"cart-service/repository/order"

	"github.com/google/uuid"
)

type svc struct {
	repository order.Repository
}

func NewSvc(repo order.Repository) *svc {
	return &svc{
		repository: repo,
	}
}

type Service interface {
	CreateOrder(bReq model.Order) (*uuid.UUID, error)
	UpdatePayment(bReq model.UpdateRequest) (*string, error)
}

func (s *svc) CreateOrder(bReq model.Order) (*uuid.UUID, error) {
	orderID, refCode, err := s.repository.CreateOrder(bReq)
	if err != nil {
		println("error insert data")
		return nil, err
	}

	_, err = s.repository.CreateOrderItemsLogs(model.OrderItemsLogs{
		OrderID:    *orderID,
		RefCode:    *refCode,
		FromStatus: "",
		ToStatus:   model.OrderStatusPending,
		Notes:      "Order created",
	})
	if err != nil {
		println("error insert logs")
		return nil, err
	}

	return orderID, nil
}

func (s *svc) UpdatePayment(bReq model.UpdateRequest) (*string, error) {
	refCode, err := s.repository.UpdateOrder(bReq)
	if err != nil {
		return nil, err
	}

	_, err = s.repository.CreateOrderItemsLogs(model.OrderItemsLogs{
		OrderID:    bReq.OrderID,
		RefCode:    *refCode,
		FromStatus: model.OrderStatusPending,
		ToStatus:   model.OrderStatusPaid,
		Notes:      "Payment success",
	})
	if err != nil {
		return nil, err
	}

	updateOK := "Payment Success"
	return &updateOK, nil
}
