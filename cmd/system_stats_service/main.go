package main

import (
	"flag"
	"log"
	"net"

	"github.com/olezhek28/system-stats-daemon/internal/app/api/stats_service_v1"
	"github.com/olezhek28/system-stats-daemon/internal/config"
	"github.com/olezhek28/system-stats-daemon/internal/service/stats"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
	"google.golang.org/grpc"
)

const host = "localhost"

var (
	port       string
	configPath string
)

func init() {
	flag.StringVar(&port, "port", "7002", "daemon port")
	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config file")
}

func main() {
	flag.Parse()

	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("failed to get config: %s", err)
	}

	lis, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	desc.RegisterStatsServiceV1Server(grpcServer, stats_service_v1.NewStatsServiceV1(stats.NewStatsService(cfg)))

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
