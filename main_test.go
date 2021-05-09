package main_test

import (
     "os"
     "bytes"
     _ "fmt"
     _ "strings"
     "testing"
     "reflect"
     "io/ioutil"
     _ "log"
     "net/http"
     "net/http/httptest"
     "encoding/json"
	 _ "github.com/prasadadireddi/scytaleapi/api"
	 _ "github.com/prasadadireddi/scytaleapi/api/models"
	 "github.com/prasadadireddi/scytaleapi/api/router"
	 _ "github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
    code := m.Run()
    os.Exit(code)
}

func TestGetWorkloads(t *testing.T) {
	var jsonStr = string("[{\"SpiffeID\":\"test\",\"selectors\":[\"Python:Java\",\"Java:Python\"]},{\"SpiffeID\":\"extratest\",\"selectors\":[\"Python1:Java1\",\"Java1:Python1\"]}]")

    req, _ := http.NewRequest("GET", "/api/v1/workloads", nil)
    response := executeRequest(req)
    var input interface{}
	var output interface{}

    checkResponseCode(t, http.StatusOK, response.Code)
    // body := string(response.Body);
    bytes, err := ioutil.ReadAll(response.Body)
    body := string(bytes)

	err = json.Unmarshal([]byte(jsonStr), &input)
	err = json.Unmarshal([]byte(body), &output)
	err = err

    // if body := response.Body.String(); body != jsonStr {
    if ! reflect.DeepEqual(input, output) {
        t.Errorf("Expected data not found. \nGot: %s \nExpected: %s", body, jsonStr)
    }
}

func TestWorkloadsSorted(t *testing.T) {
    var jsonStr = string("[{\"SpiffeID\":\"extratest\",\"selectors\":[\"Python1:Java1\",\"Java1:Python1\"]},{\"SpiffeID\":\"test\",\"selectors\":[\"Python:Java\",\"Java:Python\"]}]")

    req, _ := http.NewRequest("GET", "/api/v1/workloads/sorted", nil)
    response := executeRequest(req)

    var input interface{}
	var output interface{}

    checkResponseCode(t, http.StatusOK, response.Code)
    // body := string(response.Body);
    bytes, err := ioutil.ReadAll(response.Body)
    body := string(bytes)

	err = json.Unmarshal([]byte(jsonStr), &input)
	err = json.Unmarshal([]byte(body), &output)
	err = err

    // if body := response.Body.String(); body != jsonStr {
    if ! reflect.DeepEqual(input, output) {
        t.Errorf("Expected data not found. \nGot: %s \nExpected: %s", body, jsonStr)
    }
}


func TestCreateWorkload(t *testing.T) {
    var jsonStr = []byte(`{"SpiffeID":"example.com","selectors":["type:demo"]}`)
    req, _ := http.NewRequest("POST", "/api/v1/workload", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    response := executeRequest(req)
    checkResponseCode(t, http.StatusCreated, response.Code)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    if m["SpiffeID"] != "example.com" {
        t.Errorf("Expected SpiffeID to be 'example.com'. Got '%v'", m["SpiffeID"])
    }

}

func TestUpdateWorkload(t *testing.T) {
    var jsonStr = []byte(`{"selectors":["type:demo"]}`)
    req, _ := http.NewRequest("PUT", "/api/v1/workload/test", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    response := executeRequest(req)
    checkResponseCode(t, http.StatusCreated, response.Code)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    if m["SpiffeID"] != "test" {
        t.Errorf("Expected SpiffeID to be 'test'. Got '%v'", m["SpiffeID"])
    }

}

func TestDeleteWorkload(t *testing.T) {
    req, _ := http.NewRequest("DELETE", "/api/v1/workload/test", nil)
    req.Header.Set("Content-Type", "application/json")

    response := executeRequest(req)
    checkResponseCode(t, http.StatusCreated, response.Code)

    var m map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &m)

    if m["SpiffeID"] != "test" {
        t.Errorf("Expected SpiffeID to be 'test'. Got '%v'", m["SpiffeID"])
    }

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    // a.Router.ServeHTTP(rr, req)
    a := router.New()
    a.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}
