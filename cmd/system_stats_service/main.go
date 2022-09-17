package main

import (
	"log"
	"net"

	"github.com/olezhek28/system-stats-daemon/internal/app/api/stats_service_v1"
	"github.com/olezhek28/system-stats-daemon/internal/service/stats"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
	"google.golang.org/grpc"
)

const port = ":7002"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	desc.RegisterStatsServiceV1Server(grpcServer, stats_service_v1.NewStatsServiceV1(stats.NewStatsService()))
	grpcServer.Serve(lis)
}
