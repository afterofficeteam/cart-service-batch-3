package cart

import (
	"cart-service/proto/cart"
	cartSvc "cart-service/usecases/cart"
	"cart-service/util/helpers"
	"cart-service/util/middleware"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

func (h *Handler) DetailCart(ctx context.Context, bReq *cart.CartDetailRequest) (*cart.CartDetailResponse, error) {
	return h.svc.GetDetails(bReq)
}

func (h *Handler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.PathValue("user_id"))
	if err != nil {
		helpers.HandleResponse(w, http.StatusBadRequest, "invalid user ID format")
		return
	}

	if limiter := middleware.GetLimiter(fmt.Sprintf("%v", userID)); !limiter.Allow() {
		helpers.HandleResponse(w, http.StatusTooManyRequests, "To many request, please try again later")
		return
	}

	productID, err := uuid.Parse(r.PathValue("product_id"))
	if err != nil {
		helpers.HandleResponse(w, http.StatusBadRequest, "invalid product ID format")
		return
	}

	if userID == uuid.Nil || productID == uuid.Nil {
		helpers.HandleResponse(w, http.StatusBadRequest, "user ID or product ID cannot be empty")
		return
	}

	updateOK, err := h.svc.Delete(userID, productID)
	if err != nil {
		helpers.HandleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.HandleResponse(w, http.StatusOK, updateOK)
}
