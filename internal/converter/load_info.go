package converter

import (
	"github.com/olezhek28/system-stats-daemon/internal/model"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
)

// ToDescLoadInfo ...
func ToDescLoadInfo(loadInfo *model.LoadInfo) *desc.LoadInfo {
	return &desc.LoadInfo{
		Load1Min:  loadInfo.Load1Min,
		Load5Min:  loadInfo.Load5Min,
		Load15Min: loadInfo.Load15Min,
	}
}
