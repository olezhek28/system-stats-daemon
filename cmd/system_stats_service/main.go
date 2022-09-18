package main

import (
	"log"
	"net"

	"github.com/olezhek28/system-stats-daemon/internal/app/api/stats_service_v1"
	"github.com/olezhek28/system-stats-daemon/internal/service/stats"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
	"google.golang.org/grpc"
)

const address = "localhost:7002"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	desc.RegisterStatsServiceV1Server(grpcServer, stats_service_v1.NewStatsServiceV1(stats.NewStatsService()))

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
