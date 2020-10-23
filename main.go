package main

import (
	"net/http"
	"time"

	"github.com/robertlestak/go-zpages/pkg/zpages"
)

func main() {
	// configure a driver
	h := &zpages.HTTP{
		Name:        "google.com",
		Address:     "https://google.com",
		Method:      "GET",
		StatusCodes: []int{200, 301, 302},
	}
	// add driver to healthz
	d := []interface{}{h}
	z := &zpages.ZPages{Drivers: d}
	// create status object
	// VERSION and ENV environment variables added automatically
	z.Status = func() map[string]interface{} {
		return map[string]interface{}{
			"Arbitrary": "data",
			"Foo":       "bar",
			"Time":      time.Now().String(),
		}
	}
	// register HTTP handlers
	z.HTTPHandlers()
	// start HTTP server
	http.ListenAndServe(":8080", nil)
}
