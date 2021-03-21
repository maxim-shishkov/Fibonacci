package apiGRPC

import (
	f "Fibonacci/src"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct{}

func (c *GRPCServer) Get(cnx context.Context, req *FibonacciRequest) (*FibonacciResponse, error) {
	x := fmt.Sprint( req.GetX() )
	y := fmt.Sprint( req.GetY() )

	m,err := f.GetFibonacci(x, y)

	if err != nil {
		log.Println(err)
		return nil, err
	}

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
