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

In addition to dependency health checks, this package also creates a `/statusz` endpoint which returns a JSON `map[string]interface{}`.

This is automatically populated with the environment variables `VERSION` and `ENV`, and additional key/value pairs can be added when initialized.

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
