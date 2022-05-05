package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/nikit34/grpc_basic_tutorial/complete_course/pb"
	"github.com/nikit34/grpc_basic_tutorial/complete_course/service"
)


func createUser(userStore service.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}

const (
	secretKey = "secret"
	tokenDuration = 15 * time.Minute
)

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/proto.LaptopService/"
	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage": {"admin"},
		laptopServicePath + "RateLaptop": {"admin", "user"},
	}
}

func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()

	log.Printf("Start server on port %d", *port)

	userStore := service.NewInMemoryUserStore()
	err := seedUsers(userStore)
	if err != nil {
		log.Fatal("cannot seed users")
	}

	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()
	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
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