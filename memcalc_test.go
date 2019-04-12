package k8sresource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemConvertStringToFloat(t *testing.T) {
	tcs := []struct {
		Name      string
		Value     string
		Expected  float64
		ShouldErr bool
	}{
		{Name: "megabytes < 1Gi", Value: "256Mi", Expected: 256 * Mi, ShouldErr: false},
		{Name: "1Gi", Value: "1Gi", Expected: 1 * Gi, ShouldErr: false},
		{Name: "10Gi", Value: "10Gi", Expected: 10 * Gi, ShouldErr: false},
		{Name: "1024Mi", Value: "1024Mi", Expected: 1 * Gi, ShouldErr: false},
		{Name: "1500Mi", Value: "1500Mi", Expected: 1500 * Mi, ShouldErr: false},
		{Name: "Invalid suffix", Value: "1500Xi", Expected: 0, ShouldErr: true},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			f, err := memToFloat64(tc.Value)
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

func TestMemConvertFloatToString(t *testing.T) {
	tcs := []struct {
		Name     string
		Value    float64
		Expected string
	}{
		{Name: "megabytes < 1Gi", Expected: "256Mi", Value: 256 * Mi},
		{Name: "1Gi", Expected: "1024Mi", Value: 1 * Gi},
		{Name: "10Gi", Expected: "10240Mi", Value: 10 * Gi},
		{Name: "1024Mi", Expected: "1024Mi", Value: 1024 * Mi},
		{Name: "1500Mi", Expected: "1500Mi", Value: 1500 * Mi},
		{Name: "-256Mi", Expected: "-256Mi", Value: -256 * Mi},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			f := float64ToMi(tc.Value)
			assert.Equal(t, tc.Expected, f)
		})
	}
}

func TestMemAdd(t *testing.T) {
	tcs := []struct {
		Name     string
		Value1   string
		Value2   string
		Expected float64
	}{
		{Name: "less than 1Gi", Value1: "256Mi", Value2: "256Mi", Expected: 512 * Mi},
		{Name: "greater than 1Gi", Value1: "256Mi", Value2: "1Gi", Expected: 256*Mi + 1*Gi},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			m, err := NewMemFromString(tc.Value1)
			assert.NoError(t, err)
			mAdded, err := m.Add(tc.Value2)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, mAdded.ToFloat64())
		})
	}
}

func TestMemSub(t *testing.T) {
	tcs := []struct {
		Name     string
		Value1   string
		Value2   string
		Expected float64
	}{
		{Name: "zero", Value1: "256Mi", Value2: "256Mi", Expected: 0},
		{Name: "less than 1Gi", Value1: "256Mi", Value2: "128Mi", Expected: 128 * Mi},
		{Name: "greater than 1Gi", Value1: "1Gi", Value2: "512Mi", Expected: 512 * Mi},
		{Name: "negative", Value1: "512Mi", Value2: "1024Mi", Expected: -512 * Mi},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			m, err := NewMemFromString(tc.Value1)
			assert.NoError(t, err)
			mAdded, err := m.Sub(tc.Value2)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, mAdded.ToFloat64())
		})
	}
}
