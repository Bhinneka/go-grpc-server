package repository

import (
	"errors"
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/model"
	"sync"
)

//bhinnekanerRepositoryInMemory implementation
type bhinnekanerRepositoryInMemory struct {
	db map[string]*model.Bhinnekaner
	sync.RWMutex
}

//NewBhinnekanerRepositoryInMemory
func NewBhinnekanerRepositoryInMemory(db map[string]*model.Bhinnekaner) *bhinnekanerRepositoryInMemory {
	return &bhinnekanerRepositoryInMemory{db: db}
}

//Save
func (r *bhinnekanerRepositoryInMemory) Save(bhinnekaner *model.Bhinnekaner) <-chan RepositoryResult {
	output := make(chan RepositoryResult)
	go func() {
		defer close(output)
		r.Lock()
		defer r.Unlock()

		r.db[bhinnekaner.ID] = bhinnekaner

		output <- RepositoryResult{Result: bhinnekaner}

	}()
	return output
}

//FindByID
func (r *bhinnekanerRepositoryInMemory) FindByID(id string) <-chan RepositoryResult {
	output := make(chan RepositoryResult)
	go func() {
		defer close(output)

		r.RLock()
		defer r.RUnlock()

		bhinnekaner, ok := r.db[id]
		if !ok {
			output <- RepositoryResult{Error: errors.New("bhinnekaner not found")}
			return
		}

		output <- RepositoryResult{Result: bhinnekaner}
	}()
	return output
}
