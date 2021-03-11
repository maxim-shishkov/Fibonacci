package apiGRPC

import (
	f "Fibonacci/src"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct{}

func (c *GRPCServer) Get(cnx context.Context, req *FibonacciRequest) (*FibonacciResponse, error) {
	x := int(req.GetX())
	y := int(req.GetY())

	m := f.GetFibonacci(x, y)
	return &FibonacciResponse{Result: m}, nil
}

func StartServer() {
	s := grpc.NewServer()
	srv := &GRPCServer{}
	RegisterGetFibonacciServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Start Server")

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
