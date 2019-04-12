// Package memcalc provides utility functions for working with Kubernetes-style memory resources
// which are expressed as strings such as 256Mi or 2Gi.  Methods are provided for converting
// between string representations and floats, and for math operations.
package memcalc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	_ = iota
	_
	// Mi = Megabytes
	Mi float64 = 1 << (10 * iota)
	// Gi = Gigabytes
	Gi
)

// Memory allows converstion between string and float representations and basic math operations on memory types
type Memory interface {
	// Add will parse the memory expressed as a string and return a new memory instance
	// equal to the sum of the current instance plus m
	Add(m string) (Memory, error)

	// AddF will return a new memory instance equal to the sum of the current instance plus m
	AddF(m float64) Memory

	// Sub will parse the memory expressed as a string and return a new memory instance
	// equal to the current instance minus m
	Sub(m string) (Memory, error)

	// Add will return a new memory instance equal to the current instance minus m
	SubF(m float64) Memory

	// ToString returns the Kubernetes-style memory value as a string rounded up to the nearest
	// megabyte.  Values over 1Gi will still be returned as an equivalent Mi value.
	ToString() string

	// ToFloat returns the memory value as a float64
	ToFloat64() float64
}

type memory struct {
	m float64
}

// New returns a new memory instance initialized at 0
func New() Memory {
	return memory{0}
}

// NewFromString parses a Kubernetes-style memory string (e.g., 256Mi, 1Gi)
func NewFromString(m string) (Memory, error) {
	f, err := memToFloat64(m)
	if err != nil {
		return nil, err
	}
	return memory{f}, nil
}

// NewFromFloat creates a new memory instance initialized to m
func NewFromFloat(m float64) Memory {
	return memory{m}
}

// Add will parse the memory expressed as a string and return a new memory instance
// equal to the sum of the current instance plus m
func (s memory) Add(m string) (Memory, error) {
	f, err := memToFloat64(m)
	if err != nil {
		return nil, err
	}
	return memory{s.m + f}, nil
}

// Sub will parse the memory expressed as a string and return a new memory instance
// equal to the current instance minus m
func (s memory) Sub(m string) (Memory, error) {
	f, err := memToFloat64(m)
	if err != nil {
		return nil, err
	}
	return memory{s.m - f}, nil
}

// AddF will return a new memory instance equal to the sum of the current instance plus m
func (s memory) AddF(m float64) Memory {
	return memory{s.m + m}
}

// Add will return a new memory instance equal to the current instance minus m
func (s memory) SubF(m float64) Memory {
	return memory{s.m - m}
}

// ToString returns the Kubernetes-style memory value as a string rounded up to the nearest
// megabyte.  Values over 1Gi will still be returned as an equivalent Mi value.
func (s memory) ToString() string {
	return float64ToMi(s.m)
}

// ToFloat returns the memory value as a float
func (s memory) ToFloat64() float64 {
	return s.m
}

func memToFloat64(s string) (float64, error) {
	switch {
	case strings.HasSuffix(s, "Mi"):
		mem, err := strconv.ParseFloat(strings.TrimSuffix(s, "Mi"), 64)
		if err != nil {
			return 0, fmt.Errorf("failed to convert memory string %s to float", s)
		}
		return mem * Mi, nil
	case strings.HasSuffix(s, "Gi"):
		mem, err := strconv.ParseFloat(strings.TrimSuffix(s, "Gi"), 64)
		if err != nil {
			return 0, fmt.Errorf("failed to convert memory string %s to float", s)
		}
		return mem * Gi, nil
	default:
		return 0, fmt.Errorf("failed to convert memory string %s to float, unknown units", s)
	}
}

func float64ToMi(m float64) string {
	return fmt.Sprintf("%dMi", int(math.Ceil(m/Mi)))
}
