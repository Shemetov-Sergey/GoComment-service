package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Shemetov-Sergey/GoComment-service/pkg/config"
	"github.com/Shemetov-Sergey/GoComment-service/pkg/db"
	"github.com/Shemetov-Sergey/GoComment-service/pkg/middleware"
	"github.com/Shemetov-Sergey/GoComment-service/pkg/pb"
	"github.com/Shemetov-Sergey/GoComment-service/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", ":"+c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.LoggingInterceptor,
		),
	)

	pb.RegisterCommentServiceServer(grpcServer, &s)
	pb.RegisterCensorServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
