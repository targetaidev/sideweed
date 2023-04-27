// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"testing"
)

func TestGetHealthCheckURL_Valid(t *testing.T) {
	testCases := []struct {
		name            string
		endpoint        string
		healthCheckPath string
		healthCheckPort int
		want            string
	}{
		{
			name:            "PortSetToZero",
			endpoint:        "http://server1:9000",
			healthCheckPath: "/health",
			want:            "http://server1:9000/health",
		},
		{
			name:            "PortSetToNonZeroValue",
			endpoint:        "http://server1:9000",
			healthCheckPath: "/health",
			healthCheckPort: 4242,
			want:            "http://server1:4242/health",
		},
		{
			name:            "PortSetToUpperLimit",
			endpoint:        "http://server1:9000",
			healthCheckPath: "/health",
			healthCheckPort: portUpperLimit,
			want:            "http://server1:65535/health",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			healthCheckURL, err := getHealthCheckURL(tc.endpoint, tc.healthCheckPath, tc.healthCheckPort)
			if err != nil {
				t.Errorf("Expected no error, got %q", err)
			}

			if healthCheckURL != tc.want {
				t.Errorf("Expected %q, got %q", tc.want, healthCheckURL)
			}
		})
	}
}

func TestGetHealthCheckURL_Invalid(t *testing.T) {
	want := ""

	testCases := []struct {
		name            string
		endpoint        string
		healthCheckPath string
		healthCheckPort int
	}{
		{
			name:     "BadEndpoint",
			endpoint: "bad",
		},
		{
			name:            "PortNumberBelowLowerLimit",
			endpoint:        "http://server1:9000",
			healthCheckPath: "/health",
			healthCheckPort: portLowerLimit - 1,
		},
		{
			name:            "PortNumberAboveUpperLimit",
			endpoint:        "http://server1:9000",
			healthCheckPath: "/health",
			healthCheckPort: portUpperLimit + 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			healthCheckURL, err := getHealthCheckURL(tc.endpoint, tc.healthCheckPath, tc.healthCheckPort)

			if err == nil {
				t.Errorf("Expected an error")
			}

			if healthCheckURL != want {
				t.Errorf("Expected %q, got %q", want, healthCheckURL)
			}
		})
	}
}
