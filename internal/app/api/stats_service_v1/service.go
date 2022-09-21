package stats_service_v1

import (
	"github.com/olezhek28/system-stats-daemon/internal/service/stats"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
)

// Implementation ...
type Implementation struct {
	desc.UnimplementedStatsServiceV1Server

	statsService *stats.Service
}

// NewStatsServiceV1 return new instance of Implementation.
func NewStatsServiceV1(statsService *stats.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedStatsServiceV1Server{},

		statsService,
	}
}
