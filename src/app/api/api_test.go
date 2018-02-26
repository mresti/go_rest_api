package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server     *httptest.Server
	reader     io.Reader
	rootUrl    string
	anyUrl     string
	faviconUrl string
	statsURL   string
)

func init() {
	server = httptest.NewServer(Handlers())

	rootUrl = fmt.Sprintf("%s/", server.URL)
	anyUrl = fmt.Sprintf("%s/", server.URL)
	faviconUrl = fmt.Sprintf("%s/", server.URL)
	statsURL = fmt.Sprintf("%s/stats", server.URL)
}

func TestRootURL(t *testing.T) {
	reader = strings.NewReader("")
	request, err := http.NewRequest(http.MethodGet, rootUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	checkResponseCode(t, res.StatusCode, http.StatusOK)
	expected := `{"message":"Hello T3chFest 2018!!"}`
	checkBody(t, res.Body, expected)
}

func TestRootURL_BadRequest(t *testing.T) {
	reader = strings.NewReader("")
	request, err := http.NewRequest(http.MethodPost, rootUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	checkResponseCode(t, res.StatusCode, http.StatusMethodNotAllowed)
	expected := `{"error":"Invalid HTTP Method"}`
	checkBody(t, res.Body, expected)
}

func TestStatsURL(t *testing.T) {
	reader = strings.NewReader("")
	request, err := http.NewRequest(http.MethodGet, statsURL, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	checkResponseCode(t, res.StatusCode, http.StatusOK)
	expected := `{"visits":1}`
	checkBody(t, res.Body, expected)
}

func TestFaviconURL(t *testing.T) {
	reader = strings.NewReader("")
	request, err := http.NewRequest(http.MethodGet, faviconUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	checkResponseCode(t, res.StatusCode, http.StatusOK)
	expected := `{"message":"Hello T3chFest 2018!!"}`
	checkBody(t, res.Body, expected)
}

func checkBody(t *testing.T, bodyParm io.ReadCloser, expected string) {
	body, err := ioutil.ReadAll(bodyParm)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(body) {
		t.Errorf("Expected the message %s. Got %s", expected, body)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
