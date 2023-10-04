package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadEnv(t *testing.T) {
	t.Run("load .env", func(t *testing.T) {
		env := LoadEnv(filepath.Join("testdata", ".env"))
		require.Equal(t, env.HOST, "127.0.0.1")
	})

	t.Run("panic on missing file", func(t *testing.T) {
		require.Panics(t, func() {
			LoadEnv("missing")
		})
	})

	t.Run("panic on malformed .env", func(t *testing.T) {
		err := os.Unsetenv("HOST")
		require.NoError(t, err)
		require.Panics(t, func() {
			LoadEnv(filepath.Join("testdata", "broken.env"))
		})
	})
}
