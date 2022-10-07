//go:build darwin
// +build darwin

package cpu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStat(t *testing.T) {
	t.Run("test success get stats", func(t *testing.T) {
		cpuInfo, err := GetStats()

		require.NoError(t, err)
		require.NotNil(t, cpuInfo.UserModeTime)
		require.IsType(t, int64(1), cpuInfo.UserModeTime)
		require.NotNil(t, cpuInfo.SystemModeTime)
		require.IsType(t, int64(1), cpuInfo.SystemModeTime)
		require.NotNil(t, cpuInfo.IdleTime)
		require.IsType(t, int64(1), cpuInfo.IdleTime)
	})
}
