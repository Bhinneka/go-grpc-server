package main

import (
	"fmt"
	"os"

	"github.com/Bhinneka/go-grpc-server/internal/app/datasource"
	bhinnekanerGrpc "github.com/Bhinneka/go-grpc-server/internal/app/grpc/server"
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/presenter"
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/query"
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/repository"
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/usecase"
)

//GrpcPortDefault
const GrpcPortDefault = 3000

//db Global In Memory Database
var db = datasource.GetBhinnekanerInMemory()

func main() {

	bhinnekanerRepository := repository.NewBhinnekanerRepositoryInMemory(db)
	bhinnekanerQuery := query.NewBhinnekanerQueryInMemory(db)

	bhinnekanerUsecase := usecase.NewBhinnekanerUsecaseInMemory(bhinnekanerQuery, bhinnekanerRepository)

	bhinnekanerGrpcHandler := presenter.NewGrpcHandler(bhinnekanerUsecase)

	grpcServer, err := bhinnekanerGrpc.NewGrpcServer(bhinnekanerGrpcHandler)

	if err != nil {
		fmt.Printf("Error create grpc server: %s", err.Error())
		os.Exit(1)
	}

	err = grpcServer.Serve(GrpcPortDefault)

	if err != nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}

}
