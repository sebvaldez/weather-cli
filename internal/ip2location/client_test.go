package ip2location

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get(t *testing.T) {
	// Set up a test HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
            "ip": "127.0.0.1",
            "country_code": "US",
            "country_name": "United States",
            "region_name": "California",
            "city_name": "Mountain View",
            "latitude": 37.3860,
            "longitude": -122.0838,
            "zip_code": "94043",
            "time_zone": "-07:00"
        }`)
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.Client = ts.Client()
	client.baseURL = ts.URL // Modify the unexported baseURL directly

	location, err := client.Get()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if location.City != "Mountain View" {
		t.Errorf("Expected city 'Mountain View', got '%s'", location.City)
	}
}
