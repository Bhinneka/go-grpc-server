package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"

	pb "github.com/Bhinneka/go-grpc-server/api/protogo/bhinnekaner"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	client, err := clientBhinnekaner("localhost:3000")

	if err != nil {
		log.Fatalf("error 1 %v", err)
	}

	GetBhinnekaners(client)
}

func clientBhinnekaner(serverHost string) (pb.BhinnekanerServiceClient, error) {

	conn, err := grpc.Dial(serverHost, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	//defer conn.Close()

	client := pb.NewBhinnekanerServiceClient(conn)

	return client, nil
}

func AddBhinnekaner(client pb.BhinnekanerServiceClient) {
	md := metadata.Pairs("authorization", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	bhinnekanerRequest := &pb.BhinnekanerRequest{ID: "B4", Name: "Push", Email: "push@yahoo.com"}
	resp, err := client.Add(ctx, bhinnekanerRequest)

	if err != nil {
		log.Fatalf("error 1 %v", err)
	}

	fmt.Println(resp)
}

func GetBhinnekaners(client pb.BhinnekanerServiceClient) {
	md := metadata.Pairs("authorization", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	param := &pb.BhinnekanerQuery{}
	resStream, err := client.FindAll(ctx, param)

	if err != nil {
		log.Fatalf("error 1 %v", err)
	}

	for {
		foo, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error 1 %v", err)
		}

		fmt.Println(foo)
	}

}

func GetBhinnekaner(client pb.BhinnekanerServiceClient) {
	md := metadata.Pairs("authorization", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	param := &pb.BhinnekanerQuery{ID: "B3"}
	resStream, err := client.FindByID(ctx, param)

	if err != nil {
		log.Fatalf("error 1 %v", err)
	}

	foo, err := resStream.Recv()

	if err != nil {
		log.Fatalf("error 1 %v", err)
	}

	fmt.Println(foo)
}
