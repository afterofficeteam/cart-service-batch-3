package cart

import (
	"cart-service/proto/cart"
	cartSvc "cart-service/usecases/cart"
	"context"
)

type Handler struct {
	svc cartSvc.Service
	cart.UnimplementedCartServiceServer
}

func NewHandler(svc cartSvc.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) InsertCart(ctx context.Context, bReq *cart.CartInsertRequest) (*cart.CartInsertResponse, error) {
	return h.svc.Insert(bReq)
}
