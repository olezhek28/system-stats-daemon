//go:build darwin
// +build darwin

package cpu

import (
	"strconv"

	"github.com/olezhek28/system-stats-daemon/internal/helper"
	"github.com/olezhek28/system-stats-daemon/internal/model"
)

const (
	userModeTimePos   = 16
	systemModeTimePos = 17
	idleModeTimePos   = 18
)

// GetStats ...
func GetStats() (*model.CPUInfo, error) {
	fields, err := helper.GetLoadStats()
	if err != nil {
		return nil, err
	}

	userModeTime, err := strconv.ParseInt(fields[userModeTimePos], 10, 64)
	if err != nil {
		return nil, err
	}

	systemModeTime, err := strconv.ParseInt(fields[systemModeTimePos], 10, 64)
	if err != nil {
		return nil, err
	}

	idleTime, err := strconv.ParseInt(fields[idleModeTimePos], 10, 64)
	if err != nil {
		return nil, err
	}

	return &model.CPUInfo{
		UserModeTime:   userModeTime,
		SystemModeTime: systemModeTime,
		IdleTime:       idleTime,
	}, nil
}
