//go:build darwin
// +build darwin

package load

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStat(t *testing.T) {
	t.Run("test success get stats", func(t *testing.T) {
		loadInfo, err := GetStats()

		require.NoError(t, err)
		require.NotNil(t, loadInfo.Load1Min)
		require.IsType(t, 1.0, loadInfo.Load1Min)
		require.NotNil(t, loadInfo.Load5Min)
		require.IsType(t, 1.0, loadInfo.Load5Min)
		require.NotNil(t, loadInfo.Load15Min)
		require.IsType(t, 1.0, loadInfo.Load15Min)
	})
}
