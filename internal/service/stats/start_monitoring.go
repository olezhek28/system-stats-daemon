package stats

import (
	"github.com/olezhek28/system-stats-daemon/internal/helper/cpu"
	"github.com/olezhek28/system-stats-daemon/internal/helper/disk"
	"github.com/olezhek28/system-stats-daemon/internal/helper/load"
	"github.com/olezhek28/system-stats-daemon/internal/model"
)

// StartMonitoring ...
func (s *Service) StartMonitoring(responsePeriod int64, rangeTime int64) (*model.DeviceInfo, error) {
	cpuInfo, err := cpu.GetStats()
	if err != nil {
		return nil, err
	}

	diskInfo, err := disk.GetStats()
	if err != nil {
		return nil, err
	}

	loadInfo, err := load.GetStats()
	if err != nil {
		return nil, err
	}

	return &model.DeviceInfo{
		CPUInfo:  cpuInfo,
		DiskInfo: diskInfo,
		LoadInfo: loadInfo,
	}, nil
}
