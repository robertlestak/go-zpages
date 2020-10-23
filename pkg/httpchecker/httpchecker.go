package httpchecker

import (
	"bytes"
	"fmt"
	"net/http"
)

// Checker wraps a basic HTTP checker
type Checker struct {
	Address     *string
	Method      *string
	Body        *[]byte
	StatusCodes *[]int
}

// Ping connects to a HTTP endpoint and checks if the status code is in the defined list
func (hc *Checker) Ping() error {
	c := http.Client{}
	req, rerr := http.NewRequest(*hc.Method, *hc.Address, bytes.NewReader(*hc.Body))
	if rerr != nil {
		return rerr
	}
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	var sc bool
	for _, s := range *hc.StatusCodes {
		if res.StatusCode == s {
			sc = true
			break
		}
	}
	if !sc {
		return fmt.Errorf("status code error. Expected one of %+v, got %v", hc.StatusCodes, res.StatusCode)
	}
	return nil
}
