//go:build darwin
// +build darwin

package disk

import (
	"strconv"

	"github.com/olezhek28/system-stats-daemon/internal/helper"
	"github.com/olezhek28/system-stats-daemon/internal/model"
)

const (
	kbtPos = 13
	tpsPos = 14
	mbsPos = 15
)

// GetStats ...
func GetStats() (*model.DiskInfo, error) {
	fields, err := helper.GetLoadStats()
	if err != nil {
		return nil, err
	}

	kbt, err := strconv.ParseFloat(fields[kbtPos], 64)
	if err != nil {
		return nil, err
	}

	tps, err := strconv.ParseFloat(fields[tpsPos], 64)
	if err != nil {
		return nil, err
	}

	mbs, err := strconv.ParseFloat(fields[mbsPos], 64)
	if err != nil {
		return nil, err
	}

	return &model.DiskInfo{
		Kbt: kbt,
		Tps: tps,
		Mbs: mbs,
	}, nil
}
