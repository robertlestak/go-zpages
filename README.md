# go-zpages

Standard z-page health check methods to be imported as needed.

## Supported HTTP Libraries

### net/http

If application uses default `net/http` package, `HTTPHandlers()` method will attach handlers automatically.

### gorilla/mux

If application uses `gorilla/mux` package, `HTTPHandlersMux(r *mux.Router)` will attach handlers automatically.

### Custom HTTP Integrations

All handlers are exported and available for use in custom integrations.

See `pkg/zpages/zpages.go` `HTTPHandlers` method for usage.

It is recommended to use the same paths as defined in `HTTPHandlers` for consistency across applications.

## Status Page

In addition to dependency health checks, this package also creates a `/statusz` endpoint.

This is configurable by defining a `func() map[string]interface{}` that will return the desired values when `/statusz` is called.

This is automatically populated with the environment variables `VERSION` and `ENV`, and additional key/value pairs can be added when initialized.

For example:

```
z := new(zpages.ZPages)
z.Status = func() map[string]interface{} {
	return map[string]interface{}{
		"Time":      time.Now().String(),
		"Arbitrary": "data",
		"Foo":       "bar",
		"Dynamic":	func() string {
			return internalDynamicDataAPI()
		}
	}
}
z.HTTPHandlers()

// will produce:
// {
//		"Arbitrary": "data",
//		"Foo": "bar",
//		"Dynamic": "output from internalDynamicDataAPI()",
//		"Time": "2020-10-23 14:16:48.522569 -0700 PDT m=+0.727506354"
// }
```

## Supported Drivers

All health checks are configured with a `Driver` object.

Each driver struct is exposed through the `zpages` package.

See `pkg/zpages/drivers.go` for full list of driver structs.

```
type HTTP struct {
	Name        string
	Address     string
	Method      string
	Body        []byte
	StatusCodes []int
}
```

`Name` is a user-identifiable field to reference the checked resource.

See driver configuration below for specific driver configurations.

See [./docs/drivers.md](./docs/drivers.md) for supported drivers and configurations.

See `main.go` for a basic example usage.
