package order

import (
	model "cart-service/models"
	orderSvc "cart-service/usecases/order"
	"cart-service/util/helpers"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	svc       orderSvc.Service
	validator *validator.Validate
}

func NewHandler(
	svc orderSvc.Service,
	validator *validator.Validate,
) *Handler {
	return &Handler{
		svc:       svc,
		validator: validator,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var bReq model.Order
	if err := json.NewDecoder(r.Body).Decode(&bReq); err != nil {
		helpers.HandleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bReq.RefCode = helpers.GenerateRefCode()

	if bReq.ProductOrder == nil {
		bReq.ProductOrder = json.RawMessage("[]")
	}

	if err := h.validator.Struct(bReq); err != nil {
		helpers.HandleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bRes, err := h.svc.CreateOrder(bReq)
	if err != nil {
		helpers.HandleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.HandleResponse(w, http.StatusCreated, bRes)
}
