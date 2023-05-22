package integration

import (
	"testing"

	"github.com/cavaliercoder/go-zabbix"
	"github.com/cavaliercoder/go-zabbix/test"
)

func TestProxy(t *testing.T) {
	session := test.GetTestSession(t)

	params := zabbix.ProxyGetParams{}

	proxies, err := session.GetProxies(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting proxies: %v", err)
		}
	}

	if len(proxies) == 0 {
		t.Skip("No proxies found")
	}

	for i, proxie := range proxies {
		if proxie.ProxyID == "" {
			t.Fatalf("Proxie %d has no proxy ID", i)
		}
	}

	t.Logf("Validated %d proxies", len(proxies))
}
