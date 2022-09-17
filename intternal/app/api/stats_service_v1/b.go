package stats_service_v1

import (
	"strconv"
	"time"

	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
)

func (i *Implementation) StartMonitoring(req *desc.StartMonitoringRequest, stream desc.StatsServiceV1_StartMonitoringServer) error {
	k := int64(0)

	for {
		stream.Send(&desc.StartMonitoringResponse{Title: strconv.FormatInt(k, 10)})
		k++

		time.Sleep(1 * time.Second)

		if k == 20 {
			break
		}
	}

	return nil
}
