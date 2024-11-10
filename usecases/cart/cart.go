package cart

import (
	cartProto "cart-service/proto/cart"
	"cart-service/repository/cart"
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
