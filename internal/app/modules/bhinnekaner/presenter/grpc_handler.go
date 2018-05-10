package presenter

import (

	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/model"
	"github.com/Bhinneka/go-grpc-server/internal/app/modules/bhinnekaner/usecase"

	pb "github.com/Bhinneka/go-grpc-server/api/protogo/bhinnekaner"

	context "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GrpcHandler
type GrpcHandler struct {
	bhinnekanerUsecase usecase.BhinnekanerUsecase
}

//NewGrpcHandler
func NewGrpcHandler(bhinnekanerUsecase usecase.BhinnekanerUsecase) *GrpcHandler {
	return &GrpcHandler{bhinnekanerUsecase}
}

//Add
func (g *GrpcHandler) Add(ctx context.Context, arg *pb.BhinnekanerRequest) (*pb.BhinnekanerResponse, error) {

	var b model.Bhinnekaner
	b.ID = arg.ID
	b.Name = arg.Name
	b.Email = arg.Email

	usecaseResult := <-g.bhinnekanerUsecase.Save(&b)

	if usecaseResult.Error != nil {
		return nil, status.Error(codes.InvalidArgument, usecaseResult.Error.Error())
	}

	bhinnekanerResult, ok := usecaseResult.Result.(*model.Bhinnekaner)

	if !ok {
		return nil, status.Error(codes.Internal, "Result is not Bhinnekaner")
	}

	bhinnekanerResponse := &pb.BhinnekanerResponse{ID: bhinnekanerResult.ID, Name: bhinnekanerResult.Name, Email: bhinnekanerResult.Email}

	return bhinnekanerResponse, nil
}

//FindByID
func (g *GrpcHandler) FindByID(arg *pb.BhinnekanerQuery, stream pb.BhinnekanerService_FindByIDServer) error {

	id := arg.ID
	usecaseResult := <-g.bhinnekanerUsecase.GetByID(id)

	if usecaseResult.Error != nil {
		return status.Error(codes.InvalidArgument, usecaseResult.Error.Error())
	}

	bhinnekanerResult, ok := usecaseResult.Result.(*model.Bhinnekaner)

	if !ok {
		return status.Error(codes.Internal, "Result is not Bhinnekaner")
	}

  bhinnekanerResponse := &pb.BhinnekanerResponse{ID: bhinnekanerResult.ID, Name: bhinnekanerResult.Name, Email: bhinnekanerResult.Email}

	if err := stream.Send(bhinnekanerResponse); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

//FindAll
func (g *GrpcHandler) FindAll(arg *pb.BhinnekanerQuery, stream pb.BhinnekanerService_FindAllServer) error {

	usecaseResult := <-g.bhinnekanerUsecase.GetAll()

	if usecaseResult.Error != nil {
		return status.Error(codes.InvalidArgument, usecaseResult.Error.Error())
	}

	bhinnekaners, ok := usecaseResult.Result.(model.Bhinnekaners)

	if !ok {
		return status.Error(codes.Internal, "Result is not Bhinnekaners")
	}

	for _, v := range bhinnekaners {
		bhinnekanerResponse := &pb.BhinnekanerResponse{ID: v.ID, Name: v.Name, Email: v.Email}

		if err := stream.Send(bhinnekanerResponse); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
