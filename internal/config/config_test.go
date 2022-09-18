package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("invalid config file", func(t *testing.T) {
		_, err := GetConfig("/tmp/bla.bla")
		require.Error(t, err)

		file, err := os.CreateTemp("", "config")
		if err != nil {
			t.FailNow()
			return
		}

		_, err = file.Write([]byte("bla bla json"))
		if err != nil {
			t.FailNow()
			return
		}

		_, err = GetConfig(file.Name())
		require.Error(t, err)
	})
}
