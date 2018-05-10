package datasource

import (
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/model"
)

//GetBhinnekanerInMemory
func GetBhinnekanerInMemory() map[string]*model.Bhinnekaner {
	db := make(map[string]*model.Bhinnekaner)

	db["B1"] = &model.Bhinnekaner{ID: "B1", Name: "Wuriyanto", Email: "wuriyanto@bhinneka.com"}
	db["B2"] = &model.Bhinnekaner{ID: "B2", Name: "fork", Email: "fork@bhinneka.com"}
	db["B3"] = &model.Bhinnekaner{ID: "B3", Name: "commit", Email: "commit@bhinneka.com"}

	return db
}
