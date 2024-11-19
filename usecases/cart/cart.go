package cart

import (
	cartProto "cart-service/proto/cart"
	"cart-service/repository/cart"

	"github.com/google/uuid"
)

type svc struct {
	repository cart.Repository
}

func NewSvc(repo cart.Repository) *svc {
	return &svc{
		repository: repo,
	}
}

type Service interface {
	Insert(req *cartProto.CartInsertRequest) (*cartProto.CartInsertResponse, error)
	GetDetails(req *cartProto.CartDetailRequest) (*cartProto.CartDetailResponse, error)
	Delete(userID, productID uuid.UUID) (*int, error)
}

func (s *svc) Insert(req *cartProto.CartInsertRequest) (*cartProto.CartInsertResponse, error) {
	insertOK, err := s.repository.Insert(req)
	if err != nil {
		return nil, err
	}

	return &cartProto.CartInsertResponse{
		Msg: *insertOK,
	}, nil
}

func (s *svc) GetDetails(req *cartProto.CartDetailRequest) (*cartProto.CartDetailResponse, error) {
	return s.repository.GetDetails(req)
}

func (s *svc) Delete(userID, productID uuid.UUID) (*int, error) {
	return s.repository.Delete(userID, productID)
}
