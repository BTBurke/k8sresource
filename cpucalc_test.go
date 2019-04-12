package k8sresource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToInt(t *testing.T) {
	tcs := []struct {
		Name      string
		Value     string
		Expected  int
		ShouldErr bool
	}{
		{Name: "100m", Value: "100m", Expected: 100, ShouldErr: false},
		{Name: "0.5", Value: "0.5", Expected: 500, ShouldErr: false},
		{Name: "1", Value: "1", Expected: 1000, ShouldErr: false},
		{Name: "1.0", Value: "1.0", Expected: 1000, ShouldErr: false},
		{Name: "1500m", Value: "1500m", Expected: 1500, ShouldErr: false},
		{Name: "Invalid suffix", Value: "1500x", Expected: 0, ShouldErr: true},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			f, err := cpuIntFromString(tc.Value)
			switch tc.ShouldErr {
			case true:
				assert.Error(t, err)
			default:
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, f)
			}
		})
	}
}

func TestCPUAdd(t *testing.T) {
	tcs := []struct {
		Name     string
		Value1   string
		Value2   string
		Expected int
	}{
		{Name: "less than 1", Value1: "256m", Value2: "256m", Expected: 512},
		{Name: "greater than 1", Value1: "600m", Value2: "500m", Expected: 1100},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			m, err := NewCPUFromString(tc.Value1)
			assert.NoError(t, err)
			mAdded, err := m.Add(tc.Value2)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, mAdded.ToMillicores())
		})
	}
}

func TestCPUSub(t *testing.T) {
	tcs := []struct {
		Name     string
		Value1   string
		Value2   string
		Expected int
	}{
		{Name: "zero", Value1: "256m", Value2: "256m", Expected: 0},
		{Name: "less than 1", Value1: "256m", Value2: "128m", Expected: 128},
		{Name: "greater than 1", Value1: "1000m", Value2: "0.5", Expected: 500},
		{Name: "negative", Value1: "500m", Value2: "1000m", Expected: -500},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			m, err := NewCPUFromString(tc.Value1)
			assert.NoError(t, err)
			mAdded, err := m.Sub(tc.Value2)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, mAdded.ToMillicores())
		})
	}
}
