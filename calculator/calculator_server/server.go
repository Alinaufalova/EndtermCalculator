package main

import (
	"com.grpc.alina/calculator/calculatorpb"
	"com.grpc.alina/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	calculatorpb.UnimplementedGreetServiceServer
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}

func (s *Server) PrimeNumberD(req *calculatorpb.PrimeNumberDRequest, stream calculatorpb.PrimeNumberD_PrimeNumberDServer) error {
	fmt.Printf("PrimeNumberD function was invoked with %v \n", req)
	firstName := req.Get().GetFirstName()
	for i := 0; i < 10; i++ {
		res := &calculatorpb.PrimeNumberDResponse{Result: fmt.Sprintf("%d) Hello, %v\n", i, firstName)}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending Prime Number Decomposition responses: %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}
