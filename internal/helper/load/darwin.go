//go:build darwin
// +build darwin

package load

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/olezhek28/system-stats-daemon/internal/model"
)

const (
	load1MinPos  = 1
	load5MinPos  = 2
	load15MinPos = 3
)

// GetStats ...
func GetStats() (*model.LoadInfo, error) {
	cmd := exec.Command("sysctl", "-n", "vm.loadavg")
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
