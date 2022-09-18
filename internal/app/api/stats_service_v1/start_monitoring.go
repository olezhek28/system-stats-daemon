package stats_service_v1

import (
	"time"

	"github.com/olezhek28/system-stats-daemon/internal/converter"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// StartMonitoring ...
func (i *Implementation) StartMonitoring(req *desc.StartMonitoringRequest, stream desc.StatsServiceV1_StartMonitoringServer) error {
	for {
		res, err := i.statsService.StartMonitoring(req.GetResponsePeriod(), req.GetRangeTime())
		if err != nil {
			return err
		}

		stream.Send(&desc.StartMonitoringResponse{
			CpuInfo:     converter.ToDescCPUInfo(res.CPUInfo),
			DiskInfo:    converter.ToDescDiskInfo(res.DiskInfo),
			LoadInfo:    converter.ToDescLoadInfo(res.LoadInfo),
			CollectedAt: timestamppb.New(time.Now()),
		})

		time.Sleep(2 * time.Second)
	}

	return nil
}
