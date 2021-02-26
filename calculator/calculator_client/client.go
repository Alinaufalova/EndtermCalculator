package main

import (
	"com.grpc.alina/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewGreetServiceClient(conn)
	doGreetingEveryone(c)



}


func primeNumberDServer(c greetpb.GreetServiceClient) {
	ctx := context.Background()
	req := &calculatorpb.primeNumberDRequest{Calculator: &calculatorpb.Calculator{
		Number: "123";
	}}

	stream, err := c.primeNumberD (ctx, req)
	if err != nil {
		log.Fatalf("error while calling primeNumberD RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("error while reciving from primeNumberD RPC %v", err)
		}
		log.Printf("response from primeNumberD:%v \n", res.GetResult())
	}

}