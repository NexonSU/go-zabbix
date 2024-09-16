package integration

import (
	"testing"

	"github.com/fabiang/go-zabbix"
	"github.com/fabiang/go-zabbix/test"
)

func TestHostgroupsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	session := test.GetTestSession(t)

	params := zabbix.HostgroupGetParams{}

	hostgroups, err := session.GetHostgroups(params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting Hostgroups: %v", err)
		}
	}

	if len(hostgroups) == 0 {
		t.Skip("No Hostgroups found")
	}

	for i, hostgroup := range hostgroups {
		if hostgroup.GroupID == "" {
			t.Fatalf("Hostgroup %d returned in response body has no Group ID", i)
		}
	}

	t.Logf("Validated %d Hostgroups", len(hostgroups))
}
