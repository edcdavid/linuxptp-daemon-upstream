package utils_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/k8snetworkplumbingwg/linuxptp-daemon/pkg/testhelpers"
	"github.com/k8snetworkplumbingwg/linuxptp-daemon/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	teardownTests := testhelpers.SetupTests()
	defer teardownTests()
	os.Exit(m.Run())
}

type testCase struct {
	ifname                 string
	expectedClockIdentifier string
	checkPrefix            bool // if true, just check that result starts with expected prefix
}

func Test_GetClockIdentifier(t *testing.T) {
	testCases := []testCase{
		// Special clock identifiers (should return as-is)
		{"CLOCK_REALTIME", "CLOCK_REALTIME", false},
		{"master", "master", false},
		{"/dev/ptp0", "/dev/ptp0", false},
		{"/dev/ptp1", "/dev/ptp1", false},
		// Fallback cases (interfaces that don't match or have no PHC - will return original or /dev/ptpX)
		{"wlan", "wlan", false},   // no PHC, returns original
		{"lo", "lo", false},       // no PHC, returns original
		{"virbr", "virbr", false}, // no PHC, returns original
		{"docker", "docker", false}, // no PHC, returns original
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s->%s", tc.ifname, tc.expectedClockIdentifier), func(t *testing.T) {
			result := utils.GetClockIdentifier(tc.ifname)
			if tc.checkPrefix {
				assert.True(t, len(result) > 0 && (result == tc.ifname || result[:len(tc.expectedClockIdentifier)] == tc.expectedClockIdentifier))
			} else {
				assert.Equal(t, tc.expectedClockIdentifier, result)
			}
		})
	}
}

// Test_GetClockIdentifier_RealInterfaces tests with real network interfaces
// These tests may return /dev/ptpX if the interface has PHC support,
// or the original name if PHC lookup fails
func Test_GetClockIdentifier_RealInterfaces(t *testing.T) {
	testInterfaces := []string{
		"eth0",
		"eth1",
		"enP2s2f0np0",
		"enP1s1f1np1",
		"ens1f3np3",
	}
	
	for _, iface := range testInterfaces {
		t.Run(iface, func(t *testing.T) {
			result := utils.GetClockIdentifier(iface)
			// Result should either be /dev/ptpX or the original interface name
			if result != iface {
				assert.Regexp(t, `^/dev/ptp\d+$`, result, "If changed, should be /dev/ptpX format")
			}
		})
	}
}

// Test_GetClockIdentifier_VLAN tests VLAN interface handling
// PTP runs on the base interface, so VLAN tags should be stripped
func Test_GetClockIdentifier_VLAN(t *testing.T) {
	testCases := []struct {
		ifname string
	}{
		{"eth1.100"},
		{"eth1.100.XYZ"},
		{"enP2s2f0np0.100"},
		{"enP1s1f1np1.200"},
		{"enP10s5f3np2.300.XYZ"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.ifname, func(t *testing.T) {
			result := utils.GetClockIdentifier(tc.ifname)
			// If PHC lookup succeeded, result should be /dev/ptpX (WITHOUT VLAN tag)
			// If it failed, result should be original name
			if result != tc.ifname {
				assert.Regexp(t, `^/dev/ptp\d+$`, result, "Should be /dev/ptpX format without VLAN tag")
				assert.NotContains(t, result, ".", "VLAN tag should NOT be in PHC ID")
			}
		})
	}
}
