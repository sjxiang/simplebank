package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig("..")
	require.NoError(t, err)

	t.Logf("%+v\n", cfg)
}

// $ go test -v ./config_test.go ./config.go