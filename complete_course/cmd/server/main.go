package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/nikit34/grpc_basic_tutorial/complete_course/pb"
	"github.com/nikit34/grpc_basic_tutorial/complete_course/service"
)


func unaryInterceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("--> sream interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

const (
	secretKey = "secret"
	tokenDuration = 15 * time.Minute
)

func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()

	log.Printf("Start server on port %d", *port)

	userStore := service.NewInMemoryUserStore()
	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()

	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	pb.RegisterAuthServiceServer(grpcServer, authServer)
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