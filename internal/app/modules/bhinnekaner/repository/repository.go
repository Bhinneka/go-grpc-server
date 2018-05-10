package repository

import (
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/model"
)

//RepositoryResult (Generic Result for repository)
type RepositoryResult struct {
	Result interface{}
	Error  error
}

//BhinnekanerRepository (Generic repository for Bhinnekaner)
type BhinnekanerRepository interface {
	Save(*model.Bhinnekaner) <-chan RepositoryResult
	FindByID(string) <-chan RepositoryResult
}
