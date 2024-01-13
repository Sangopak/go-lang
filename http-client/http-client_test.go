package main

import (
	"testing"
	"github.com/jarcoal/httpmock"
)

func TestGetGttpSuccess(t *testing.T) {
	httpmock.Activate()
  	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost:8090/hello",
    httpmock.NewStringResponder(200, `{"message":"hello"}`))

	actualResponse, _ := GetHttp("http://localhost:8090/hello")

	if actualResponse.Status != "200" {
		t.Error("Did not get the correct Http status 200")
	}
}