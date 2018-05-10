package usecase

import (
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/model"
)

//UsecaseResult (Generic result for usecase)
type UsecaseResult struct {
	Result interface{}
	Error  error
}

//BhinnekanerUsecase
type BhinnekanerUsecase interface {
	Save(*model.Bhinnekaner) <-chan UsecaseResult
	GetByID(string) <-chan UsecaseResult
	GetAll() <-chan UsecaseResult
}
