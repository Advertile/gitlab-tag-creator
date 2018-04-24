package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gitlab "github.com/xanzy/go-gitlab"
)

func TestIsValidVersionType(t *testing.T) {
	if isValidVersionType("major") == false {
		t.Errorf("The version type \"major\" should be valid")
	}
	if isValidVersionType("minor") == false {
		t.Errorf("The version type \"minor\" should be valid")
	}
	if isValidVersionType("patch") == false {
		t.Errorf("The version type \"patch\" should be valid")
	}

	if isValidVersionType("majores") == true {
		t.Errorf("The version type \"majores\" should not be valid")
	}
}

func TestBumpVersion(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/ProjectID/repository/tags", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[{"name": "1.0.1"},{"name": "1.0.0"}]`)
	})

	v, _ := bumpVersion("ProjectID", client, "major")
	if v != "2.0.0" {
		t.Errorf("The major version was not increased correctly")
	}

	v, _ = bumpVersion("ProjectID", client, "minor")
	if v != "1.1.0" {
		t.Errorf("The minor version was not increased correctly")
	}

	v, _ = bumpVersion("ProjectID", client, "patch")
	if v != "1.0.2" {
		t.Errorf("The patch version was not increased correctly")
	}
}

func TestBumpVersionWithNoTags(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/projects/ProjectID/repository/tags", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[]`)
	})

	_, err := bumpVersion("ProjectID", client, "major")

	if err.Error() != "There are no tags yet. Create the first tag manually" {
		t.Errorf("Wrong error message")
	}
}

// setup sets up a test HTTP server along with a gitlab.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (*http.ServeMux, *httptest.Server, *gitlab.Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the Gitlab client being tested.
	client := gitlab.NewClient(nil, "")
	client.SetBaseURL(server.URL)

	return mux, server, client
}

// teardown closes the test HTTP server.
func teardown(server *httptest.Server) {
	server.Close()
}
