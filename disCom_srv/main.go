package main

import (
	"context"
	pb "disCom_srv/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type server struct {
	pb.UnimplementedCodeServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (s *server) GetAns(ctx context.Context, request *pb.CalRequest) (*pb.CalResponse, error) {
	log.Printf("Received: %v", request.Number)

	var result []*pb.Message
	num, _ := strconv.Atoi(request.Number)

	for i := 1; i < num; i++ {

		if isPrime(i) {
			mess := &pb.Message{
				Time:  uint64(time.Now().Unix()),
				Value: uint32(i),
			}
			result = append(result, mess)
		}
	}
	response := pb.CalResponse{
		Id:   "1",
		Data: result,
	}
	return &response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCodeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
