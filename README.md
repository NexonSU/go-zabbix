# go-zabbix

Go bindings for the Zabbix API

## Getting started

Get the package:

```
go get "github.com/NexonSU/go-zabbix"
```

```go
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/NexonSU/go-zabbix"
)

func main() {
	// Default approach - without session caching
	session, err := zabbix.NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix")
	if err != nil {
		panic(err)
	}

	version, err := session.GetVersion()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to Zabbix API v%s", version)
}
```

### Use session builder with caching.

You can use own cache by implementing SessionAbstractCache interface.
Optionally an http.Client can be passed to the builder, allowing to skip TLS verification, pass proxy settings, etc.

```go
func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true
			}
		}
	}

	cache := zabbix.NewSessionFileCache().SetFilePath("./zabbix_session")
	session, err := zabbix.CreateClient("http://zabbix/api_jsonrpc.php").
		WithCache(cache).
		WithHTTPClient(client).
		WithCredentials("Admin", "zabbix").
		Connect()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	version, err := session.GetVersion()

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("Connected to Zabbix API v%s", version)
}
```

## Running the tests

### Unit tests
Running the unit tests:

```bash
go test -v -short "./..."
# or:
make unittests
```

### Integration tests

To run the integration tests against a specific Zabbix Server version, you'll need Docker. Then start the containers:

```bash
export ZBX_VERSION=6.4
docker compose up -d
# server should be running in a minute
# run tests:
go test -v -run Integration "./..."
# or:
make integration
```

## License

Released under the [GNU GPL License](https://github.com/NexonSU/go-zabbix?tab=GPL-2.0-1-ov-file)
