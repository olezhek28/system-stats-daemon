package stats

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/olezhek28/system-stats-daemon/internal/converter"
	"github.com/olezhek28/system-stats-daemon/internal/helper/cpu"
	"github.com/olezhek28/system-stats-daemon/internal/helper/disk"
	"github.com/olezhek28/system-stats-daemon/internal/helper/load"
	"github.com/olezhek28/system-stats-daemon/internal/model"
	desc "github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const decimals = 2

type data struct {
	statData []*model.DeviceInfo
	m        sync.Mutex
}

// StartMonitoring ...
func (s *Service) StartMonitoring(stream desc.StatsServiceV1_StartMonitoringServer, responsePeriod int64, rangeTime int64) error {
	collectTicker := time.NewTicker(1 * time.Second)
	statsInfo := &data{}

	go func() {
		for {
			select {
			case <-collectTicker.C:
				res, err := getStats()
				if err != nil {
					log.Println(stream.Context(), "failed to get stats: %s", err)
					continue
				}

				statsInfo.m.Lock()
				statsInfo.statData = append(statsInfo.statData, res)
				statsInfo.m.Unlock()

			case <-stream.Context().Done():
				log.Println("data collection interrupted")
				return
			}
		}
	}()

	responseTicker := time.NewTicker(time.Duration(responsePeriod) * time.Second)
	for {
		select {
		case <-responseTicker.C:
			if len(statsInfo.statData) >= int(rangeTime) {
				statsInfo.m.Lock()
				res := statsInfo.statData[:rangeTime]
				statsInfo.statData = statsInfo.statData[responsePeriod:]
				statsInfo.m.Unlock()

				avgStats := calcAvg(res)

				err := stream.Send(&desc.StartMonitoringResponse{
					CpuInfo:     converter.ToDescCPUInfo(avgStats.CPUInfo),
					DiskInfo:    converter.ToDescDiskInfo(avgStats.DiskInfo),
					LoadInfo:    converter.ToDescLoadInfo(avgStats.LoadInfo),
					CollectedAt: timestamppb.New(time.Now()),
				})
				if err != nil {
					return err
				}
			}

		case <-stream.Context().Done():
			log.Println("sending data interrupted")
			return nil
		}
	}
}

func getStats() (*model.DeviceInfo, error) {
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

func calcAvg(stats []*model.DeviceInfo) *model.DeviceInfo {
	var sumUserModeTime, sumSystemModeTime, sumIdleTime int64
	var sumKbt, sumTps, sumMbs float64
	var sumLoad1Min, sumLoad5Min, sumLoad15Min float64

	for _, stat := range stats {
		sumUserModeTime += stat.CPUInfo.UserModeTime
		sumSystemModeTime += stat.CPUInfo.SystemModeTime
		sumIdleTime += stat.CPUInfo.IdleTime

		sumKbt += stat.DiskInfo.Kbt
		sumTps += stat.DiskInfo.Tps
		sumMbs += stat.DiskInfo.Mbs

		sumLoad1Min += stat.LoadInfo.Load1Min
		sumLoad5Min += stat.LoadInfo.Load5Min
		sumLoad15Min += stat.LoadInfo.Load15Min
	}

	return &model.DeviceInfo{
		CPUInfo: &model.CPUInfo{
			UserModeTime:   sumUserModeTime / int64(len(stats)),
			SystemModeTime: sumSystemModeTime / int64(len(stats)),
			IdleTime:       sumIdleTime / int64(len(stats)),
		},
		DiskInfo: &model.DiskInfo{
			Kbt: round(sumKbt/float64(len(stats)), decimals),
			Tps: round(sumTps/float64(len(stats)), decimals),
			Mbs: round(sumMbs/float64(len(stats)), decimals),
		},
		LoadInfo: &model.LoadInfo{
			Load1Min:  round(sumLoad1Min/float64(len(stats)), decimals),
			Load5Min:  round(sumLoad5Min/float64(len(stats)), decimals),
			Load15Min: round(sumLoad15Min/float64(len(stats)), decimals),
		},
	}
}

func round(num float64, decimals int) float64 {
	return math.Round(num*float64(decimals)) / float64(decimals)
}
