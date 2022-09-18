package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"

	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
	"google.golang.org/grpc"
)

const host = "localhost"

var (
	port           string
	responsePeriod int64
	rangeTime      int64
)

func init() {
	flag.StringVar(&port, "port", "7002", "daemon port")
	flag.Int64Var(&responsePeriod, "n", 5, "period for sending statistics (sec)")
	flag.Int64Var(&rangeTime, "m", 15, "the range for which the average statistics are collected (sec)")
}

func main() {
	flag.Parse()

	// nolint:staticcheck
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to dial connection" + err.Error())
		return
	}
	defer conn.Close()

	client := desc.NewStatsServiceV1Client(conn)

	req := &desc.StartMonitoringRequest{
		ResponsePeriod: responsePeriod,
		RangeTime:      rangeTime,
	}

	stream, err := client.StartMonitoring(context.Background(), req)
	if err != nil {
		fmt.Println("failed to start monitoring")
		return
	}

	for {
		data, errRecv := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("statistics collection completed")
			break
		}
		if errRecv != nil {
			log.Fatalf(" failed to get stats, err = %v", err)
		}

		printStats(data)
	}
}

func printStats(data *desc.StartMonitoringResponse) {
	fmt.Println("\nCPU:")
	fmt.Println("\tuser mode time:", data.GetCpuInfo().GetUserModeTime())
	fmt.Println("\tsystem mode time:", data.GetCpuInfo().GetSystemModeTime())
	fmt.Println("\tidle time:", data.GetCpuInfo().GetIdleTime())

	fmt.Println("Disk:")
	fmt.Println("\tkbt:", data.GetDiskInfo().GetKbt())
	fmt.Println("\ttps:", data.GetDiskInfo().GetTps())
	fmt.Println("\tmbs:", data.GetDiskInfo().GetMbs())

	fmt.Println("Avg load:")
	fmt.Println("\tavg load 1min:", data.GetLoadInfo().GetLoad1Min())
	fmt.Println("\tavg load 5min:", data.GetLoadInfo().GetLoad5Min())
	fmt.Println("\tavg load 15min:", data.GetLoadInfo().GetLoad15Min())

	fmt.Println("collected at:", data.GetCollectedAt().AsTime())
}
