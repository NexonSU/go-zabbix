package integration

import (
	"testing"

	"github.com/NexonSU/go-zabbix"
	"github.com/NexonSU/go-zabbix/test"
)

func TestProxyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

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
