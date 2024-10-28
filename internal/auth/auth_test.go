package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		input         http.Header
		expectedKey   string
		expectedError string
	}{
		{
			name: "normal header",
			input: http.Header{
				"Authorization": []string{"ApiKey 123123123"},
			},
			expectedKey:   "123123123",
			expectedError: "",
		},

		{
			name:          "empty header",
			input:         http.Header{},
			expectedKey:   "",
			expectedError: "no authorization header included",
		},

		{
			name: "wrong token",
			input: http.Header{
				"Authorization": []string{"Bearer 123123123"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "too short auth header",
			input: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.input)
			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("GetAPIKey() error = %v, expectedError %v", err, tt.expectedError)
				}
			} else if err != nil {
				t.Errorf("GetAPIKey() unexpected error = %v", err)
			}
			if got != tt.expectedKey {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.expectedKey)
			}
		})
	}
}
