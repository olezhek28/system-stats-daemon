package helper

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetLoadStats ...
func GetLoadStats() ([]string, error) {
	cmd := exec.Command("iostat")
	res, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to exec %s: %w", cmd.String(), err)
	}

	return strings.Fields(string(res)), nil
}
