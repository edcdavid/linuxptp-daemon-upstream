package utils

import (
	"sync"

	ptpv1 "github.com/k8snetworkplumbingwg/ptp-operator/api/v1"
)

// phcCache stores the mapping of interface names to PHC IDs
// This is populated from the NodePtpDevice CR and used as the source of truth
var (
	phcCache     map[string]string
	phcCacheLock sync.RWMutex
)

func init() {
	phcCache = make(map[string]string)
}

// UpdatePhcCacheFromNodePtpDevice updates the PHC cache from a NodePtpDevice CR
// This should be called whenever the NodePtpDevice CR is updated
func UpdatePhcCacheFromNodePtpDevice(nodePtpDevice *ptpv1.NodePtpDevice) {
	if nodePtpDevice == nil {
		return
	}

	phcCacheLock.Lock()
	defer phcCacheLock.Unlock()

	// Clear existing cache
	phcCache = make(map[string]string)

	// Populate from NodePtpDevice status
	for _, device := range nodePtpDevice.Status.Devices {
		if device.Name != "" && device.PhcId != "" {
			phcCache[device.Name] = device.PhcId
		}
	}
}

// GetPhcIdFromCache retrieves the PHC ID for an interface from the cache
// Returns empty string if not found
func GetPhcIdFromCache(ifname string) string {
	phcCacheLock.RLock()
	defer phcCacheLock.RUnlock()

	if phcId, found := phcCache[ifname]; found {
		return phcId
	}
	return ""
}

// ClearPhcCache clears the PHC ID cache
func ClearPhcCache() {
	phcCacheLock.Lock()
	defer phcCacheLock.Unlock()
	phcCache = make(map[string]string)
}

