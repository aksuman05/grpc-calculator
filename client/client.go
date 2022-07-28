package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/aksuman05/grpc-calulator/calculatorpb"
	"github.com/aksuman05/grpc-calculator/tree/main/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	c := calculatorpb.NewSumClient(conn)

	// numbers to add
	num := calculatorpb.Numbers{
		A: 10,
		B: 5,
	}

	// call Add service
	res, err := c.Add(context.Background(), &calculatorpb.SumRequest{Numbers: &num})
	if err != nil {
		log.Fatalf("failed to call Add: %v", err)
	}
	fmt.Println(res.Result)
}
