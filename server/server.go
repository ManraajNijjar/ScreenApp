package main

import (
	"context"
	"log"
	"net"

	"../../Screen/screenpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sample(ctx context.Context, req *screenpb.SampleRequest) (*screenpb.SampleResponse, error) {
	responseValue := req.GetCalculation().GetFirstInt() + req.GetCalculation().GetSecondInt()

	res := &screenpb.SampleResponse{
		Result: responseValue,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	screenpb.RegisterScreenServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
