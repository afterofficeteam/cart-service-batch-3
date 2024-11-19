package main

import (
	"cart-service/config"
	"cart-service/repository/cart"
	"cart-service/transport/procedures"
	"cart-service/transport/routes"
	"cart-service/util/middleware"
	"database/sql"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	cartHandler "cart-service/handlers/cart"
	cartSvc "cart-service/usecases/cart"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConn, err := config.ConnectToDatabase(config.Connection{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	listener, err := net.Listen("tcp", "localhost:"+cfg.GrpcPort)
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryLoggingInterceptor),
	)

	var wg sync.WaitGroup
	wg.Add(2)

	procedures, routes := setupTransport(dbConn, grpcServer, listener)
	go routes.Run(cfg.HttpPort, &wg)
	go procedures.RunRpcServer(cfg.GrpcPort, &wg)

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	<-stopSignal
	log.Println("Shutting down servers...")

	grpcServer.GracefulStop()
	log.Println("gRPC server stopped")

	routes.ShutdownHTTP()

	wg.Done()
	log.Println("All servers stopped successfully")
}

func setupTransport(db *sql.DB, server *grpc.Server, listener net.Listener) (*procedures.Procedures, *routes.Routes) {
	cartRepository := cart.NewStore(db)
	cartSvc := cartSvc.NewSvc(cartRepository)
	cartHandler := cartHandler.NewHandler(cartSvc)

	procedures := &procedures.Procedures{
		Listen: listener,
		Grpc:   server,
		Cart:   cartHandler,
	}

	routes := &routes.Routes{
		Cart: cartHandler,
	}

	return procedures, routes
}
