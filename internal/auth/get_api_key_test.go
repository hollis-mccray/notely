package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key           string
		value         string
		expect        string
		expectedError string
	}{
		{
			expectedError: "no authorization header included",
		},
		{
			key:           "Authorization",
			expectedError: "no authorization header included",
		},
		{
			key:           "Authorization",
			value:         "tacotacotaco",
			expectedError: "malformed authorization header",
		},
		{
			key:           "Authorization",
			value:         "Bearer tacotacotaco",
			expectedError: "malformed authorization header",
		},
		{
			key:           "Authorization",
			value:         "ApiKey tacotacotaco",
			expect:        "tacotacotaco",
			expectedError: "no error extected",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedError) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}
