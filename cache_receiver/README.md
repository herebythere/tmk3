# Cache Receiver (SuperCache)

Deployable Cache as webservice.

## Setup

Create `config/config.json` as a configuration file. Use `config/config.json.example` as an example config.

Then run the following command:

`python3 build_and_run_supercache.py`

## Web API

SupderCache has an API with a single endpoint: `/`

### Requests

The following is an example of a JSON request body for a `hello world` request to SuperCache.

```JSON
// json
["INCR", "MYCOUNTER"]
```

SuperCache requests adhere to the following types:

```Golang
// Golang
type Statement = []interface{}
```

```Typescript
// Typescript
interface Statement = unknown[]
```

### Responses

The following is an example of the body of a response from SuperCache. Due to the nature of remote caches, anything can be returned.

```Golang
// Golang
type Statement = interface{}
```

```Typescript
// Typescript
type Statement = unknown
```

### Errors

The following is an example of a JSON response body for an error from SuperCache.

```JSON
[{
    "kind": "failed to exec",
    "message": "could not complete exec"
}]
```

SQL requests adhere to the following types:

```Golang
// Golang
type ErrorEntity struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}
```

```Typescript
// Typescript
interface ErrorEntity {
	kind    string
	message string
}
```

## License

Apache-2.0