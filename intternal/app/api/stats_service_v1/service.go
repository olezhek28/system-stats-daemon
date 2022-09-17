package stats_service_v1

import (
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
)

// Implementation ...
type Implementation struct {
	desc.UnimplementedStatsServiceV1Server
}

// NewStatsServiceV1 return new instance of Implementation.
func NewStatsServiceV1() *Implementation {
	return &Implementation{
		desc.UnimplementedStatsServiceV1Server{},
	}
}

func newMockStatsServiceV1(i Implementation) *Implementation {
	return &Implementation{
		desc.UnimplementedStatsServiceV1Server{},
	}
}
