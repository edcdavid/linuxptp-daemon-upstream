package daemon

import (
	ptpv1 "github.com/k8snetworkplumbingwg/ptp-operator/api/v1"
	ptputils "github.com/k8snetworkplumbingwg/linuxptp-daemon/pkg/utils"
)

// InitializePhcCacheFromNodePtpDevice initializes the PHC cache from NodePtpDevice CR
// This is a wrapper to avoid circular dependency between daemon and utils packages
func InitializePhcCacheFromNodePtpDevice(nodePtpDevice *ptpv1.NodePtpDevice) {
	ptputils.UpdatePhcCacheFromNodePtpDevice(nodePtpDevice)
}


