package converter

import (
	"github.com/olezhek28/system-stats-daemon/internal/model"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
)

// ToDescDiskInfo ...
func ToDescDiskInfo(diskInfo *model.DiskInfo) *desc.DiskInfo {
	return &desc.DiskInfo{
		Kbt: diskInfo.Kbt,
		Tps: diskInfo.Tps,
		Mbs: diskInfo.Mbs,
	}
}
