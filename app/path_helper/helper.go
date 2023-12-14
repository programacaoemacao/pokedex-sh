package pathhelper

import (
	"fmt"
	"os"
	"strings"
)

func GetRepoRootPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", nil
	}

	parts := strings.Split(currentDir, "pokedex-sh")
	if len(parts) < 1 {
		return "", fmt.Errorf("can't split the path")
	}

	fullPath := parts[0] + "pokedex-sh"
	return fullPath, nil
}
