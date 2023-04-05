package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRespBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"id": 1, "title": "delectus aut autem", "completed": false}`)
	}))

	defer ts.Close()

	expected := `{"id": 1, "title": "delectus aut autem", "completed": false}`

	testCases := []struct {
		name        string
		url         string
		expected    string
		expectedErr error
	}{
		{
			name:        "valid URL",
			url:         ts.URL,
			expected:    expected,
			expectedErr: nil,
		},
		{
			name:        "invalid URL",
			url:         "invalid-url",
			expected:    "",
			expectedErr: errors.New("Get \"invalid-url\": unsupported protocol scheme \"\""),
		},
	}

	for _, tc := range testCases {
		actual := getRespBody(tc.url)

		assert.Equal(t, tc.expected, actual)
	}
}
