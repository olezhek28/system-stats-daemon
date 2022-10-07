//go:build linux
// +build linux

package disk

import (
	"github.com/olezhek28/system-stats-daemon/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetStats ...
func GetStats() (*model.DiskInfo, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
