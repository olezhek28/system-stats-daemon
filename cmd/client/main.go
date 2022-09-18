package main

import (
	"context"
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
		if errRecv == io.EOF {
			fmt.Println("statistics collection completed")
			break
		}
		if errRecv != nil {
			log.Fatalf(" failed to get stats, err = %v", err)
		}

		log.Println(data)
	}
}
