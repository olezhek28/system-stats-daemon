//go:build linux
// +build linux

package load

import (
	"github.com/olezhek28/system-stats-daemon/internal/model"
)

const (
	load1MinPos  = 0
	load5MinPos  = 1
	load15MinPos = 2
)

// GetStats ...
func GetStats() (*model.CPUInfo, error) {
	cmd := exec.Command("cat", "/proc/loadavg")
	res, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to exec %s: %w", cmd.String(), err)
	}

	fields := strings.Fields(string(res))

	userModeTime, err := strconv.ParseFloat(fields[load1MinPos], 64)
	if err != nil {
		return nil, err
	}

	systemModeTime, err := strconv.ParseFloat(fields[load5MinPos], 64)
	if err != nil {
		return nil, err
	}

	idleTime, err := strconv.ParseFloat(fields[load15MinPos], 64)
	if err != nil {
		return nil, err
	}

	return &model.LoadInfo{
		Load1Min:  userModeTime,
		Load5Min:  systemModeTime,
		Load15Min: idleTime,
	}, nil
}
