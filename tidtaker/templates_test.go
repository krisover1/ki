package main

import (
	"testing"
	"time"

	"github.com/pocketbase/pocketbase/tools/types"
)

func TestFormatDate(t *testing.T) {
	tests := []struct {
		name     string
		input    string // RFC3339
		expected string
	}{
		{"zero value", "", "-"},
		{"april", "2026-04-21T10:00:00Z", "21. april 2026"},
		{"januar", "2026-01-01T00:00:00Z", "1. januar 2026"},
		{"desember", "2025-12-31T23:59:59Z", "31. desember 2025"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dt types.DateTime
			if tt.input != "" {
				parsed, err := time.Parse(time.RFC3339, tt.input)
				if err != nil {
					t.Fatalf("ugyldig testdato: %v", err)
				}
				if err := dt.Scan(parsed); err != nil {
					t.Fatalf("kunne ikke scanne dato: %v", err)
				}
			}
			got := formatDate(dt)
			if got != tt.expected {
				t.Errorf("formatDate(%q) = %q, ønsker %q", tt.input, got, tt.expected)
			}
		})
	}
}
