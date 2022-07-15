package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rsmarincu/glassnode/pkg/fees/repository"
	"github.com/rsmarincu/glassnode/pkg/fees/usecases"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"

	feespb "github.com/rsmarincu/glassnode/api"
	feesGrpc "github.com/rsmarincu/glassnode/pkg/fees/grpc"

	"github.com/rsmarincu/glassnode/pkg/common"
)

type Opts struct {
	Port        string `short:"p" long:"port"`
	GatewayPort string `short:"g" long:"gateway_port"`
}

var opts Opts

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatalf("failed to parse args: %w", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", opts.Port))
	if err != nil {
		log.Fatal("failed to listen: %w", err)
	}

	connURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
	)
	db, err := common.ConnectPostgress(context.Background(), common.NewDefaultConnectionOptions(connURL))
	if err != nil {
		log.Fatalf("Failed connecting to the database: %w", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	ethRepository := repository.NewETHRepository(db)
	feesService := usecases.NewFeesService(ethRepository)
	feespb.RegisterFeesServer(grpcServer, feesGrpc.NewServiceHandler(feesService))

	log.Printf("Initializing gRPC server on port %d", opts.Port)
	go func() {
		log.Fatalln(grpcServer.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%s", opts.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to dial server: %w", err)
	}

	gatewayMux := runtime.NewServeMux()
	err = feespb.RegisterFeesHandler(context.Background(), gatewayMux, conn)
	if err != nil {
		log.Fatalf("Failed to register gateway: %w", err)
	}

	gatewayServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", opts.GatewayPort),
		Handler: gatewayMux,
	}

	log.Printf("Serving gateway on port %s", opts.GatewayPort)
	log.Fatalln(gatewayServer.ListenAndServe())
}
