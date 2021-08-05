package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"grpc-app/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	res := &proto.AddResponse{Result: result}
	return res, nil
}

func (server *server) Average(averageStream proto.AppService_AverageServer) error {
	total := int64(0)
	count := int64(0)
	for {
		avgReq, err := averageStream.Recv()
		if err == io.EOF {
			average := total / count
			log.Printf("Average(%d) = %d", count, average)
			averageResponse := &proto.AverageResponse{Result: average}
			averageStream.SendAndClose(averageResponse)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		total += avgReq.GetNo()
		count++
	}
	return nil
}

func (s *server) GeneratePrime(req *proto.PrimeNumberRequest, resStream proto.AppService_GeneratePrimeServer) error {
	rangeStart := req.GetRangeStart()
	rangeEnd := req.GetRangeEnd()
	log.Printf("GeneratePrime(%d, %d)", rangeStart, rangeEnd)
	for i := rangeStart; i <= rangeEnd; i++ {
		if isPrime(i) {
			fmt.Println("Sending Prime No ", i)
			res := &proto.PrimeNumberResponse{PrimeNumber: i}
			resStream.Send(res)
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(n int64) bool {
	for i := int64(2); i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//bidirectional streaming
func (s *server) GreetEveryone(stream proto.AppService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln(err)
		}
		greetingData := req.GetGreeting()
		fmt.Println("Received ", greetingData)
		greetMsg := fmt.Sprintf("Hi %s %s!", greetingData.GetFirstName(), greetingData.GetLastName())
		greetRes := &proto.GreetEveryoneResponse{Message: greetMsg}
		stream.Send(greetRes)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
