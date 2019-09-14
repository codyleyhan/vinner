package vinner_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/codyleyhan/vinner"
)

func TestService_GetVehicle(t *testing.T) {
	service := vinner.NewService()

	tests := map[string]struct {
		expected *vinner.Vehicle
		err      error
	}{
		"1G6AR5SX5F0123869": {
			expected: &vinner.Vehicle{
				VIN:       "1G6AR5SX5F0123869",
				Year:      2015,
				Make:      "CADILLAC",
				Model:     "CTS",
				Trim:      "",
				Doors:     4,
				BodyClass: vinner.Sedan,
			},
		},
		"4S4BRBCC3E3215258": {
			expected: &vinner.Vehicle{
				VIN:       "4S4BRBCC3E3215258",
				Year:      2014,
				Make:      "SUBARU",
				Model:     "Outback",
				Trim:      "Premium + CWP",
				Doors:     0,
				BodyClass: vinner.Wagon,
			},
		},
		"5S4BRBCC3E3215258": {
			err: vinner.InvalidVIN,
		},
		"4S4BRBCC3E321525": {
			err: vinner.InvalidVIN,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			vehicle, err := service.GetVehicle(context.Background(), name)
			require.Equal(t, tt.err, err)
			if tt.err == nil {
				require.NotNil(t, vehicle)
				assert.Equal(t, *tt.expected, *vehicle)
			}
		})
	}
}
