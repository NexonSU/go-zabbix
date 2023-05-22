package integration

import (
	"testing"

	"github.com/cavaliercoder/go-zabbix"
	"github.com/cavaliercoder/go-zabbix/test"
)

func TestHostInterfaces(t *testing.T) {
	session := test.GetTestSession(t)

	params := zabbix.HostInterfaceGetParams{}

	hostInterfaces, err := session.GetHostInterfaces(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting interfaces: %v", err)
		}
	}

	if len(hostInterfaces) == 0 {
		t.Skip("No interfaces found")
	}

	for i, hostInterface := range hostInterfaces {
		if hostInterface.InterfaceID == "" {
			t.Fatalf("Interface %d has no interface ID", i)
		}
	}

	t.Logf("Validated %d interfaces", len(hostInterfaces))
}
