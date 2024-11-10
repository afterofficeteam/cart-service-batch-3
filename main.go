package main

import (
	"cart-service/config"
	"cart-service/procedures"
	"cart-service/repository/cart"
	"cart-service/util/middleware"
	"database/sql"
	"net"

	cartHandler "cart-service/handlers/cart"
	cartSvc "cart-service/usecases/cart"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	dbConn, err := config.ConnectToDatabase(config.Connection{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		return
	}
	defer dbConn.Close()

	listener, err := config.NetworkListener("tcp", "localhost:"+cfg.AppPort)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryLoggingInterceptor),
	)

	procedures := setupProcedures(dbConn, grpcServer, listener)
	procedures.RunRpcServer(cfg.AppPort)
}

func setupProcedures(db *sql.DB, server *grpc.Server, listener net.Listener) *procedures.Procedures {
	cartRepository := cart.NewStore(db)
	cartSvc := cartSvc.NewSvc(cartRepository)
	cartHandler := cartHandler.NewHandler(cartSvc)

	return &procedures.Procedures{
		Listen: listener,
		Grpc:   server,
		Cart:   cartHandler,
	}
}
