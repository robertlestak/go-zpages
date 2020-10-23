package zpages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// ZPages contains the config params for health check endpoints
type ZPages struct {
	Drivers []interface{}
	Status  map[string]interface{}
}

// Response contains an execution response
type Response struct {
	Type  string
	Name  string
	Error string
}

func pingDriver(d interface{}, res chan *Response) {
	switch v := d.(type) {
	case *CloudFront:
		e := v.Ping()
		res <- v.Response(e)
	case *DynamoDB:
		e := v.Ping()
		res <- v.Response(e)
	case *Elasticsearch:
		e := v.Ping()
		res <- v.Response(e)
	case *HTTP:
		e := v.Ping()
		res <- v.Response(e)
		break
	case *Rekognition:
		e := v.Ping()
		res <- v.Response(e)
		break
	case *Redis:
		e := v.Ping()
		res <- v.Response(e)
		break
	case *S3:
		e := v.Ping()
		res <- v.Response(e)
		break
	case *SQL:
		e := v.Ping()
		res <- v.Response(e)
		break
	default:
		res <- &Response{Error: "driver not supported"}
		break
	}
}

// Ping iterates through all drivers and executes `Ping` method
// on configured and initialized driver
// TODO: Convert `Checker` struct to `interface` which implements `Ping` method. Currenly duplicating code.
func (z *ZPages) Ping() []*Response {
	var e []*Response
	c := make(chan *Response, len(z.Drivers))
	for _, d := range z.Drivers {
		go pingDriver(d, c)
	}
	for i := 0; i < len(z.Drivers); i++ {
		e = append(e, <-c)
	}
	return e
}

// Up exposes the most basic health check endpoint
// to confirm application is online, but does not check
// any upstream dependencies
func (z *ZPages) Up(w http.ResponseWriter, r *http.Request) {
	res := z.Ping()
	var e bool
	for _, v := range res {
		if v.Error != "" {
			e = true
		}
	}
	if e {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	jd, jerr := json.Marshal(res)
	if jerr != nil {
		http.Error(w, jerr.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(jd))
}

// Ready exposes the most basic health check endpoint
// currently wraps Up
func (z *ZPages) Ready(w http.ResponseWriter, r *http.Request) {
	z.Up(w, r)
}

// Live exposes the most basic health check endpoint
// currently wraps Up
func (z *ZPages) Live(w http.ResponseWriter, r *http.Request) {
	z.Up(w, r)
}

// Statusz returns the status object
func (z *ZPages) Statusz(w http.ResponseWriter, r *http.Request) {
	if z.Status == nil {
		z.Status = make(map[string]interface{})
	}
	v := os.Getenv("VERSION")
	if _, ok := z.Status["Version"]; !ok && v != "" {
		z.Status["Version"] = v
	}
	var env string
	env = os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("ENVIRONMENT")
	}
	if _, ok := z.Status["Environment"]; !ok && env != "" {
		z.Status["Environment"] = env
	}
	jd, jerr := json.Marshal(z.Status)
	if jerr != nil {
		http.Error(w, jerr.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(jd))
}

// HTTPHandlers registers the default HTTP handlers
func (z *ZPages) HTTPHandlers() {
	// rate limit to 1 request per second
	http.HandleFunc("/healthz", z.Up)
	http.HandleFunc("/livez", z.Live)
	http.HandleFunc("/readyz", z.Ready)
	http.HandleFunc("/statusz", z.Statusz)
}

// HTTPHandlersMux registers HTTP handlers on Gorilla Mux
func (z *ZPages) HTTPHandlersMux(r *mux.Router) {
	// rate limit to 1 request per second
	r.HandleFunc("/healthz", z.Up)
	r.HandleFunc("/livez", z.Live)
	r.HandleFunc("/readyz", z.Ready)
	r.HandleFunc("/statusz", z.Statusz)
}
