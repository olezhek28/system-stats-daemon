package stats_service_v1

import (
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
)

// StartMonitoring ...
func (i *Implementation) StartMonitoring(req *desc.StartMonitoringRequest, stream desc.StatsServiceV1_StartMonitoringServer) error {
	return i.statsService.StartMonitoring(stream, req.GetResponsePeriod(), req.GetRangeTime())
}
