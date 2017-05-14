package netswi

import (
	"net/http/httptest"
	"net/http"
	"testing"
	"fmt"
)

func TestNetSwitching(t *testing.T) {
	local_server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer local_server.Close()

	client, err := NewClientBindedToIP("127.0.0.1")
	if err != nil { t.Fail() }

	// External resource should be unavailable
	if _, err := client.Get("http://google.com"); err == nil {
		t.Fail()
	}

	// TLS too
	if _, err := client.Get("https://google.com"); err == nil {
		t.Fail()
	}

	// Local server should be reachable
	if _, err := client.Get(local_server.URL); err != nil {
		t.Fail()
	}
}
