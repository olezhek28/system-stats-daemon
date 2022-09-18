package stats

import (
	"sync"

	"github.com/olezhek28/system-stats-daemon/internal/config"
	"github.com/olezhek28/system-stats-daemon/internal/model"
)

// Service ...
type Service struct {
	statData []*model.DeviceInfo
	m        sync.Mutex
	config   *config.Config
}

// NewStatsService ...
func NewStatsService(config *config.Config) *Service {
	return &Service{
		config: config,
	}
}
