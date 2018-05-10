package query

import (
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/model"
)

//bhinnekanerQueryInMemory implementation
type bhinnekanerQueryInMemory struct {
	db map[string]*model.Bhinnekaner
}

//NewBhinnekanerQueryInMemory
func NewBhinnekanerQueryInMemory(db map[string]*model.Bhinnekaner) *bhinnekanerQueryInMemory {
	return &bhinnekanerQueryInMemory{db: db}
}

//FindAll
func (q *bhinnekanerQueryInMemory) FindAll() <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		defer close(output)

		var bhinnekaners model.Bhinnekaners
		for _, v := range q.db {
			bhinnekaners = append(bhinnekaners, *v)
		}

		output <- QueryResult{Result: bhinnekaners}
	}()
	return output
}
