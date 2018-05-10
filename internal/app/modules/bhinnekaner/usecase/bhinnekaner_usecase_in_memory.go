package usecase

import (
	"errors"
	
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/model"
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/query"
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/repository"
)

//bhinnekanerUsecaseInMemory
type bhinnekanerUsecaseInMemory struct {
	bhinnekanerQuery      query.BhinnekanerQuery
	bhinnekanerRepository repository.BhinnekanerRepository
}

//NewBhinnekanerUsecaseInMemory
func NewBhinnekanerUsecaseInMemory(bhinnekanerQuery query.BhinnekanerQuery, bhinnekanerRepository repository.BhinnekanerRepository) *bhinnekanerUsecaseInMemory {
	return &bhinnekanerUsecaseInMemory{bhinnekanerQuery: bhinnekanerQuery, bhinnekanerRepository: bhinnekanerRepository}
}

//Save
func (u *bhinnekanerUsecaseInMemory) Save(bhinnekaner *model.Bhinnekaner) <-chan UsecaseResult {
	output := make(chan UsecaseResult)

	go func() {
		bhinnekanerResult := <-u.bhinnekanerRepository.Save(bhinnekaner)

		if bhinnekanerResult.Error != nil {
			output <- UsecaseResult{Error: bhinnekanerResult.Error}
			return
		}

		b, ok := bhinnekanerResult.Result.(*model.Bhinnekaner)

		if !ok {
			output <- UsecaseResult{Error: errors.New("Result is not Bhinnekaner")}
			return
		}

		output <- UsecaseResult{Result: b}
	}()

	return output
}

//GetByID
func (u *bhinnekanerUsecaseInMemory) GetByID(id string) <-chan UsecaseResult {
	output := make(chan UsecaseResult)

	go func() {
		bhinnekanerResult := <-u.bhinnekanerRepository.FindByID(id)

		if bhinnekanerResult.Error != nil {
			output <- UsecaseResult{Error: bhinnekanerResult.Error}
			return
		}

		bhinnekaner, ok := bhinnekanerResult.Result.(*model.Bhinnekaner)

		if !ok {
			output <- UsecaseResult{Error: errors.New("Result is not Bhinnekaner")}
			return
		}

		output <- UsecaseResult{Result: bhinnekaner}
	}()

	return output
}

//GetAll
func (u *bhinnekanerUsecaseInMemory) GetAll() <-chan UsecaseResult {
	output := make(chan UsecaseResult)

	go func() {
		bhinnekanerResult := <-u.bhinnekanerQuery.FindAll()

		if bhinnekanerResult.Error != nil {
			output <- UsecaseResult{Error: bhinnekanerResult.Error}
			return
		}

		bhinnekaners, ok := bhinnekanerResult.Result.(model.Bhinnekaners)

		if !ok {
			output <- UsecaseResult{Error: errors.New("Result is not Products")}
			return
		}

		output <- UsecaseResult{Result: bhinnekaners}
	}()

	return output
}
