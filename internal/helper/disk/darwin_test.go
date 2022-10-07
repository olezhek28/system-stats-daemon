//go:build darwin
// +build darwin

package disk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStat(t *testing.T) {
	t.Run("test success get stats", func(t *testing.T) {
		diskInfo, err := GetStats()

		require.NoError(t, err)
		require.NotNil(t, diskInfo.Kbt)
		require.IsType(t, 1.0, diskInfo.Kbt)
		require.NotNil(t, diskInfo.Tps)
		require.IsType(t, 1.0, diskInfo.Tps)
		require.NotNil(t, diskInfo.Mbs)
		require.IsType(t, 1.0, diskInfo.Mbs)
	})
}
