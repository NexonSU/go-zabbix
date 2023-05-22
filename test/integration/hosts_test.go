package integration

import (
	"testing"

	"github.com/fabiang/go-zabbix"
	"github.com/fabiang/go-zabbix/test"
)

func testHost(t *testing.T, params zabbix.HostGetParams) {
	session := test.GetTestSession(t)

	hosts, err := session.GetHosts(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting Hosts: %v", err)
		}
	}

	if len(hosts) == 0 {
		t.Skip("No Hosts found")
	}

	for i, host := range hosts {
		if host.HostID == "" {
			t.Fatalf("Host %d returned in response body has no Host ID", i)
		}
	}

	t.Logf("Validated %d Hosts", len(hosts))
}

func TestHostsTemplates(t *testing.T) {
	params := zabbix.HostGetParams{
		IncludeTemplates: true,
	}

	testHost(t, params)
}

func TestHostsGroups(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectGroups: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsApplications(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectApplications: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsDiscoveries(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectDiscoveries: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)

}

func TestHostsDiscoveryRule(t *testing.T) {

	params := zabbix.HostGetParams{
		SelectDiscoveryRule: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsGraphs(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectGraphs: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsHostDiscovery(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectHostDiscovery: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsWebScenarios(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectWebScenarios: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsInterfaces(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectInterfaces: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsInventory(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectInventory: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsItems(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectItems: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsMacros(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectMacros: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsParentTemplates(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectParentTemplates: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsScreens(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectScreens: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsTriggers(t *testing.T) {
	params := zabbix.HostGetParams{
		SelectTriggers: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}
