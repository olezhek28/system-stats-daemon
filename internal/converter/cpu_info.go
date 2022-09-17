package converter

import (
	"github.com/olezhek28/system-stats-daemon/internal/model"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
)

// ToDescCPUInfo ...
func ToDescCPUInfo(cpuInfo *model.CPUInfo) *desc.CPUInfo {
	return &desc.CPUInfo{
		UserModeTime:   cpuInfo.UserModeTime,
		SystemModeTime: cpuInfo.SystemModeTime,
		IdleTime:       cpuInfo.IdleTime,
	}
}
