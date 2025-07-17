package integration

import (
	"testing"

	"github.com/NexonSU/go-zabbix"
	"github.com/NexonSU/go-zabbix/test"
)

func TestMaintenanceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	session := test.GetTestSession(t)

	params := zabbix.MaintenanceGetParams{}

	maintenances, err := session.GetMaintenance(&params)
	if err != nil {
		if _, ok := err.(*zabbix.NotFoundError); !ok {
			t.Fatalf("Error getting maintenances: %v", err)
		}
	}

	if len(maintenances) == 0 {
		t.Skip("No maintenances found")
	}

	for i, maintenance := range maintenances {
		if maintenance.MaintenanceID == "" {
			t.Fatalf("Maintenance %d has no maintenance ID", i)
		}
	}

	t.Logf("Validated %d maintenance", len(maintenances))
}
