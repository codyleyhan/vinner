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

func TestService_GetMakes(t *testing.T) {
	service := vinner.NewService()

	makes, err := service.GetMakes(context.Background())
	require.NoError(t, err)
	require.NotNil(t, makes)

	assert.NotEmpty(t, makes)
}

func TestService_GetModelsForMake(t *testing.T) {
	service := vinner.NewService()

	tests := map[string]struct {
		req vinner.GetModelsRequest
		err error
	}{
		"honda no year": {
			req: vinner.GetModelsRequest{Make: "honda"},
		},
		"honda 2019": {
			req: vinner.GetModelsRequest{Make: "honda", Year: 2019},
		},
		"no make passed": {
			req: vinner.GetModelsRequest{Make: "", Year: 2019},
			err: vinner.InvalidModelsRequest,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			makes, err := service.GetModels(context.Background(), test.req)
			require.Equal(t, test.err, err)
			if err == nil {
				require.NotNil(t, makes)
				assert.NotEmpty(t, makes)
			}
		})
	}
}
