package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"encoding/json"
)

// test api with incorrect chain parameter
// should return a 400 response
func TestAPIHandlerResponseStatusWithInvalidChain(t *testing.T) {

	// create http httpRequest and add url parameters
	httpRequest, err := http.NewRequest("GET", "/api/", nil)
	if err != nil {
		t.Fatal(err)
	}
	urlQuery := httpRequest.URL.Query()
	urlQuery.Add("address", "0x158190da5b0cb33a3afc8103a4c304c86cb8410c")
	urlQuery.Add("chain", "mainnettt")
	httpRequest.URL.RawQuery = urlQuery.Encode()

	// create the httpRequest recorder and the api handler and start the test
	httpRequestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(APIHandler)
	handler.ServeHTTP(httpRequestRecorder, httpRequest)

	//if response status is anything other than 400, fail
	if apiResponseStatus := httpRequestRecorder.Code; apiResponseStatus != 400 {
		t.Errorf("handler returned wrong status code: got %d want %d",
			apiResponseStatus, 400)
	}
}

// test api with correct parameters
func TestAPIHandlerResponseObject(t *testing.T) {
	var apiResponseObject APIResponseObject

	// create http httpRequest and add url parameters
	httpRequest, err := http.NewRequest("GET", "/api/", nil)
	if err != nil {
		t.Fatal(err)
	}
	urlQuery := httpRequest.URL.Query()
	urlQuery.Add("address", "0x0A8BF44EAd38A45425305D65a85703331b0a30BD")
	urlQuery.Add("chain", "mainnet")
	httpRequest.URL.RawQuery = urlQuery.Encode()

	// create the httpRequestRecorder and the api handler and start the test
	httpRequestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(APIHandler)
	handler.ServeHTTP(httpRequestRecorder, httpRequest)

	// if api response code is not 200, fail the test
	if apiResponseStatus := httpRequestRecorder.Code; apiResponseStatus != 200 {
		t.Errorf("handler returned wrong status code: got %d want %d",
			apiResponseStatus, 200)
	}

	// read api response body into buffer
	buffer, err := ioutil.ReadAll(httpRequestRecorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	// if unable to parse the json object into associated go struct, fail the test
	err = json.Unmarshal(buffer, &apiResponseObject)
	if err != nil {
		t.Errorf("unable to parse api response object into struct")
	}

}