package main

import (
	"context"
	"fmt"
	"io"
	"log"

	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:7002", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to dial connection" + err.Error())
		return
	}
	defer conn.Close()

	client := desc.NewStatsServiceV1Client(conn)

	req := &desc.StartMonitoringRequest{
		ResponsePeriod: 5,
		RangeTime:      15,
	}

	stream, err := client.StartMonitoring(context.Background(), req)
	if err != nil {
		fmt.Println("failed to start monitoring")
		return
	}

	for {
		data, errRecv := stream.Recv()
		if errRecv == io.EOF {
			fmt.Println("finish")
			break
		}
		if errRecv != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}

		log.Println(data)
	}
}
