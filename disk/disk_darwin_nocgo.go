// +build darwin
// +build !cgo

package disk

import "github.com/Azareal/gopsutil/internal/common"

func IOCounters(names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
