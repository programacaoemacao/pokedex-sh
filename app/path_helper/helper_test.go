package pathhelper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRepoRootPath(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		currentDir, err := GetRepoRootPath()
		require.NoError(t, err)
		require.Contains(t, currentDir, "pokedex-sh")
	})
}
