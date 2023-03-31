package util

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test - model for testing
type Test struct {
	Title       string
	Description string
	Method      string
	URL         string
	Headers     map[string]interface{}
	Body        string
	PreRequest  func()
	Request     func(w http.ResponseWriter, r *http.Request)
	PostRequest func(resp []byte) error
}

func TestUseCases(t *testing.T, testCases []Test) {
	for _, testCase := range testCases {
		fmt.Println(testCase.Title, "-", testCase.Description)

		// setup before testing
		testCase.PreRequest()

		// request
		payload := strings.NewReader(testCase.Body)

		req, err := http.NewRequest(testCase.Method, testCase.URL, payload)
		if err != nil {
			t.Errorf("Error while testing %s, Error - %s", testCase.URL, err.Error())
			return
		}

		response := httptest.NewRecorder()

		testCase.Request(response, req)

		// validate and clear any data after testing
		err = testCase.PostRequest(response.Body.Bytes())
		if err != nil {
			t.Errorf("Error while testing %s, Response - %s, Error - %s", testCase.URL, response.Body.String(), err.Error())
			return
		}
	}
}
