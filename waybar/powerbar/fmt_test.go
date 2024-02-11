package main

import (
	"testing"
	"time"
)

func TestApplyFmt(t *testing.T) {
	cases := []struct {
		name               string
		inputFmtStr        string
		inputBatteryStatus BatteryStatus
		expected           string
	}{
		{
			name:        "Full battery status with usage",
			inputFmtStr: "FULL - {usage}W",
			inputBatteryStatus: BatteryStatus{
				Usage: 0,
			},
			expected: "FULL - 0.00W",
		},
		{
			name:        "capacity, time until full, charging",
			inputFmtStr: "{state} {capacity}% - {usage}W - {H}h {M}m",
			inputBatteryStatus: BatteryStatus{
				Capacity:      33,
				State:         Charging,
				TimeUntilFull: time.Minute * 372,
				IsCharging:    true,
				Usage:         10.24,
			},
			expected: "Charging 33% - 10.24W - 6h 12m",
		},
		{
			name:        "capacity, time left",
			inputFmtStr: "{capacity}% - {usage}W - {H}h {M}m",
			inputBatteryStatus: BatteryStatus{
				Capacity:       33,
				TimeUntilEmpty: time.Minute * 372,
				Usage:          3.24,
			},
			expected: "33% - 3.24W - 6h 12m",
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(c *testing.T) {
			actual := applyFmt(tc.inputFmtStr, tc.inputBatteryStatus)
			if actual != tc.expected {
				t.Errorf("\nexpected: %s\nactual:   %s", tc.expected, actual)
			}
		})
	}
}
