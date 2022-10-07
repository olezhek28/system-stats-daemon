package stats

import (
	"github.com/olezhek28/system-stats-daemon/internal/config"
)

// Service ...
type Service struct {
	config *config.Config
}

// NewStatsService ...
func NewStatsService(config *config.Config) *Service {
	return &Service{
		config: config,
	}
}
