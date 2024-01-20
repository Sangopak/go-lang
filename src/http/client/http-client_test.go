package src

 import (
 	"errors"
 	"fmt"
 	"strings"
 	"testing"

 	"github.com/jarcoal/httpmock"
 )

 var TestEndpoint string = "http://localhost:8090/hello"
 func TestGetGttpSuccess(t *testing.T) {
 	httpmock.Activate()
   	defer httpmock.DeactivateAndReset()

 	// Exact URL match with success
 	httpmock.RegisterResponder("GET", TestEndpoint,
     httpmock.NewStringResponder(200, `{"message":"hello"}`))

 	actualResponse, _ := GetHttp(TestEndpoint)

 	if actualResponse.Status != "200" {
 		t.Error("Did not get the correct Http status 200")
 	}
 }

 func TestGetHttpFailure(t *testing.T) {
 	httpmock.Activate()
 	defer httpmock.DeactivateAndReset()

   	// Exact URL match with an error
   	httpmock.RegisterResponder("GET", TestEndpoint,
 	httpmock.NewErrorResponder(errors.New("500 Internal Server Error")))

 	_, actualError := GetHttp(TestEndpoint)
 	errorMessage := fmt.Sprint(actualError)
 	if !strings.Contains(errorMessage, "500 Internal Server Error"){
 		t.Error("Did not get the 500 error")
 	}
 } 