package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/nikit34/grpc_basic_tutorial/complete_course/pb"
	"github.com/nikit34/grpc_basic_tutorial/complete_course/service"
)


func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()

	log.Printf("Start server on port %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}