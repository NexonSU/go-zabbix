package integration

import (
	"testing"

	"github.com/NexonSU/go-zabbix"
	"github.com/NexonSU/go-zabbix/test"
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

func TestHostsTemplatesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		IncludeTemplates: true,
	}

	testHost(t, params)
}

func TestHostsGroupsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectGroups: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsApplicationsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectApplications: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsDiscoveriesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectDiscoveries: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)

}

func TestHostsDiscoveryRuleIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectDiscoveryRule: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsGraphsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectGraphs: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsHostDiscoveryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectHostDiscovery: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsWebScenariosIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectWebScenarios: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsInterfacesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectInterfaces: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsInventoryIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectInventory: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsItemsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectItems: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsMacrosIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectMacros: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsParentTemplatesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectParentTemplates: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsScreensIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectScreens: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}

func TestHostsTriggersIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	params := zabbix.HostGetParams{
		SelectTriggers: zabbix.SelectExtendedOutput,
	}

	testHost(t, params)
}
