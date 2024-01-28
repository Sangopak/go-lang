package src

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var TestEndpoint string = "http://localhost:8090/hello"

func TestGetGttpSuccess(t *testing.T) {
	expectedResponse := "200"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match with success
	httpmock.RegisterResponder("GET", TestEndpoint,
		httpmock.NewStringResponder(200, `{"message":"hello"}`))

	actualResponse, _ := GetHttp(TestEndpoint)

	assert.Equal(t, expectedResponse, actualResponse.Status, "Did not get the correct Http status 200")
}

func TestGetHttpFailure(t *testing.T) {
	expectedResponse := "500 Internal Server Error"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match with an error
	httpmock.RegisterResponder("GET", TestEndpoint,
		httpmock.NewErrorResponder(errors.New("500 Internal Server Error")))

	_, actualError := GetHttp(TestEndpoint)
	errorMessage := fmt.Sprint(actualError)

	assert.True(t, strings.Contains(errorMessage, expectedResponse), "Did not get the 500 error")
}

func BenchmarkGetHttp(b *testing.B) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match with success
	httpmock.RegisterResponder("GET", TestEndpoint,
		httpmock.NewStringResponder(200, `{"message":"hello"}`))

	for i := 0; i < b.N; i++ {
		GetHttp(TestEndpoint)
	}

}
