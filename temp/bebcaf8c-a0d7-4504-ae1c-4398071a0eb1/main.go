package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"demoService/internal/prospects"
	"demoService/internal/users"
	 
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func startGRPC(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	
	prospectsServer := prospects.Server{}
	prospects.RegisterProspectsServiceServer(grpcServer, &prospectsServer)
	
	usersServer := users.Server{}
	users.RegisterUsersServiceServer(grpcServer, &usersServer)
	 
	log.Println("gRPC server ready...")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	wg.Done()
}

func startHTTP(wg *sync.WaitGroup) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Register grpc-gateway
	rmux := runtime.NewServeMux()
	
	prospectsClient := prospects.NewProspectsServiceClient(conn)
	err = prospects.RegisterProspectsServiceHandlerClient(ctx, rmux, prospectsClient)
	if err != nil {
		log.Fatal(err)
	}
	
	usersClient := users.NewUsersServiceClient(conn)
	err = users.RegisterUsersServiceHandlerClient(ctx, rmux, usersClient)
	if err != nil {
		log.Fatal(err)
	}
	
	
	handler := cors.Default().Handler(rmux)
	log.Println("rest server ready...")

	err = http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()

}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go startGRPC(&wg)

	wg.Add(1)
	go startHTTP(&wg)

	wg.Wait()
}
