package main

import (
    "log"
    "net"

    "google.golang.org/grpc"

    menusService "example.com/myproject/internal/menus"
    ordersService "example.com/myproject/internal/orders"
    menuspb "example.com/myproject/proto/menus"
    orderspb "example.com/myproject/proto/orders"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()

    menuspb.RegisterMenuServiceServer(grpcServer, menusService.NewMenuService())
    orderspb.RegisterOrderServiceServer(grpcServer, ordersService.NewOrderService())

    log.Println("Server is running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
